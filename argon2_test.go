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

package argon2_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/matthewhartstonge/argon2"
)

var (
	config = argon2.Config{
		HashLength:  32,
		SaltLength:  16,
		TimeCost:    1,
		MemoryCost:  32 * 1024,
		Parallelism: 1,
		Mode:        argon2.ModeArgon2id,
		Version:     argon2.Version13,
	}
	password        = []byte("password")
	salt            = []byte("saltsalt")
	expectedHash    = []byte{139, 118, 66, 92, 63, 17, 51, 11, 184, 106, 68, 37, 211, 16, 139, 244, 189, 217, 38, 53, 116, 148, 139, 173, 176, 3, 182, 239, 235, 210, 75, 155}
	expectedEncoded = []byte("$argon2id$v=19$m=32768,t=1,p=1$c2FsdHNhbHQ$i3ZCXD8RMwu4akQl0xCL9L3ZJjV0lIutsAO27+vSS5s")
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
	r, err := config.HashRaw(password)
	mustBeTruthy(t, "r.Config", r.Config)
	mustBeTruthy(t, "r.Salt", r.Salt)
	mustBeTruthy(t, "r.Hash", r.Hash)
	mustBeFalsey(t, "err", err)
}

func TestHashEncoded(t *testing.T) {
	enc, err := config.HashEncoded(password)
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
	r, err := config.Hash(password, salt)
	mustBeTruthy(t, "r.Config", r.Config)
	mustBeTruthy(t, "r.Salt", r.Salt)
	mustBeTruthy(t, "r.Hash", r.Hash)
	mustBeFalsey(t, "err", err)

	if !bytes.Equal(r.Hash, expectedHash) {
		t.Logf("ref: %v", expectedHash)
		t.Logf("act: %v", r.Hash)
		t.Error("hashes do not match")
	}

	enc := r.Encode()
	mustBeTruthy(t, "encoded", enc)

	if !bytes.Equal(enc, expectedEncoded) {
		t.Logf("ref: %s", string(expectedEncoded))
		t.Logf("act: %s", string(enc))
		t.Error("encoded strings do not match")
	}
}

func TestVerifyRaw(t *testing.T) {
	r, err := config.HashRaw(password)
	mustBeTruthy(t, "r.Config", r.Config)
	mustBeTruthy(t, "r.Salt", r.Salt)
	mustBeTruthy(t, "r.Hash", r.Hash)
	mustBeFalsey(t, "err1", err)

	ok, err := r.Verify(password)
	mustBeTruthy(t, "ok", ok)
	mustBeFalsey(t, "err2", err)
}

func TestVerifyEncoded(t *testing.T) {
	encoded, err := config.HashEncoded(password)
	mustBeTruthy(t, "encoded", encoded)
	mustBeFalsey(t, "err1", err)

	ok, err := argon2.VerifyEncoded(password, encoded)
	mustBeTruthy(t, "ok", ok)
	mustBeFalsey(t, "err2", err)
}

func TestSecureZeroMemory(t *testing.T) {
	pwd := append(make([]byte, 0, len(password)), password...)

	// SecureZeroMemory should erase up to cap(pwd) --> let's test that too
	argon2.SecureZeroMemory(pwd[0:0])

	for _, b := range pwd {
		if b != 0 {
			t.Error("pwd must only contain 0x00")
		}
	}
}

func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = config.Hash(password, salt)
	}
}

func BenchmarkMemoryConstrainedDefaults(b *testing.B) {
	cfg := argon2.MemoryConstrainedDefaults()
	for i := 0; i < b.N; i++ {
		_, _ = cfg.Hash(password, salt)
	}
}

func BenchmarkRecommendedDefaults(b *testing.B) {
	cfg := argon2.RecommendedDefaults()
	for i := 0; i < b.N; i++ {
		_, _ = cfg.Hash(password, salt)
	}
}

func BenchmarkVerify(b *testing.B) {
	r, err := config.Hash(password, salt)
	if err != nil {
		b.Error(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = r.Verify(password)
	}
}

func BenchmarkEncode(b *testing.B) {
	r, err := config.Hash(password, salt)
	if err != nil {
		b.Error(err)
	}

	b.SetBytes(int64(len(expectedEncoded)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r.Encode()
	}
}

func BenchmarkDecode(b *testing.B) {
	b.SetBytes(int64(len(expectedEncoded)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = argon2.Decode(expectedEncoded)
	}
}

func BenchmarkSecureZeroMemory16(b *testing.B) {
	const numBytes int = 16
	buf := make([]byte, numBytes)
	b.SetBytes(int64(numBytes))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		argon2.SecureZeroMemory(buf)
	}
}

func BenchmarkSecureZeroMemory64(b *testing.B) {
	const numBytes int = 64
	buf := make([]byte, numBytes)
	b.SetBytes(int64(numBytes))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		argon2.SecureZeroMemory(buf)
	}
}

func BenchmarkSecureZeroMemory256(b *testing.B) {
	const numBytes int = 256
	buf := make([]byte, numBytes)
	b.SetBytes(int64(numBytes))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		argon2.SecureZeroMemory(buf)
	}
}

func BenchmarkSecureZeroMemory1024(b *testing.B) {
	const numBytes int = 1024
	buf := make([]byte, numBytes)
	b.SetBytes(int64(numBytes))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		argon2.SecureZeroMemory(buf)
	}
}

func BenchmarkSecureZeroMemory4096(b *testing.B) {
	const numBytes int = 4096
	buf := make([]byte, numBytes)
	b.SetBytes(int64(numBytes))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		argon2.SecureZeroMemory(buf)
	}
}

func BenchmarkSecureZeroMemory1048576(b *testing.B) {
	const numBytes int = 1048576
	buf := make([]byte, numBytes)
	b.SetBytes(int64(numBytes))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		argon2.SecureZeroMemory(buf)
	}
}
