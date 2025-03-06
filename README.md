# 🔐 argon2

[![Go Reference](https://pkg.go.dev/badge/github.com/matthewhartstonge/argon2.svg)](https://pkg.go.dev/github.com/matthewhartstonge/argon2) [![Go Report Card](https://goreportcard.com/badge/github.com/matthewhartstonge/argon2)](https://goreportcard.com/report/github.com/matthewhartstonge/argon2) [![go](https://github.com/matthewhartstonge/argon2/actions/workflows/go.yml/badge.svg)](https://github.com/matthewhartstonge/argon2/actions/workflows/go.yml)

argon2 provides a pure Go implementation for Argon2 password hashing.

Intended to be a drop in replacement for lhecker's [argon2](https://github.com/lhecker/argon2)
library.

## tl;dr
```go
package main

import (
    "fmt"

    "github.com/matthewhartstonge/argon2"
)


func main() {
    argon := argon2.DefaultConfig()

    // Waaahht??! It includes magic salt generation for me ! Yasss...
    encoded, err := argon.HashEncoded([]byte("p@ssw0rd"))
    if err != nil {
        panic(err) // 💥
    }
    
    fmt.Println(string(encoded))
    // > $argon2id$v=19$m=65536,t=1,p=4$WXJGqwIB2qd+pRmxMOw9Dg$X4gvR0ZB2DtQoN8vOnJPR2SeFdUhH9TyVzfV98sfWeE

    ok, err := argon2.VerifyEncoded([]byte("p@ssw0rd"), encoded)
    if err != nil {
        panic(err) // 💥
    }
    
    matches := "no 🔒"
    if ok {
        matches = "yes 🔓"
    }
    fmt.Printf("Password Matches: %s\n", matches)
}
```

## Example
For a fuller example check out [_example/example.go](./_example/example.go) for 
a step-by-step introduction.

```
go run _example/example.go
```

## Limitations
* `Config.Parallelism` is a `uint8` instead of `uint32` as required by the
    underlying crypto library
* The [crypto](https://golang.org/x/crypto/argon2) implementation does not 
    support generation using Argon2d. Argon2id is now generally recommended.
* Errors still need to be properly implemented at the Go end 
    * This is mainly a case of implementing the PHC/Argon2 C++ pre-hash validation checks.

👌

## Benchmarks

The following manual benchmark was performed on a `i7-7700 @ 3.60GHz` with 
`AData DDR4 2132MHz` memory.

Note: 
- The native benchmarks are in a separate branch for reference in order to keep
  go mod dependencies tidy.

```
goos: windows
goarch: amd64
pkg: github.com/matthewhartstonge/argon2
BenchmarkHash
BenchmarkHash-8                               50          23479754 ns/op
BenchmarkNativeArgonBindingsHash
BenchmarkNativeArgonBindingsHash-8            38          31814984 ns/op
BenchmarkVerify
BenchmarkVerify-8                             49          22755661 ns/op
BenchmarkNativeArgonBindingsVerify
BenchmarkNativeArgonBindingsVerify-8          38          32342853 ns/op
BenchmarkEncode
BenchmarkEncode-8                        8693931               142 ns/op         604.42 MB/s
BenchmarkDecode
BenchmarkDecode-8                        5852835               208 ns/op         414.27 MB/s
BenchmarkSecureZeroMemory16
BenchmarkSecureZeroMemory16-8          289155232              4.09 ns/op        3914.11 MB/s
BenchmarkSecureZeroMemory64
BenchmarkSecureZeroMemory64-8          284339167              4.21 ns/op       15215.14 MB/s
BenchmarkSecureZeroMemory256
BenchmarkSecureZeroMemory256-8         160209489              7.51 ns/op       34093.00 MB/s
BenchmarkSecureZeroMemory1024
BenchmarkSecureZeroMemory1024-8         75083060              16.4 ns/op       62457.50 MB/s
BenchmarkSecureZeroMemory4096
BenchmarkSecureZeroMemory4096-8         20678958              60.3 ns/op       67922.23 MB/s
BenchmarkSecureZeroMemory1048576
BenchmarkSecureZeroMemory1048576-8         52404             22442 ns/op       46724.63 MB/s
PASS
ok      github.com/matthewhartstonge/argon2     18.481s
```

## Versioning Strategy

The API is stable and has been running in production for many years now, therefore won't be changing.

This library has a single dependency on `golang.org/x/crypto`. This means that as the version of Go is updated there, this library will roll up it's version to a new minor.
Any CVEs/security patches that come through via dependabot, without a resulting Go version update, will become a patch release.

As of 2025, `golang.org/x` libraries only support `1.(N-1).0` as their versioning strategy. If you're interested you can read about the implemented [automate go directive maintenance in golang.org/x repositories](https://go.googlesource.com/proposal/+/master/design/69095-x-repo-continuous-go.md) proposal. 

For example:

- If a version of `/x/crypto` now requires `go@1.28.0`, `argon2` will go from `v1.0.0` -> `v1.1.0`.
- If `x/crypto` releases a version that resolves CVEs with no requirement to upgrade Go, then `argon2` will go from `v1.0.0` -> `v1.0.1`.