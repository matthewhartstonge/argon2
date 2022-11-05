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

package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/matthewhartstonge/argon2"
)

var password = []byte("password")

func main() {
	// First we need a argon2.Config instance.
	// It contains all essential settings for Argon2.
	cfg := argon2.DefaultConfig()

	fmt.Printf("HashLength:  %d\n", cfg.HashLength)
	fmt.Printf("SaltLength:  %d\n", cfg.SaltLength)
	fmt.Printf("TimeCost:    %d\n", cfg.TimeCost)
	fmt.Printf("MemoryCost:  %d MB\n", cfg.MemoryCost/1024)
	fmt.Printf("Parallelism: %d\n", cfg.Parallelism)
	fmt.Printf("Mode:        %s\n", cfg.Mode)
	fmt.Printf("Version:     %s\n", cfg.Version)
	fmt.Println()

	// NOTE The following can be shortened by writing for instance:
	//
	//   encoded, err := cfg.HashEncoded(password)
	//

	// Let's hash the password "password"! We pass `nil` as the second
	// argument to make argon2 generate a salt for us.
	raw, err := cfg.Hash(password, nil)
	if err != nil {
		fmt.Printf("Error during hashing: %s\n", err.Error())
		os.Exit(1)
	}

	// The Raw struct contains 3 members:
	//   - A reference to the Config which created it
	//   - The Hash
	//   - And the Salt
	// It also has a method to turn it into a encoded string.
	fmt.Printf("Hash:        %s\n", hex.EncodeToString(raw.Hash))
	fmt.Printf("Salt:        %s\n", hex.EncodeToString(raw.Salt))
	fmt.Printf("Encoded:     %s\n", string(raw.Encode()))
	fmt.Println()

	// Let's say you're storing the encoded string representation in a database
	// and you pulled the following string `encoded` from there (instead of us
	// creating it here) and you want to verify it against an unencrypted user
	// password for login authentication.
	encoded := raw.Encode()

	// NOTE The following can be shortened by writing for instance:
	//
	//   ok, err := argon2.VerifyEncoded(password, encoded)
	//

	// We can then proceed to first decode the string back into a Raw struct.
	raw, err = argon2.Decode(encoded)
	if err != nil {
		fmt.Printf("Error during decoding: %s\n", err.Error())
		os.Exit(1)
	}

	// The Raw struct then allows us to verify it against a unencrypted password.
	ok, err := raw.Verify(password)
	if err != nil {
		fmt.Printf("Error during verification: %s\n", err.Error())
		os.Exit(1)
	}

	matches := "no"
	if ok {
		matches = "yes"
	}
	fmt.Printf("Matches:     %s\n", matches)
}
