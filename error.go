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
)

const (
	ARGON2_MIN_TIME = uint32(1)
	ARGON2_MAX_TIME = uint32(4294967295)
)
