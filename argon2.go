/*
 * Copyright 2022. Matthew Hartstonge <matt@mykro.co.nz>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package argon2 implements the key derivation function Argon2.
//
// Argon2 was selected as the winner of the Password Hashing Competition and
// can be used to derive cryptographic keys from passwords.
package argon2

import (
	"crypto/rand"
	"crypto/subtle"
	"golang.org/x/crypto/argon2"
)

// Mode exists for type check purposes. See Config.
type Mode uint32

const (
	// modeArgon2d is added here for completeness, but is not exposed as the
	// x/crypto library does not expose argon2d as an option.
	//
	// modeArgon2d is faster and uses data-depending memory access,
	// which makes it highly resistant against GPU cracking attacks and
	// suitable for applications with no (!) threats from
	// side-channel timing attacks (eg. cryptocurrencies).
	modeArgon2d = iota

	// ModeArgon2i uses data-independent memory access, which is
	// preferred for password hashing and password-based key derivation
	// (e.g. hard drive encryption), but it's slower as it makes
	// more passes over the memory to protect from TMTO attacks.
	ModeArgon2i

	// ModeArgon2id is a hybrid of Argon2i and Argon2d, using a
	// combination of data-depending and data-independent memory accesses,
	// which gives some of Argon2i's resistance to side-channel cache timing
	// attacks and much of Argon2d's resistance to GPU cracking attacks.
	ModeArgon2id
)

// String simply maps a ModeArgon{d,i,id} constant to a "Argon{d,i,id}" string
// or returns "unknown" if `m` does not match one of the constants.
func (m Mode) String() string {
	switch m {
	case modeArgon2d:
		return "Argon2d"
	case ModeArgon2i:
		return "Argon2i"
	case ModeArgon2id:
		return "Argon2id"
	default:
		return "unknown"
	}
}

// Version exists for type check purposes. See Config.
type Version uint32

const (
	// Version10 of the Argon2 algorithm. Deprecated: Use Version13 instead.
	Version10 = 0x10

	// Version13 of the Argon2 algorithm. Recommended.
	Version13 = 0x13
)

// String simply maps a Version{10,13} constant to a "{10,13}" string
// or returns "unknown" if `v` does not match one of the constants.
func (v Version) String() string {
	switch v {
	case Version10:
		return "10"
	case Version13:
		return "13"
	default:
		return "unknown"
	}
}

// Config contains all configuration parameters for the Argon2 hash function.
type Config struct {
	// HashLength specifies the length of the resulting hash in Bytes.
	//
	// Must be > 0.
	HashLength uint32

	// SaltLength specifies the length of the resulting salt in Bytes,
	// if one of the helper methods is used.
	//
	// Must be > 0.
	SaltLength uint32

	// TimeCost specifies the number of iterations of argon2.
	//
	// Must be > 0.
	// If you use ModeArgon2i this should *always* be >= 3 due to TMTO attacks.
	// Additionally if you can afford it you might set it to >= 10.
	TimeCost uint32

	// MemoryCost specifies the amount of memory to use in Kibibytes.
	//
	// Must be > 0.
	MemoryCost uint32

	// Parallelism specifies the amount of threads to use.
	//
	// Must be > 0.
	Parallelism uint8

	// Mode specifies the hashing method used by argon2.
	//
	// If you're writing a server and unsure what to choose,
	// use ModeArgon2i with a TimeCost >= 3.
	Mode Mode

	// Version specifies the argon2 version to be used.
	Version Version
}

// DefaultConfig returns a Config struct suitable for most servers. These
// default settings are based on RFC9106 recommendations.
//
// Refer:
//   - https://datatracker.ietf.org/doc/html/rfc9106#section-7.4
//   - https://datatracker.ietf.org/doc/html/rfc9106#section-4
//
// The memory constrained settings result in around 50ms of computation time
// while using 64 MiB of memory during hashing. Tested on an Intel Core i7-7700
// @ 3.6 GHz with DDR4 @ 2133 MHz.
func DefaultConfig() Config {
	return MemoryConstrainedDefaults()
}

// RecommendedDefaults provides configuration based on the first recommended
// option as described in RFC9106.
//
// If a uniformly safe option that is not tailored to your application or
// hardware is acceptable, select Argon2id with t=1 iteration, p=4 lanes,
// m=2^(21) (2 GiB of RAM), 128-bit salt, and 256-bit tag size.
func RecommendedDefaults() Config {
	return Config{
		HashLength:  32, // 32 * 8 = 256-bits
		SaltLength:  16, // 16 * 8 = 128-bits
		TimeCost:    1,
		MemoryCost:  2 * 1024 * 1024, // 2^(21) (2 GiB of RAM)
		Parallelism: 4,
		Mode:        ModeArgon2id,
		Version:     Version13,
	}
}

// MemoryConstrainedDefaults provides configuration based on the second
// recommended option as described in RFC9106.
//
// If much less memory is available, a uniformly safe option is Argon2id with
// t=3 iterations, p=4 lanes, m=2^(16) (64 MiB of RAM), 128-bit salt, and
// 256-bit tag size.
func MemoryConstrainedDefaults() Config {
	return Config{
		HashLength:  32, // 32 * 8 = 256-bits
		SaltLength:  16, // 16 * 8 = 128-bits
		TimeCost:    3,
		MemoryCost:  64 * 1024, // 2^(16) (64MiB of RAM)
		Parallelism: 4,
		Mode:        ModeArgon2id,
		Version:     Version13,
	}
}

// Hash takes a password and optionally a salt and returns an Argon2 hash.
//
// If salt is nil an appropriate salt of Config.SaltLength bytes is generated
// for you.
func (c *Config) Hash(pwd []byte, salt []byte) (Raw, error) {
	if pwd == nil {
		return Raw{}, ErrPwdTooShort
	}

	if salt == nil {
		salt = make([]byte, c.SaltLength)
		_, err := rand.Read(salt)

		if err != nil {
			return Raw{}, err
		}
	}

	hash := make([]byte, c.HashLength)
	switch c.Mode {
	case ModeArgon2i:
		hash = argon2.Key(pwd, salt, c.TimeCost, c.MemoryCost, c.Parallelism, c.HashLength)
	case ModeArgon2id:
		hash = argon2.IDKey(pwd, salt, c.TimeCost, c.MemoryCost, c.Parallelism, c.HashLength)
	case modeArgon2d:
		return Raw{}, ErrModeUnsupported
	}

	return Raw{
		Config: *c,
		Salt:   salt,
		Hash:   hash,
	}, nil
}

// HashRaw is a helper function around Hash()
// which automatically generates a salt for you.
func (c *Config) HashRaw(pwd []byte) (Raw, error) {
	return c.Hash(pwd, nil)
}

// HashEncoded is a helper function around Hash() which automatically
// generates a salt and encodes the result for you.
func (c *Config) HashEncoded(pwd []byte) (encoded []byte, err error) {
	r, err := c.Hash(pwd, nil)
	if err == nil {
		encoded = r.Encode()
	}
	return
}

// Raw wraps a salt and hash pair including the Config with which it was
// generated.
//
// A Raw struct is generated using Decode() or the Hash*() methods above.
//
// You MUST ensure that a Raw instance is not changed after creation,
// otherwise you risk race conditions. If you do need to change it during
// runtime use a Mutex and simply create a copy of your shared Raw
// instance in the critical section and store it on your local stack.
// That way your critical section is very short, while allowing you to safely
// call all the member methods on your local "immutable" copy.
type Raw struct {
	Config Config
	Salt   []byte
	Hash   []byte
}

// Verify returns true if `pwd` matches the hash in `raw` and otherwise false.
func (raw *Raw) Verify(pwd []byte) (bool, error) {
	r, err := raw.Config.Hash(pwd, raw.Salt)
	if err != nil {
		return false, err
	}
	return subtle.ConstantTimeCompare(r.Hash, raw.Hash) == 1, nil
}

// VerifyEncoded returns true if `pwd` matches the encoded hash `encoded` and
// otherwise false.
//
// Note: Only supports verifying `argon2i` and `argon2id` hashes. As `x/crypto`
// doesn't expose the generation of `argon2d` hashing, we can't generate an
// `argon2d` hash to verify the incoming password against.
//
// Refer: https://github.com/matthewhartstonge/argon2/issues/80#issuecomment-2679636640
func VerifyEncoded(pwd []byte, encoded []byte) (bool, error) {
	r, err := Decode(encoded)
	if err != nil {
		return false, err
	}
	return r.Verify(pwd)
}

// SecureZeroMemory is a helper method which sets all bytes in `b`
// (up to its capacity) to `0x00`, erasing its contents.
func SecureZeroMemory(b []byte) {
	b = b[:cap(b):cap(b)]
	for i := range b {
		b[i] = 0
	}
}
