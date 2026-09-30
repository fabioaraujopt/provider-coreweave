# Provider CoreWeave

<div style="text-align: center;">

[![CI](https://github.com/fabioaraujopt/provider-coreweave/actions/workflows/ci.yml/badge.svg)](https://github.com/fabioaraujopt/provider-coreweave/actions/workflows/ci.yml)
[![GitHub release](https://img.shields.io/github/release/fabioaraujopt/provider-coreweave/all.svg)](https://github.com/fabioaraujopt/provider-coreweave/releases)
[![License](https://img.shields.io/github/license/fabioaraujopt/provider-coreweave)](LICENSE)

</div>

`provider-coreweave` is a [Crossplane](https://crossplane.io/) provider that is
built using [Upjet](https://github.com/crossplane/upjet) code generation tools
and exposes XRM-conformant managed resources for
[CoreWeave](https://www.coreweave.com/), wrapping CoreWeave's official
[`coreweave/coreweave`](https://registry.terraform.io/providers/coreweave/coreweave/latest)
Terraform provider.

All CoreWeave resources are built on the Terraform Plugin Framework, so the
provider runs them in-process (Upjet's no-fork architecture): no Terraform CLI,
no provider binary and no workspace on disk inside the pod.

> [!NOTE]
> The provider is in **alpha**. Every kind is `v1alpha1` and fields may still
> change between minor releases. A kind moves to `v1beta1` once it passes
> create, update, import and delete end-to-end tests against CoreWeave.

## Resources

| API group | Kinds |
|---|---|
| `cks` | `Cluster` |
| `networking` | `VPC` |
| `objectstorage` | `Bucket`, `BucketPolicy`, `BucketVersioning`, `BucketSettings`, `BucketLifecycleConfiguration`, `BucketInventory`, `AccessKey`, `OrganizationAccessPolicy` |
| `inference` | `Gateway`, `Deployment`, `CapacityClaim` |
| `sandbox` | `ManagedRunner` |
| `workloadfederation` | `OIDCConfig` |

Every kind is available namespaced (`<group>.coreweave.m.crossplane.io`, the
Crossplane v2 default) and cluster-scoped (`<group>.coreweave.crossplane.io`,
legacy).

Cross-resource references: `Cluster.vpcIdRef`, `Cluster.sharedStorageClusterRef`,
`ManagedRunner.clusterIdRef`, `Deployment.gatewayIdsRefs`,
`Deployment.model.bucketRef`, and `bucketRef` on every bucket sub-resource.

## Getting Started

### Install the provider

Pick a version from the [releases](https://github.com/fabioaraujopt/provider-coreweave/releases).
Requires Crossplane v2.

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-coreweave
spec:
  package: ghcr.io/fabioaraujopt/provider-coreweave:v0.1.0
```

Or with the Crossplane CLI:

```bash
crossplane xpkg install provider ghcr.io/fabioaraujopt/provider-coreweave:v0.1.0
```

### Install only the resources you need

The provider supports Crossplane's safe-start, so its resources are
[ManagedResourceDefinitions](https://docs.crossplane.io/latest/managed-resources/managed-resource-definitions/)
that only become CRDs when a
[ManagedResourceActivationPolicy](https://docs.crossplane.io/latest/managed-resources/managed-resource-activation-policies/)
activates them. The Crossplane Helm chart activates everything by default. To
activate only object storage, install Crossplane with
`--set provider.defaultActivations={}` and apply:

```yaml
apiVersion: apiextensions.crossplane.io/v1alpha1
kind: ManagedResourceActivationPolicy
metadata:
  name: coreweave-object-storage
spec:
  activate:
    - "*.objectstorage.coreweave.m.crossplane.io"
```

### Configure credentials

Create an API access token in the CoreWeave Cloud Console, store it in a
Secret, and point a `ClusterProviderConfig` (or a namespaced `ProviderConfig`)
at it:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: coreweave-creds
  namespace: crossplane-system
stringData:
  credentials: |
    {"token": "CW-SECRET-..."}
---
apiVersion: coreweave.m.crossplane.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: coreweave-creds
      namespace: crossplane-system
      key: credentials
```

| Key | Required | Description |
|---|---|---|
| `token` | yes | CoreWeave API access token (`CW-SECRET-...`). |
| `endpoint` | no | CoreWeave API endpoint. Defaults to `https://api.coreweave.com/`. |
| `s3_endpoint` | no | AI Object Storage endpoint. Defaults to `https://cwobject.com`; use `http://cwlota.com` from inside CoreWeave. |
| `http_timeout` | no | HTTP client timeout, e.g. `30s`. |

> [!WARNING]
> Don't set `COREWEAVE_API_TOKEN` or `COREWEAVE_*_ENDPOINT` on the provider
> pod. The upstream Terraform provider lets those variables override every
> ProviderConfig.

### Create a resource

```yaml
apiVersion: objectstorage.coreweave.m.crossplane.io/v1alpha1
kind: Bucket
metadata:
  name: my-unique-bucket-name   # bucket names are global
  namespace: default
spec:
  forProvider:
    zone: US-EAST-04A
  providerConfigRef:
    kind: ClusterProviderConfig
    name: default
```

More examples: [`examples/namespaced/`](examples/namespaced/) (hand-written)
and [`examples-generated/`](examples-generated/) (generated from the Terraform
docs, one per kind).

## Supply Chain Security

Release packages are signed with keyless [cosign](https://github.com/sigstore/cosign)
and carry SLSA build provenance and an SPDX SBOM as
[GitHub artifact attestations](https://docs.github.com/actions/security-for-github-actions/using-artifact-attestations).
Each [release](https://github.com/fabioaraujopt/provider-coreweave/releases)
lists the package digest and the commands to verify it:

```bash
cosign verify ghcr.io/fabioaraujopt/provider-coreweave@<digest> \
  --certificate-identity-regexp '^https://github.com/fabioaraujopt/provider-coreweave/\.github/workflows/publish-provider-package\.yml@' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com
gh attestation verify oci://ghcr.io/fabioaraujopt/provider-coreweave@<digest> \
  --repo fabioaraujopt/provider-coreweave
```

The latest release is scanned for vulnerabilities every week; findings appear
in the repository's Security tab.

## Developing

```bash
make submodules    # crossplane/build
make generate      # Terraform schema -> API types, controllers, CRDs, examples
make build         # provider binary, image and package
make local-deploy  # kind cluster + Crossplane + this provider
```

See [CONTRIBUTING.md](CONTRIBUTING.md) for prerequisites, adding a resource and
testing against CoreWeave, and [docs/UPJET_GUIDE.md](docs/UPJET_GUIDE.md) for
how the provider is put together: the in-process Terraform provider, external
names per resource, and bumping CoreWeave's Terraform provider.

Releases follow [docs/RELEASING.md](docs/RELEASING.md).

## Report a Bug

For filing bugs, suggesting improvements, or requesting new resources or
features, open an [issue](https://github.com/fabioaraujopt/provider-coreweave/issues/new/choose).
Report security issues privately as described in [SECURITY.md](SECURITY.md).

For general help with Crossplane and Upjet, the
[Crossplane Slack](https://slack.crossplane.io) is the best place to ask.

## License

The provider is released under the [Apache 2.0 license](LICENSE).
