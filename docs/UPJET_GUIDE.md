# Building a Crossplane provider for CoreWeave with Upjet

This is how `provider-coreweave` was built, what the current (Upjet v2, 2026)
best practices are, and how to maintain it. It follows the upstream
[Generating a provider](https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md)
guide and deviates where CoreWeave's provider needs it.

## 1. What Upjet does

Upjet reads a Terraform provider's schema (`terraform providers schema -json`)
and generates, for every resource you opt in:

- Go API types and CRDs, in two flavours since Upjet v2:
  cluster-scoped (`coreweave.crossplane.io`) and namespaced Crossplane v2
  managed resources (`coreweave.m.crossplane.io`, with `ProviderConfig` and
  `ClusterProviderConfig`).
- A controller per resource that translates `spec.forProvider` into Terraform
  config, calls the Terraform provider's CRUD, and writes back
  `status.atProvider`.
- Example manifests (`examples-generated/`) from the provider docs.

There are three ways a generated controller can talk to Terraform:

| Mode | Upjet option | When |
|---|---|---|
| Terraform CLI (legacy) | `WithIncludeList` | Only when you cannot import the provider's Go code |
| No-fork, Plugin SDKv2 | `WithTerraformPluginSDKIncludeList` + `WithTerraformProvider` | SDKv2 resources |
| **No-fork, Plugin Framework** | `WithTerraformPluginFrameworkIncludeList` + `WithTerraformPluginFrameworkProvider` | Framework resources, **what CoreWeave uses** |

No-fork means the provider's Go code runs inside the Crossplane controller.
It's the recommended mode: no `terraform` processes, far less memory, faster
reconciles, and a tiny image. Upbound's AWS/Azure/GCP providers use it.

## 2. Why the Terraform provider is forked

`terraform-provider-coreweave` is 100 % Terraform Plugin Framework, which is
ideal for no-fork. But its provider constructor lives in `internal/provider`,
and Go forbids importing `internal/` from another module.

The fork (branch `crossplane`, tag `v0.24.0-xp.1`) adds exactly one file:

```go
// xpprovider/xpprovider.go
package xpprovider

func New(version string) fwprovider.Provider {
    return provider.New(version)()
}
```

This is the same pattern Upbound uses (`xpprovider` packages in their AWS and
Azure forks). `provider-coreweave/go.mod` points at it with:

```
replace github.com/coreweave/terraform-provider-coreweave => github.com/your-org/terraform-provider-coreweave v0.24.0-xp.1
```

Worth proposing upstream to CoreWeave: if they accept a public `xpprovider`
package, the fork goes away.

## 3. Scaffolding

```bash
# "Use this template" on crossplane/upjet-provider-template, then:
PROVIDER_NAME_LOWER=coreweave PROVIDER_NAME_NORMAL=CoreWeave \
ORGANIZATION_NAME=your-org CRD_ROOT_GROUP=crossplane.io ./hack/prepare.sh
```

Then in the `Makefile`:

```make
TERRAFORM_PROVIDER_SOURCE  := coreweave/coreweave
TERRAFORM_PROVIDER_REPO    := https://github.com/coreweave/terraform-provider-coreweave
TERRAFORM_PROVIDER_VERSION := 0.24.0
TERRAFORM_DOCS_PATH        := docs/resources
```

`TERRAFORM_VERSION` stays at 1.5.7: the template refuses 1.6+ because of the
BSL licence. The CLI is only used at generation time to dump `schema.json`.

Because everything is no-fork, the image (`cluster/images/provider-coreweave/Dockerfile`)
was stripped of the Terraform CLI, the provider binary and `terraformrc.hcl`.

## 4. Wiring the in-process provider

`config/provider.go`

```go
ujconfig.WithIncludeList([]string{}),                              // nothing via CLI
ujconfig.WithTerraformPluginFrameworkIncludeList(ExternalNameConfigured()),
ujconfig.WithTerraformPluginFrameworkProvider(xpprovider.New("")),
ujconfig.WithSchemaTraversers(&ujconfig.SingletonListEmbedder{}),
```

- The include list must be emptied explicitly; its default is `.+`, and a
  resource in two lists makes Upjet panic.
- `SingletonListEmbedder` renders `MaxItems: 1` blocks as objects instead of
  one-element arrays. It's the default for every new Upjet provider; turning
  it on later is a breaking API change.

`internal/clients/coreweave.go` maps the ProviderConfig secret onto the
Terraform provider block and hands Upjet a fresh framework provider:

```go
ps.Configuration = map[string]any{"token": creds["token"], ...}
ps.FrameworkProvider = xpprovider.New(providerVersion)
```

`cmd/provider/main.go` adds an `OperationTrackerStore` to both controller
option sets. The template doesn't set one (its sample is CLI-based), and
framework controllers dereference it, so without it the provider panics on
the first reconcile.

