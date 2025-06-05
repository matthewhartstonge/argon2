# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.3.2](https://github.com/matthewhartstonge/argon2/compare/v1.3.1...v1.3.2) (2025-06-05)


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.38.0 to 0.39.0 ([#93](https://github.com/matthewhartstonge/argon2/issues/93)) ([313b810](https://github.com/matthewhartstonge/argon2/commit/313b810d8b4f4c6cf8076ad8f9614dd284b2dd62))

## [1.3.1](https://github.com/matthewhartstonge/argon2/compare/v1.3.0...v1.3.1) (2025-05-06)


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.37.0 to 0.38.0 ([#90](https://github.com/matthewhartstonge/argon2/issues/90)) ([4d70c13](https://github.com/matthewhartstonge/argon2/commit/4d70c13642edccdd2f8acc1255be8e8bb79dcafa))

## [1.3.0](https://github.com/matthewhartstonge/argon2/compare/v1.2.1...v1.3.0) (2025-04-22)


### Features

* return error on attempting to hash `argon2d`. ([71f8bcb](https://github.com/matthewhartstonge/argon2/commit/71f8bcb6797b19eeb4b142e7f7a6b7d56cee521c))


### Bug Fixes

* **lint:** migrates if-else blocks to tagged switches (QF1003). ([04cdec5](https://github.com/matthewhartstonge/argon2/commit/04cdec5d85be929c70f569a0da4d128ba565014c))

## [1.2.1](https://github.com/matthewhartstonge/argon2/compare/v1.2.0...v1.2.1) (2025-04-08)


### Bug Fixes

* Bump golang.org/x/crypto from 0.36.0 to 0.37.0 ([#87](https://github.com/matthewhartstonge/argon2/issues/87)) ([3c08e2b](https://github.com/matthewhartstonge/argon2/commit/3c08e2bd89c9b90792e94198ce542440d018a6db))

## [1.2.0](https://github.com/matthewhartstonge/argon2/compare/v1.1.1...v1.2.0) (2025-03-06)

### Features

* **deps:** upgrades to go@1.23.0. ([9c86600](https://github.com/matthewhartstonge/argon2/commit/9c86600313babe362ca9cbaae5d270931e72e8c2))

note: `go@1.21` introduced a change in how the go directive works, now enforcing Minimal Version Selection (MVS). The best thing to do is upgrade your Go toolchain to `n-1` and set the `go` directive in your `go.mod` file to `go 1.(N-1).0` - notice the requirement for the patch set to `0`.

If you are a library maintainer, PLEASE make sure to only set this to `0` and remove the `toolchain` directive, otherwise you will make downstream users will cry (myself included).

Why? As part of an internal Go proposal, all golang.org/x libraries now only support n-1.

> I propose that each time that a new major Go release 1.N.0 is made, the go directive in all golang.org/x repos will be upgraded to go 1.(N-1).0. For example, when Go 1.28.0 is released, golang.org/x modules would have their go directive set to go 1.27.0.

Refer: https://go.googlesource.com/proposal/+/master/design/69095-x-repo-continuous-go.md

## [1.1.1](https://github.com/matthewhartstonge/argon2/compare/v1.1.0...v1.1.1) (2024-12-23)


### Bug Fixes

* **cmd/argon2:** enable ldflags variable configuration. ([988c8fe](https://github.com/matthewhartstonge/argon2/commit/988c8fec2409ed091efe93ae9ab4d99da7b851f8))

## [1.1.0](https://github.com/matthewhartstonge/argon2/compare/v1.0.3...v1.1.0) (2024-12-21)


### Features

* **cmd/argon2:** adds argon2 cli. ([32a22f9](https://github.com/matthewhartstonge/argon2/commit/32a22f986e28f574553af5a4540b5d74439f1492))
* **goreleaser:** adds configuration for automated builds. ([c8385f8](https://github.com/matthewhartstonge/argon2/commit/c8385f89c37ee265cb8afdfa34fdb56777daf9d1))

## [1.0.3](https://github.com/matthewhartstonge/argon2/compare/v1.0.2...v1.0.3) (2024-12-12)


### Miscellaneous Chores

* release 1.0.3 ([601f492](https://github.com/matthewhartstonge/argon2/commit/601f49212436450d975dc2b4f9f30c6d0018ab8f))
* chore: Bump golang.org/x/crypto from 0.30.0 to 0.31.0 ([3bb40c5](https://github.com/matthewhartstonge/argon2/commit/3bb40c5e87345e039ea1f406b10fd07850f8120c))

## [1.0.2](https://github.com/matthewhartstonge/argon2/compare/v1.0.1...v1.0.2) (2024-11-10)


### Miscellaneous Chores

* release 1.0.2 ([64bcf86](https://github.com/matthewhartstonge/argon2/commit/64bcf86837060e454fafb59acc3d12feb3a3322d))
* chore: Bump golang.org/x/crypto from 0.28.0 to 0.29.0 ([7d61928](https://github.com/matthewhartstonge/argon2/commit/7d619289f4cafec245e657e6ea183d48f00e7cbd))
* chore: Bump golang.org/x/crypto from 0.27.0 to 0.28.0 ([dbe7209](https://github.com/matthewhartstonge/argon2/commit/dbe72099829775cac5519e4afcd5b475200bebef))

## [1.0.1](https://github.com/matthewhartstonge/argon2/compare/v1.0.0...v1.0.1) (2024-09-12)


### Miscellaneous Chores

* release 1.0.1 ([f68ab6f](https://github.com/matthewhartstonge/argon2/commit/f68ab6fc71041029ca749a899646046fb9f10ce3))
* chore: Bump golang.org/x/crypto from 0.26.0 to 0.27.0 ([#50](https://github.com/matthewhartstonge/argon2/pull/50))

## [1.0.0](https://github.com/matthewhartstonge/argon2/compare/v0.3.4...v1.0.0) (2023-12-03)


### ⚠ BREAKING CHANGES

* **deps:** bump golang.org/x/crypto from 0.13.0 to 0.16.0

### Features

* **deps:** bump golang.org/x/crypto from 0.13.0 to 0.16.0 ([8dbc527](https://github.com/matthewhartstonge/argon2/commit/8dbc52707b213d2c69660edb118b56f915eec4b0))


### Miscellaneous Chores

* release 1.0.0 ([fbe015c](https://github.com/matthewhartstonge/argon2/commit/fbe015cd6ebc9dc6890d00d5f09cd4b42583fc28))

## [0.3.4](https://github.com/matthewhartstonge/argon2/compare/v0.3.3...v0.3.4) (2023-09-06)


### Miscellaneous Chores

* **deps:** updates dependencies ([efcbcce](https://github.com/matthewhartstonge/argon2/commit/efcbcce98406c8304d41b681330864af14e7aeb1))

## [0.3.3](https://github.com/matthewhartstonge/argon2/compare/v0.3.2...v0.3.3) (2023-07-04)


### Bug Fixes

* **.github/dependabot:** fixes branch targets in dependabot configuration. ([64ab947](https://github.com/matthewhartstonge/argon2/commit/64ab947d1921dc80d8352516cf9c711044b2ed8f))

## [Unreleased]
## [v0.3.2] - 2022-11-18
### Changed
- deps: updates to `golang.org/x/crypto@v0.3.0`.
- .github: adjusts dependabot to file version updates against development.

## [v0.3.1] - 2022-11-10
### Changed
- deps: updates to `golang.org/x/crypto@v0.2.0`.

## [v0.3.0] - 2022-11-05
### Added
- .github: adds support for github actions, codeowners and dependabot.

### Changed
- deps: updates to `go@1.18` and `golang.org/x/crypto@v0.1.0`.
- readme: update badges and remove references to travis ci.
- *: update project licensing.

### Fixed
- argon2: fixes grammar and comment line length.

### Removed
- dep: removes support for dep.
- travis: removes travis configuration.

## [v0.2.1] - 2022-02-22
GPG signed release.

### Added
- tests: adds benchmarks to compare the default recommendations.

### Changed
- travis: updates to test against go1.17.
- deps: updates `go.mod` to `go@1.17` and `x/crypto` to latest.
- explicitly ignores returned error values under benchmarks (errcheck).
- argon2: updates recommended defaults as described by RFC9106.

## [v0.2.0] - 2022-02-22
### Added
- tests: adds benchmarks to compare the default recommendations.

### Changed
- travis: updates to test against go1.17.
- deps: updates `go.mod` to `go@1.17` and `x/crypto` to latest.
- explicitly ignores returned error values under benchmarks (errcheck).
- argon2: updates recommended defaults as described by RFC9106.

## [v0.1.5] - 2021-08-06
### Changed
- travis: require go >= v1.9
- deps: bumps to the latest version of `golang.org/x/crypto`.
- Upgrades to mitigate users importing a vulnerable version of `golang.org/x/crypto/ssh` that contains CVE-2020-9283.
- Upgrades to mitigate users importing a vulnerable version of `golang.org/x/text` that contains CVE-2020-14040.
- travis: updated to test against go 1.12+, migrates to go mod for dependency management.
- travis: forces use of go modules under ci.
- travis: sets `go@v1.11.4` as the lowest supported `go mod` version due to a change in go build.

## [v0.1.4] - 2021-08-06
### Changed
- readme: updated.

## [v0.1.3] - 2021-08-06
### Added
- deps: adds support for go modules.

### Changed
- travis: updates to test against `go@{1.12, 1.13, 1.14}`
- deps: update `x/crypto@master` to point to latest commit.
  - This is mainly to mitigate users from CVEs in other `x/crypto` implementations, namely CVE-2020-7919 (cryptobyte) and CVE-2020-9283 (ssh).
- readme: adds a tl;dr, updates benchmarks.

### Fixed
- _example: fixes calling location of `VerifyEncoded()`

### Removed
- tests: removes benchmarking against native bindings to remove dependencies under go mod.

## [v0.1.2] - 2018-09-18
### Added
- tests: added benchmarks to compare against native argon2 bindings.

### Changed
- readme: updated to include travis build badge.
- argon2: Updates SecureZeroMemory to match upstream for better performance.
- deps: Updated to support dep v0.5.0
- readme: Update with new benchmark stats due to SecureZeroMemory performance tweak.

### Fixed
- readme: Fixes example pathing.

## [v0.1.1] - 2018-06-14
### Changed
- deps: unpins golang.org/x/crypto from a specific revision.

## [v0.1.0] - 2018-05-30
### Added
- Initial Commit

### Fixed
- git: Fixes repo github links
- readme: Fix example link

[Unreleased]: https://github.com/matthewhartstonge/argon2/tree/master
[v0.3.2]: https://github.com/matthewhartstonge/argon2/tree/v0.3.2
[v0.3.1]: https://github.com/matthewhartstonge/argon2/tree/v0.3.1
[v0.3.0]: https://github.com/matthewhartstonge/argon2/tree/v0.3.0
[v0.2.1]: https://github.com/matthewhartstonge/argon2/tree/v0.2.1
[v0.2.0]: https://github.com/matthewhartstonge/argon2/tree/v0.2.0
[v0.1.5]: https://github.com/matthewhartstonge/argon2/tree/v0.1.5
[v0.1.4]: https://github.com/matthewhartstonge/argon2/tree/v0.1.4
[v0.1.3]: https://github.com/matthewhartstonge/argon2/tree/v0.1.3
[v0.1.2]: https://github.com/matthewhartstonge/argon2/tree/v0.1.2
[v0.1.1]: https://github.com/matthewhartstonge/argon2/tree/v0.1.1
[v0.1.0]: https://github.com/matthewhartstonge/argon2/tree/v0.1.0
