package errors

import "errors"

var (
	ErrMissingRelyingPartyName = errors.New("missing required configuration: RelyingPartyName")
	ErrMissingRelyingPartyUUID = errors.New("missing required configuration: RelyingPartyUUID")

	ErrInvalidNationalIdentityNumber = errors.New("invalid NationalIdentityNumber format")

	ErrSmartIdProviderError = errors.New("Smart-ID provider error")
)
