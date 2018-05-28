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
	// crypto does not expose argon2d as an option.
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

// DefaultConfig returns a Config struct suitable for most servers.
//
// These default settings are based on the draft RFC recommendations.
// See: https://tools.ietf.org/html/draft-irtf-cfrg-argon2-03#section-9.4
//
// These default settings result in around 20ms of computation time while using
// 64 MiB of memory.
// (Tested on an i7 4770 @ 3.4 GHz & AData PC3-12800 @ 1600 MHz).
func DefaultConfig() Config {
	return Config{
		HashLength: 32,
		// 128 bits is sufficient for all applications, but can be reduced to
		// 64 bits in the case of space constraints.
		SaltLength:  16,
		TimeCost:    1,
		MemoryCost:  64 * 1024,
		Parallelism: 4,
		Mode:        ModeArgon2id,
		Version:     Version13,
	}
}

// Hash takes a password and optionally a salt and returns an Argon2 hash.
//
// If salt is nil a appropriate salt of Config.SaltLength bytes is generated for you.
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

// Raw wraps a salt and hash pair including the Config with which it was generated.
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

// VerifyEncoded returns true if `pwd` matches the encoded hash `encoded` and otherwise false.
func VerifyEncoded(pwd []byte, encoded []byte) (bool, error) {
	r, err := Decode(encoded)
	if err != nil {
		return false, err
	}
	return r.Verify(pwd)
}

// SecureZeroMemory is a helper method which as securely as possible sets all
// bytes in `b` (up to it's capacity) to `0x00`, erasing it's contents.
//
// Using this method DOES NOT make secrets impossible to recover from memory,
// it's just a good start and generally recommended to use.
//
// Due to the nature of memory allocated by the Go runtime, SecureZeroMemory
// cannot guarantee that the data does not exist elsewhere in memory.
func SecureZeroMemory(b []byte) {
	c := cap(b)
	if c > 0 {
		b = b[:c:c]
		wipeBytes(b)
	}
}

// Wipes a byte slice with zeroes.
//
func wipeBytes(buf []byte) {
	if len(buf) == 0 {
		return
	}
	buf[0] = 0
	for bp := 1; bp < len(buf); bp *= 2 {
		copy(buf[bp:], buf[:bp])
	}
}
