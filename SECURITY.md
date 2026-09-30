# Security Policy

## Supported Versions

Security fixes go into the latest minor release (the newest `release-X.Y`
branch) and `main`. Older minor versions aren't patched while the provider is
`v0.x`.

## Reporting a Vulnerability

Please don't open a public issue for a security problem.

Report it privately through
[GitHub private vulnerability reporting](https://github.com/fabioaraujopt/provider-coreweave/security/advisories/new).
Include:

- a description of the vulnerability and its impact,
- the provider version (or commit) and Crossplane version,
- steps or manifests to reproduce it.

You should get an acknowledgement within a few working days. Once a fix is
released, the advisory is published with credit to the reporter, unless you
prefer otherwise.

Vulnerabilities in CoreWeave's Terraform provider or API belong to
[CoreWeave](https://github.com/coreweave/terraform-provider-coreweave/security);
vulnerabilities in Crossplane or Upjet themselves go to the
[Crossplane security team](https://github.com/crossplane/crossplane/security/policy).

## Verifying Releases

Release packages are signed and carry build provenance and an SBOM. See
[Supply Chain Security](README.md#supply-chain-security).
