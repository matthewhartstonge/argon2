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

package argon2

// Error represents the error code returned by argon2.
type Error string

func (e Error) Error() string {
	return string(e)
}

var (
	ErrOutputPtrNull         = Error("output pointer is null")
	ErrOutputTooShort        = Error("output is too short")
	ErrOutputTooLong         = Error("output is too long")
	ErrPwdTooShort           = Error("password is too short")
	ErrPwdTooLong            = Error("password is too long")
	ErrSaltTooShort          = Error("salt is too short")
	ErrSaltTooLong           = Error("salt is too long")
	ErrAdTooShort            = Error("associated data is too short")
	ErrAdTooLong             = Error("associated data is too long")
	ErrSecretTooShort        = Error("secret is too short")
	ErrSecretTooLong         = Error("secret is too long")
	ErrTimeTooSmall          = Error("time cost is too small")
	ErrTimeTooLarge          = Error("time cost is too large")
	ErrMemoryTooLittle       = Error("memory cost is too small")
	ErrMemoryTooMuch         = Error("memory cost is too large")
	ErrLanesTooFew           = Error("too few lanes")
	ErrLanesTooMany          = Error("too many lanes")
	ErrPwdPtrMismatch        = Error("password pointer is null, but password length is not 0")
	ErrSaltPtrMismatch       = Error("salt pointer is null, but salt length is not 0")
	ErrSecretPtrMismatch     = Error("secret pointer is null, but secret length is not 0")
	ErrAdPtrMismatch         = Error("associated data pointer is null, but ad length is not 0")
	ErrMemoryAllocationError = Error("memory allocation error")
	ErrFreeMemoryCbkNull     = Error("the free memory callback is null")
	ErrAllocateMemoryCbkNull = Error("the allocate memory callback is null")
	ErrIncorrectParameter    = Error("argon2_context context is null")
	ErrIncorrectType         = Error("there is no such version of argon2")
	ErrOutPtrMismatch        = Error("output pointer mismatch")
	ErrThreadsTooFew         = Error("not enough threads")
	ErrThreadsTooMany        = Error("too many threads")
	ErrMissingArgs           = Error("missing arguments")
	ErrEncodingFail          = Error("encoding failed")
	ErrDecodingFail          = Error("decoding failed")
	ErrThreadFail            = Error("threading failure")
	ErrDecodingLengthFail    = Error("some of encoded parameters are too long or too short")
	ErrVerifyMismatch        = Error("the password does not match the supplied hash")
	ErrModeUnsupported       = Error("argon2d hashing mode unsupported by go maintainers")
)

const (
	ARGON2_MIN_TIME = uint32(1)
	ARGON2_MAX_TIME = uint32(4294967295)
)
