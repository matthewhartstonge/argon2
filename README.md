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
goversion: go version go1.11 windows/amd64
pkg: github.com/matthewhartstonge/argon2
BenchmarkHash-8                        	      50	  27979234 ns/op
BenchmarkNativeArgonBindingsHash-8     	      30	  42269670 ns/op
BenchmarkVerify-8                      	      50	  26799596 ns/op
BenchmarkNativeArgonBindingsVerify-8   	      30	  41367290 ns/op
BenchmarkEncode-8                      	10000000	       159 ns/op	 540.54 MB/s
BenchmarkDecode-8                      	 5000000	       265 ns/op	 324.04 MB/s
BenchmarkSecureZeroMemory16-8          	300000000	         4.50 ns/op	3557.50 MB/s
BenchmarkSecureZeroMemory64-8          	300000000	         5.20 ns/op	12307.78 MB/s
BenchmarkSecureZeroMemory256-8         	200000000	         9.29 ns/op	27550.88 MB/s
BenchmarkSecureZeroMemory1024-8        	100000000	        22.9 ns/op	44633.69 MB/s
BenchmarkSecureZeroMemory4096-8        	20000000	        77.0 ns/op	53194.84 MB/s
BenchmarkSecureZeroMemory1048576-8     	   50000	     39450 ns/op	26579.48 MB/s
PASS
```