## 5. External names, resource by resource

The external name (`crossplane.io/external-name`) is how Crossplane finds the
real resource. Upjet only generates resources that have one configured, so
this table *is* the resource list.

| Terraform resource | Config | Why |
|---|---|---|
| `cks_cluster`, `networking_vpc`, `inference_*`, `workload_federation_oidc_config`, `object_storage_access_key` | `FrameworkResourceWithComputedIdentifier("id", zero-UUID)` | The API assigns the id. Before creation Upjet reads with a placeholder id; CoreWeave returns NotFound, the TF resource drops it from state, and Upjet creates it. |
| `sandbox_managed_runner` | `ParameterAsIdentifier("runner_id")` | You choose `runner_id`; TF `id` equals it. |
| `object_storage_bucket`, `object_storage_organization_access_policy` | `ParameterAsIdentifier("name")` + read `name` back from state | No `id` attribute at all, so the default "read `id` from state" would fail. |
| `bucket_policy`, `bucket_settings`, `bucket_versioning`, `bucket_lifecycle_configuration` | `IdentifierFromProvider` + read `bucket` from state | Keyed by bucket. `bucket` stays in spec so it can use `bucketRef`. |
| `bucket_inventory` | same, external name = `name`, identifiers `bucket` + `name` | Import ID is `<bucket>:<name>`. |

Things to verify against the real API (§8): that every Read answers NotFound
(not InvalidArgument) for the zero-UUID placeholder.

## 6. Groups, kinds, references, secrets

One package per API group under `config/<group>/`, shared by both scopes.
The upstream guide duplicates each file under `config/cluster/` and
`config/namespaced/`; that's only needed when a reference uses the deprecated
`Type` field. Using `TerraformName` lets Upjet resolve the right package per
scope, so one copy is enough.

- Groups/kinds are set explicitly. Upjet would otherwise derive
  `object.coreweave…/StorageBucket` from `coreweave_object_storage_bucket`.
- References: `vpc_id`, `shared_storage_cluster_id`, `cluster_id`,
  `gateway_ids` (list, so `gatewayIdsRefs`), `model.bucket`, and `bucket` on
  bucket sub-resources.
- Sensitive inputs (inference gateway W&B `api_key`, sandbox `spec.base`,
  `spec.overrides.env`) become `*SecretRef` fields automatically.
- Sensitive outputs go to the connection secret. `AccessKey` additionally
  publishes `access_key_id` / `secret_access_key`.

## 7. Generate, build, run

```bash
make submodules
go mod tidy
make generate      # terraform schema -> scraper -> generator -> controller-gen -> angryjet
make build
make run           # out-of-cluster, against the current kubecontext
make local-deploy  # kind + Crossplane + the provider
```

Check after `make generate`:

- `package/crds/` has 30 CRDs (15 × 2 scopes) plus the ProviderConfig CRDs.
- `examples-generated/` exists; compare with the hand-written `examples/`.
- Field names in `examples/` match the generated CRDs. The hand-written
  manifests were written against the schema docs, not generated types.

## 8. Testing against CoreWeave

Uptest (`make e2e`) applies `examples/` to a kind cluster, waits for Ready,
imports, and deletes. Use a dedicated CoreWeave org or project:

```bash
export UPTEST_CLOUD_CREDENTIALS='{"token":"CW-SECRET-..."}'
make e2e UPTEST_EXAMPLE_LIST="examples/namespaced/networking/v1alpha1/vpc.yaml,examples/namespaced/objectstorage/v1alpha1/bucket.yaml"
```

Start with Bucket and VPC (fast, cheap), then Cluster (slow). Resources pass
from `v1alpha1` to `v1beta1` once they pass uptest with create, update,
import and delete.

## 9. Bumping the CoreWeave provider

1. In the fork: rebase `crossplane` onto the new upstream tag and tag
   `v<new>-xp.1` (see `CROSSPLANE_FORK.md`).
2. Here: update the `replace` line and `TERRAFORM_PROVIDER_VERSION`, then
   `go mod tidy && make generate`.
3. New resources won't appear until you add them to `ExternalNameConfigs`.
   `config/schema.json` shows what's new.
4. `make crddiff` in CI flags breaking API changes.

## References

- Upjet docs: `generating-a-provider.md`, `configuring-a-resource.md`,
  `upjet-v2-upgrade.md`, `testing-with-uptest.md` in crossplane/upjet
- crossplane/upjet-provider-template (the scaffold used here)
- crossplane-contrib/provider-upjet-aws (reference for framework resources and
  `FrameworkResourceWithComputedIdentifier`)
- CoreWeave Terraform provider docs: registry.terraform.io/providers/coreweave/coreweave
