# argon2-cli

## Usage

```shell
Usage:
  argon2 [options] p@ssw0rd

Description:
  Enables hashing passwords with argon via the CLI

Options:
  -hash-len uint
    	hash length specifies the length of the resulting hash in bytes.
  -salt-len uint
    	salt length specifies the length of the resulting salt in bytes.
  -m uint
    	memory cost specifies the amount of memory to use in kibibytes.
  -p uint
    	parallelism cost
  -t uint
    	time cost specifies the number of iterations of argon2.
  -s
    	silent removes all cli output
  -help
    	prints usage
```
