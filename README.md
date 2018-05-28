# ✨ argon2 ✨

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
