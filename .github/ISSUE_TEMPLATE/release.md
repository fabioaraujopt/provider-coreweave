---
name: Release
about: Track a release. For maintainers.
labels: release
---
<!-- The full process is in docs/RELEASING.md. -->

### Checklist

* [ ] CI on `main` is green.
* [ ] For a minor release, create the release branch, e.g. `release-0.2` for `v0.2.0`.
  * Patch releases reuse the existing branch; backport fixes with `backport release-X.Y` labels first.
* [ ] Tag the release branch with a v-prefixed version (`git push origin vX.Y.Z`, or the `Tag` workflow).
* [ ] If the `Tag` workflow was used, run `Publish Provider Package` from the tag.
* [ ] The publish workflow succeeded: package pushed, signed and attested, GitHub release created.
* [ ] Release notes reviewed; every PR labelled `breaking-change` has upgrade instructions.
* [ ] Install the published package on a scratch cluster and check the provider is `Healthy`.
* [ ] Announce the release.
