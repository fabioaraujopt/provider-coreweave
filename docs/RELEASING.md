# Releasing

provider-coreweave follows the release process of the crossplane-contrib
providers: release branches per minor version, tags created from those
branches, and packages published by GitHub Actions. The only difference is the
registry: packages go to `ghcr.io/fabioaraujopt/provider-coreweave`.

## Versioning

- [Semantic versioning](https://semver.org/) with a `v` prefix: `v0.1.0`,
  `v0.1.1`, `v0.2.0-rc.1`.
- While the provider is `v0.x`, a minor release may contain breaking API
  changes. Call them out in the release notes (PRs labelled `breaking-change`).
- A version with a `-` suffix (`-rc.1`) is published as a GitHub pre-release.

## Branches

| Branch | Purpose |
|---|---|
| `main` | Development. CI runs on every push and PR; nothing is published. |
| `release-X.Y` | One per minor version (`release-0.1` for `v0.1.x`). Patch releases are tagged here. CI runs on every push. |

Fixes land on `main` first and are backported: label the PR
`backport release-X.Y` before merging it, and on merge the Backport workflow
opens a PR that cherry-picks it onto `release-X.Y`. For a PR merged without the
label, cherry-pick it yourself (`git cherry-pick -x <sha>`) into a PR against
the release branch.

## What a release produces

The [Publish Provider Package](../.github/workflows/publish-provider-package.yml)
workflow:

1. builds the provider for `linux/amd64` and `linux/arm64` and pushes the
   package to `ghcr.io/fabioaraujopt/provider-coreweave:<version>` (using the
   shared `crossplane-contrib/provider-workflows` publish workflow);
2. signs it with keyless cosign and attaches SLSA build provenance and an SPDX
   SBOM as GitHub artifact attestations;
3. creates the GitHub release with install and verification instructions and
   notes generated from the merged PRs. If a release for the tag already
   exists (for example one you drafted), its notes are kept and the install
   section is added on top.

## Cutting a minor release (vX.Y.0)

1. Open a [release issue](../.github/ISSUE_TEMPLATE/release.md) to track it.
2. Make sure CI on `main` is green.
3. Create the release branch from `main`:
   ```bash
   git checkout main && git pull
   git checkout -b release-X.Y
   git push origin release-X.Y
   ```
4. Tag it. Either push the tag yourself, which also starts the publish
   workflow:
   ```bash
   git checkout release-X.Y
   git tag -a vX.Y.0 -m "vX.Y.0"   # -s instead of -a to sign it
   git push origin vX.Y.0
   ```
   or run the [Tag](../.github/workflows/tag.yaml) workflow with "Use workflow
   from" set to `release-X.Y`. Tags created by that workflow don't start other
   workflows, so then run **Publish Provider Package** with "Use workflow from"
   set to the tag (Tags → `vX.Y.0`) and version `vX.Y.0`.
5. Wait for the publish workflow. It fails early if it isn't running from the
   tag it publishes.
6. Review the GitHub release notes; add upgrade instructions for anything
   labelled `breaking-change`.
7. Smoke-test the published package: install it on a scratch cluster and
   check the provider becomes `Healthy`.

## Cutting a patch release (vX.Y.Z)

1. Backport the fixes to `release-X.Y` (see [Branches](#branches)).
2. Tag `vX.Y.Z` on `release-X.Y` and publish, as in steps 4 to 7 above.

## Release candidates

Tag `vX.Y.0-rc.N` on `release-X.Y` and publish as usual. The GitHub release
is marked as a pre-release.

## One-time repository setup

These are settings, not code, so they aren't in this repository:

- **Make the package public.** The first publish creates
  `ghcr.io/fabioaraujopt/provider-coreweave` as a private package. In the
  package settings, set its visibility to public so Crossplane can pull it
  without credentials.
- **Protect `main` and `release-*`** with a ruleset that requires pull requests
  and the CI checks to pass.
- **Protect `v*` tags** so only maintainers can create or move them.
- **Secrets for end-to-end tests:** `UPTEST_CLOUD_CREDENTIALS`, the JSON the
  ProviderConfig secret takes (`{"token": "CW-SECRET-..."}`), for a dedicated
  test organization. See [CONTRIBUTING.md](../CONTRIBUTING.md#end-to-end-tests).
- **Renovate:** install the [Renovate GitHub App](https://github.com/apps/renovate)
  on the repository; the configuration is in `.github/renovate.json5`.
- **Private vulnerability reporting:** enable it under Settings → Security, as
  [SECURITY.md](../SECURITY.md) asks reporters to use it.
