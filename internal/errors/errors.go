package errors

import "errors"

var (
	ErrMissingRelyingPartyName = errors.New("missing required configuration: RelyingPartyName")
	ErrMissingRelyingPartyUUID = errors.New("missing required configuration: RelyingPartyUUID")

	ErrUnsupportedHashType = errors.New("unsupported hash type, allowed hash types are SHA256, SHA384 or SHA512")

	ErrSmartIdProviderError = errors.New("Smart-ID provider error")

	ErrInvalidCertificate    = errors.New("invalid certificate")
	ErrInvalidIdentityNumber = errors.New("invalid identity number")

	ErrFailedToGenerateRandomBytes = errors.New("failed to generate random bytes")

	ErrFailedToDecodeCertificate = errors.New("failed to decode certificate")
	ErrFailedToParseCertificate  = errors.New("failed to parse certificate")
)
