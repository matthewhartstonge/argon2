# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
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
