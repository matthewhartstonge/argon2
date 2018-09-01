# ✨ argon2 ✨

[![Build Status](https://travis-ci.org/matthewhartstonge/argon2.svg?branch=master)](https://travis-ci.org/matthewhartstonge/argon2)

This provides a Pure Go implemention for Argon2 password hashing. It is 
intended to be an _almost_ drop in replacement for lhecker's amazing 
[argon2](https://github.com/lhecker/argon2) library.

## Usage
See [_example/example.go](./_example/example.go) for a simple introduction and 
try it out with:

```
go run examples/example.go
```

## Limitations
* `Config.Parallelism` is now a `uint8` instead of `uint32` as required by the 
    underlying crypto library
* The [crypto](https://golang.org/x/crypto/argon2) implementation does not 
    support generation using Argon2d
* Errors still need to be properly implemented at the Go end 
    * This is mainly a case of implementing the PHC/Argon2 C++ pre-hash validation checks.

👌

## Benchmarks

Benchmarks are run on each build. Follow the build badge link to check the 
latest status on benchmarking results via TravisCI.

[![Build Status](https://travis-ci.org/matthewhartstonge/argon2.svg?branch=master)](https://travis-ci.org/matthewhartstonge/argon2) 

The following manual benchmark was performed on a `i7-4770 @ 3.40GHz` with 
`AData DDR3 1333MHz` memory.

```
goos: windows
goarch: amd64
goversion: go version go1.10.3 windows/amd64
pkg: github.com/matthewhartstonge/argon2
BenchmarkHash-8                        	      50	  27799898 ns/op
BenchmarkNativeArgonBindingsHash-8     	      30	  42499980 ns/op
BenchmarkVerify-8                      	      50	  27499686 ns/op
BenchmarkNativeArgonBindingsVerify-8   	      30	  41932776 ns/op
BenchmarkEncode-8                      	10000000	       165 ns/op	 519.01 MB/s
BenchmarkDecode-8                      	 5000000	       266 ns/op	 323.06 MB/s
BenchmarkSecureZeroMemory16-8          	100000000	        18.9 ns/op	 847.90 MB/s
BenchmarkSecureZeroMemory64-8          	50000000	        27.0 ns/op	2370.40 MB/s
BenchmarkSecureZeroMemory256-8         	50000000	        37.1 ns/op	6892.79 MB/s
BenchmarkSecureZeroMemory1024-8        	20000000	        63.0 ns/op	16266.87 MB/s
BenchmarkSecureZeroMemory4096-8        	20000000	       112 ns/op	36360.42 MB/s
BenchmarkSecureZeroMemory1048576-8     	   30000	     57833 ns/op	18131.03 MB/s
PASS
```

_Disclaimer:_ I was unable to set the GCC build optimisation flags as the 
compilation kept saying:

```
# runtime/cgo
gcc: error: "-O3: Invalid argument

# - and - 
# runtime/cgo
gcc: error: "-march=native": Invalid argument
```

Could be a windows GCC incompatibility? 