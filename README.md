# provider-coreweave

A [Crossplane](https://crossplane.io/) provider for [CoreWeave](https://www.coreweave.com/),
generated with [Upjet](https://github.com/crossplane/upjet) from CoreWeave's
official Terraform provider ([`coreweave/coreweave`](https://registry.terraform.io/providers/coreweave/coreweave/latest)).

All 15 CoreWeave resources are built on the Terraform Plugin Framework, so this
provider runs them **in-process** (Upjet's no-fork architecture): no Terraform
CLI, no provider binary and no workspace on disk inside the pod.

| Group (`*.coreweave.crossplane.io` / `*.coreweave.m.crossplane.io`) | Kinds |
|---|---|
| `cks` | `Cluster` |
| `networking` | `VPC` |
| `objectstorage` | `Bucket`, `BucketPolicy`, `BucketVersioning`, `BucketSettings`, `BucketLifecycleConfiguration`, `BucketInventory`, `AccessKey`, `OrganizationAccessPolicy` |
| `inference` | `Gateway`, `Deployment`, `CapacityClaim` |
| `sandbox` | `ManagedRunner` |
| `workloadfederation` | `OIDCConfig` |

Every kind is generated twice: cluster-scoped (`coreweave.crossplane.io`) and
namespaced Crossplane v2 (`coreweave.m.crossplane.io`).

Cross-resource references: `Cluster.vpcIdRef`, `Cluster.sharedStorageClusterRef`,
`ManagedRunner.clusterIdRef`, `Deployment.gatewayIdsRefs`,
`Deployment.model.bucketRef`, and `bucketRef` on every bucket sub-resource.

## Credentials

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: coreweave-creds
  namespace: crossplane-system
stringData:
  credentials: |
    {"token": "CW-SECRET-...", "s3_endpoint": "https://cwobject.com"}
---
apiVersion: coreweave.m.crossplane.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef: {name: coreweave-creds, namespace: crossplane-system, key: credentials}
```

Accepted keys: `token` (required), `endpoint`, `s3_endpoint` (use
`http://cwlota.com` when running inside CoreWeave), `http_timeout`.

> Do not set `COREWEAVE_API_TOKEN` / `COREWEAVE_*_ENDPOINT` on the provider pod:
> the upstream Terraform provider lets those env vars override every
> ProviderConfig.

## Build it

Requires Go 1.26.7+, Docker, and network access to the Go proxy, `buf.build`
(CoreWeave's generated API clients), `releases.hashicorp.com` and
`registry.terraform.io`.

```bash
make submodules                # crossplane/build
go mod tidy                    # resolves the forked TF provider + deps
make generate                  # schema.json -> APIs, controllers, CRDs, examples
make build                     # provider binary + xpkg
make run                       # run out-of-cluster against your current kubecontext
```

Then commit the generated code (`apis/`, `internal/controller/`, `package/crds/`,
`examples-generated/`, `config/schema.json`, `config/provider-metadata.yaml`).

## How it's put together

See [docs/UPJET_GUIDE.md](docs/UPJET_GUIDE.md) for the full walkthrough: why
the Terraform provider is forked, how each resource's external name was chosen,
and how to bump to a new CoreWeave provider release.
