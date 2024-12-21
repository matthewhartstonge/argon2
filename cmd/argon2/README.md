# argon2-cli

## Usage

```shell
NAME:
	argon2 - An Argon2id CLI hash generator

USAGE:
	argon2 [command options] p@ssw0rd

VERSION:
	v0.1.3

AUTHOR:
	Matthew Hartstonge - https://github.com/matthewhartstonge

OPTIONS:
  -m uint
    	memory cost specifies the amount of memory to use in kibibytes.
  -p uint
    	parallelism cost specifies the number of parallel threads to spawn.
  -t uint
    	time cost specifies the number of iterations of argon2.
  -hash-len uint
    	hash length specifies the length of the resulting hash in bytes.
  -salt-len uint
    	salt length specifies the length of the resulting salt in bytes.
    	
GLOBAL OPTIONS:
  -h	displays usage.
  -s	silent removes all cli output.
```
