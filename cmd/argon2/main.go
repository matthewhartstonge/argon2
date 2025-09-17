package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"os"

	"github.com/matthewhartstonge/argon2"
)

var (
	// AppName configures the binaries name.
	AppName = "argon2"
	// AppVersion outputs the binaries version.
	AppVersion = "1.4.0" // x-release-please-version
	// AppCommit specifies the exact git commit the binary has been built from.
	AppCommit = "unknown"
	// AppCommitDate specifies the date the commit was made.
	AppCommitDate = ""
)

var (
	// configure flags
	t       = flag.Uint("t", 0, "time cost specifies the number of iterations of argon2.")
	m       = flag.Uint("m", 0, "memory cost specifies the amount of memory to use in kibibytes.")
	p       = flag.Uint("p", 0, "parallelism cost specifies the number of parallel threads to spawn.")
	hashLen = flag.Uint("hash-len", 0, "hash length specifies the length of the resulting hash in bytes.")
	saltLen = flag.Uint("salt-len", 0, "salt length specifies the length of the resulting salt in bytes.")
	s       = flag.Bool("s", false, "silent removes all cli output.")
)

type config struct {
	silent bool
	argon  argon2.Config
}

func main() {
	setupFlagUsage()

	cfg, err := parseFlags()
	if err != nil {
		cliPrintln(cfg, err)
		os.Exit(1)
	}

	pw, err := parsePassword()
	if err != nil {
		cliPrintln(cfg, err)
		os.Exit(1)
	}

	out, err := hash(cfg, pw)
	if err != nil {
		cliPrintln(cfg, err)
		os.Exit(1)
	}

	fmt.Println(out)
}

// setupFlagUsage configures the binaries custom usage signature.
func setupFlagUsage() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "NAME:\n\t%s - An Argon2id CLI hash generator\n\n", AppName)
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "USAGE:\n\t%s [command options] p@ssw0rd\n\n", AppName)
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "VERSION:\n\tv%s (%s) %s\n\n", AppVersion, AppCommit, AppCommitDate)
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "AUTHOR:\n\tMatthew Hartstonge - https://github.com/matthewhartstonge\n\n")
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "OPTIONS:\n")
		flag.PrintDefaults()
	}
}

// parseFlags extracts, validates and configures argon2 by the based on the
// set flag options.
func parseFlags() (*config, error) {
	flag.Parse()

	cfg := &config{
		silent: *s,
	}

	argon := argon2.RecommendedDefaults()
	if *hashLen != 0 {
		v, err := isUint32(*hashLen)
		if err != nil {
			return cfg, err
		}
		argon.HashLength = v
	}

	if *saltLen != 0 {
		v, err := isUint32(*saltLen)
		if err != nil {
			return cfg, err
		}
		argon.SaltLength = v
	}

	if *t != 0 {
		v, err := isUint32(*t)
		if err != nil {
			return cfg, err
		}
		argon.TimeCost = v
	}

	if *m != 0 {
		v, err := isUint32(*m)
		if err != nil {
			return cfg, err
		}
		argon.MemoryCost = v
	}

	if *p != 0 {
		v, err := isUint8(*p)
		if err != nil {
			return cfg, err
		}
		argon.Parallelism = v
	}

	// inject argon config
	cfg.argon = argon

	return cfg, nil
}

// parsePassword extracts and returns the password to hash from args.
func parsePassword() (string, error) {
	args := flag.Args()
	if len(args) == 0 {
		return "", errors.New("please provide a password to hash")
	}

	return args[0], nil
}

// hash encodes and returns a stringified argon2 hash.
func hash(cfg *config, password string) (string, error) {
	cliPrintf(cfg,
		"Generating argon2id hash with m=%d, t=%d, p=%d4...\n\n",
		cfg.argon.MemoryCost,
		cfg.argon.TimeCost,
		cfg.argon.Parallelism,
	)

	enc, err := cfg.argon.HashEncoded([]byte(password))
	if err != nil {
		return "", err
	}

	return string(enc), nil
}

// isUint8 performs bounds checking for uint8
func isUint8(i uint) (uint8, error) {
	if i > math.MaxUint8 {
		return 0, fmt.Errorf("argon2: invalid uint8 value: %d", i)
	}

	return uint8(i), nil
}

// isUint32 performs bounds checking for uint32
func isUint32(i uint) (uint32, error) {
	if i > math.MaxUint32 {
		return 0, fmt.Errorf("argon2: invalid uint32 value: %d", i)
	}

	return uint32(i), nil
}

// cliPrintf provides a fmt.Printf wrapper that doesn't print if silence is
// desired.
func cliPrintf(cfg *config, format string, a ...any) {
	if cfg.silent {
		return
	}

	fmt.Printf(format, a...)
}

// cliPrintln provides a fmt.Println wrapper that doesn't print if silence is
// desired.
func cliPrintln(cfg *config, a ...any) {
	if cfg.silent {
		return
	}

	fmt.Println(a...)
}
