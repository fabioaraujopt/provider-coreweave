# Contributing

Thanks for helping improve provider-coreweave. Bug reports, resource requests
and pull requests are all welcome.

## Prerequisites

- Go 1.26.8+ (the version in `go.mod`)
- Docker (Docker Desktop, Colima or similar)
- `goimports` on your `PATH`: `go install tool` installs the version pinned in
  `go.mod` into `$(go env GOPATH)/bin`
- Network access to the Go proxy, `buf.build` (CoreWeave's generated API
  clients), `releases.hashicorp.com` and `registry.terraform.io`
- A checkout path without spaces. Make can't handle them.

If you use Colima, point the Crossplane CLI at its Docker socket:
`export DOCKER_HOST=unix://$HOME/.colima/default/docker.sock`.

## Development workflow

```bash
make submodules    # crossplane/build
make generate      # Terraform schema -> API types, controllers, CRDs, examples
make reviewable    # generate, lint and test
make build         # provider binary, image and package for your platform
make local-deploy  # kind cluster + Crossplane + this provider
```

`make run` runs the provider out-of-cluster against your current kubeconfig
context. Check which cluster that is first.

Generated code (`apis/`, `internal/controller/`, `package/crds/`,
`examples-generated/`, `config/schema.json`, `config/provider-metadata.yaml`)
is committed. CI regenerates it and fails if the result differs, so run
`make generate` and commit the output with your change.

## Adding or changing a resource

Upjet only generates resources that have an external name configured.

1. Add the Terraform resource to `ExternalNameConfigs` in
   [`config/external_name.go`](config/external_name.go). Pick the external
   name strategy from the table in
   [docs/UPJET_GUIDE.md](docs/UPJET_GUIDE.md#5-external-names-resource-by-resource).
2. Set its group, kind and references in `config/<group>/config.go`, creating
   the package and registering it in `config/provider.go` if it's a new group.
3. Run `make generate`.
4. Add an example under `examples/namespaced/<group>/v1alpha1/` and the
   cluster-scoped equivalent under `examples/cluster/`. Check it with
   `kubectl apply --dry-run=server --validate=strict` on a `make local-deploy`
   cluster.

To bump CoreWeave's Terraform provider, follow
[docs/UPJET_GUIDE.md §9](docs/UPJET_GUIDE.md#9-bumping-the-coreweave-provider).

## Testing

### Unit tests and lint

```bash
make test
make lint
```

### End-to-end tests

[Uptest](https://github.com/crossplane/uptest) applies examples to a kind
cluster with the provider installed, waits for them to become ready, imports
and deletes them. It talks to the real CoreWeave API, so use a dedicated test
organization and a token for it.

Locally:

```bash
export UPTEST_CLOUD_CREDENTIALS='{"token":"CW-SECRET-..."}'
make e2e UPTEST_EXAMPLE_LIST="examples/namespaced/objectstorage/v1alpha1/bucket.yaml"
```

On a pull request, a maintainer can comment:

```
/test-examples="examples/namespaced/objectstorage/v1alpha1/bucket.yaml"
```

That runs the same test in GitHub Actions with the repository's
`UPTEST_CLOUD_CREDENTIALS` secret and reports the result as a status check.

A kind moves from `v1alpha1` to `v1beta1` once its example passes Uptest.

## Pull requests

- Keep each PR focused on one change and describe how you tested it.
- CI must pass: lint, unit tests, generated-code check, breaking CRD change
  report and a local deploy.
- Label changes that break the API `breaking-change`, so they end up in the
  release notes with upgrade instructions.
- Label fixes that should reach a released minor version
  `backport release-X.Y` before merging (see [docs/RELEASING.md](docs/RELEASING.md)).

## Releases

See [docs/RELEASING.md](docs/RELEASING.md).

## Code of Conduct

This project follows the [CNCF Code of Conduct](CODE_OF_CONDUCT.md).
