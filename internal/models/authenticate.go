package models

type AllowedInteraction struct {
	Type          string `json:"type"`
	DisplayText60 string `json:"displayText60"`
}

type AuthenticateRequest struct {
	RelyingPartyName         string               `json:"relyingPartyName"`
	RelyingPartyUUID         string               `json:"relyingPartyUUID"`
	NationalIdentityNumber   string               `json:"nationalIdentityNumber"`
	CertificateLevel         string               `json:"certificateLevel"`
	Hash                     string               `json:"hash"`
	HashType                 string               `json:"hashType"`
	AllowedInteractionsOrder []AllowedInteraction `json:"allowedInteractionsOrder"`
}

type AuthenticateResponse struct {
	SessionID string `json:"sessionID"`
	Code      string `json:"code"`
}
