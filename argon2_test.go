// Copyright (c) 2016 Leonard Hecker
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package argon2

import (
	"bytes"
	"reflect"
	"testing"
)

var (
	CFG     = DefaultConfig()
	PWD     = []byte("password")
	SALT    = []byte("saltsalt")
	HASH    = []byte{0x96, 0x5b, 0xd4, 0x76, 0xaa, 0x7a, 0xf7, 0x2d, 0x91, 0x07, 0xad, 0xbd, 0x74, 0x2b, 0x86, 0xe3, 0x69, 0x11, 0xe7, 0x2f, 0x8e, 0x71, 0xcf, 0xf3, 0x88, 0xa5, 0x79, 0x92, 0x7d, 0xeb, 0x48, 0xe3}
	ENCODED = []byte("$argon2i$v=19$m=4096,t=3,p=1$c2FsdHNhbHQ$llvUdqp69y2RB629dCuG42kR5y+Occ/ziKV5kn3rSOM")
)

func isFalsey(obj interface{}) bool {
	if obj == nil {
		return true
	}

	value := reflect.ValueOf(obj)
	kind := value.Kind()

	return kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil()
}

func mustBeFalsey(t *testing.T, name string, obj interface{}) {
	if !isFalsey(obj) {
		t.Errorf("'%s' should be nil, but is: %v", name, obj)
	}
}

func mustBeTruthy(t *testing.T, name string, obj interface{}) {
	if isFalsey(obj) {
		t.Errorf("'%s' should be non nil, but is: %v", name, obj)
	}
}

func TestHashRaw(t *testing.T) {
	r, err := CFG.HashRaw(PWD)
	mustBeTruthy(t, "r.Config", r.Config)
	mustBeTruthy(t, "r.Salt", r.Salt)
	mustBeTruthy(t, "r.Hash", r.Hash)
	mustBeFalsey(t, "err", err)
}

func TestHashEncoded(t *testing.T) {
	enc, err := CFG.HashEncoded(PWD)
	mustBeTruthy(t, "encoded", enc)
	mustBeFalsey(t, "err", err)

	if len(enc) == 0 {
		t.Error("len(encoded) must be > 0")
	}

	for _, b := range enc {
		if b == 0 {
			t.Error("encoded must not contain 0x00")
		}
	}
}

func TestHashWithSalt(t *testing.T) {
	r, err := CFG.Hash(PWD, SALT)
	mustBeTruthy(t, "r.Config", r.Config)
	mustBeTruthy(t, "r.Salt", r.Salt)
	mustBeTruthy(t, "r.Hash", r.Hash)
	mustBeFalsey(t, "err", err)

	if !bytes.Equal(r.Hash, HASH) {
		t.Logf("ref: %v", HASH)
		t.Logf("act: %v", r.Hash)
		t.Error("hashes do not match")
	}

	enc := r.Encode()
	mustBeTruthy(t, "encoded", enc)

	if !bytes.Equal(enc, ENCODED) {
		t.Logf("ref: %s", string(ENCODED))
		t.Logf("act: %s", string(enc))
		t.Error("encoded strings do not match")
	}
}

func TestVerifyRaw(t *testing.T) {
	r, err := CFG.HashRaw(PWD)
	mustBeTruthy(t, "r.Config", r.Config)
	mustBeTruthy(t, "r.Salt", r.Salt)
	mustBeTruthy(t, "r.Hash", r.Hash)
	mustBeFalsey(t, "err1", err)

	ok, err := r.Verify(PWD)
	mustBeTruthy(t, "ok", ok)
	mustBeFalsey(t, "err2", err)
}

func TestVerifyEncoded(t *testing.T) {
	encoded, err := CFG.HashEncoded(PWD)
	mustBeTruthy(t, "encoded", encoded)
	mustBeFalsey(t, "err1", err)

	ok, err := VerifyEncoded(PWD, encoded)
	mustBeTruthy(t, "ok", ok)
	mustBeFalsey(t, "err2", err)
}

func TestSecureZeroMemory(t *testing.T) {
	pwd := append(make([]byte, 0, len(PWD)), PWD...)

	// SecureZeroMemory should erase up to cap(pwd) --> let's test that too
	SecureZeroMemory(pwd[0:0])

	for _, b := range pwd {
		if b != 0 {
			t.Error("pwd must only contain 0x00")
		}
	}
}

func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = CFG.Hash(PWD, SALT)
	}
}

func BenchmarkVerify(b *testing.B) {
	r, err := CFG.Hash(PWD, SALT)
	if err != nil {
		b.Error(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = r.Verify(PWD)
	}
}

func BenchmarkEncode(b *testing.B) {
	r, err := CFG.Hash(PWD, SALT)
	if err != nil {
		b.Error(err)
	}

	b.SetBytes(int64(len(ENCODED)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = r.Encode()
	}
}

func BenchmarkDecode(b *testing.B) {
	b.SetBytes(int64(len(ENCODED)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = Decode(ENCODED)
	}
}

func BenchmarkSecureZeroMemory16(b *testing.B) {
	const numBytes int = 16
	buf := make([]byte, numBytes)
	b.SetBytes(int64(numBytes))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SecureZeroMemory(buf)
	}
}

func BenchmarkSecureZeroMemory64(b *testing.B) {
	const numBytes int = 64
	buf := make([]byte, numBytes)
	b.SetBytes(int64(numBytes))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SecureZeroMemory(buf)
	}
}

func BenchmarkSecureZeroMemory256(b *testing.B) {
	const numBytes int = 256
	buf := make([]byte, numBytes)
	b.SetBytes(int64(numBytes))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SecureZeroMemory(buf)
	}
}

func BenchmarkSecureZeroMemory1024(b *testing.B) {
	const numBytes int = 1024
	buf := make([]byte, numBytes)
	b.SetBytes(int64(numBytes))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SecureZeroMemory(buf)
	}
}

func BenchmarkSecureZeroMemory4096(b *testing.B) {
	const numBytes int = 4096
	buf := make([]byte, numBytes)
	b.SetBytes(int64(numBytes))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SecureZeroMemory(buf)
	}
}

func BenchmarkSecureZeroMemory1048576(b *testing.B) {
	const numBytes int = 1048576
	buf := make([]byte, numBytes)
	b.SetBytes(int64(numBytes))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SecureZeroMemory(buf)
	}
}
