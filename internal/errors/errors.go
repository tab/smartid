package errors

import (
	"errors"
)

var (
	ErrMissingRelyingPartyName = errors.New("missing required configuration: RelyingPartyName")
	ErrMissingRelyingPartyUUID = errors.New("missing required configuration: RelyingPartyUUID")

	ErrUnsupportedHashType = errors.New("unsupported hash type, allowed hash types are SHA256, SHA384 or SHA512")

	ErrSmartIdProviderError   = errors.New("Smart-ID provider error")
	ErrSmartIdSessionNotFound = errors.New("Smart-ID session not found or expired")

	ErrInvalidCertificate    = errors.New("invalid certificate")
	ErrInvalidIdentityNumber = errors.New("invalid identity number")

	ErrFailedToGenerateRandomBytes = errors.New("failed to generate random bytes")

	ErrUnsupportedState  = errors.New("unsupported state, allowed states are COMPLETE or RUNNING")
	ErrUnsupportedResult = errors.New("unsupported result, allowed results are OK or USER_REFUSED, USER_REFUSED_DISPLAYTEXTANDPIN, USER_REFUSED_VC_CHOICE, USER_REFUSED_CONFIRMATIONMESSAGE, USER_REFUSED_CONFIRMATIONMESSAGE_WITH_VC_CHOICE, USER_REFUSED_CERT_CHOICE, WRONG_VC, TIMEOUT")

	ErrAuthenticationIsRunning = errors.New("authentication is still running")

	ErrFailedToDecodeCertificate = errors.New("failed to decode certificate")
	ErrFailedToParseCertificate  = errors.New("failed to parse certificate")
)
