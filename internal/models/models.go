package models

type AuthenticationRequest struct {
	RelyingPartyName         string               `json:"relyingPartyName"`
	RelyingPartyUUID         string               `json:"relyingPartyUUID"`
	NationalIdentityNumber   string               `json:"nationalIdentityNumber"`
	CertificateLevel         string               `json:"certificateLevel"`
	Hash                     string               `json:"hash"`
	HashType                 string               `json:"hashType"`
	Nonce                    string               `json:"nonce,omitempty"`
	AllowedInteractionsOrder []AllowedInteraction `json:"allowedInteractionsOrder"`
}

type AllowedInteraction struct {
	Type           string `json:"type"`
	DisplayText60  string `json:"displayText60"`
	DisplayText200 string `json:"displayText200"`
}

type AuthenticationResponse struct {
	State               string      `json:"state"`
	Result              Result      `json:"result"`
	Signature           Signature   `json:"signature"`
	Cert                Certificate `json:"cert"`
	InteractionFlowUsed string      `json:"interactionFlowUsed"`
}

type Result struct {
	EndResult      string `json:"endResult"`
	DocumentNumber string `json:"documentNumber"`
}

type Signature struct {
	Value     string `json:"value"`
	Algorithm string `json:"algorithm"`
}

type Certificate struct {
	Value            string `json:"value"`
	CertificateLevel string `json:"certificateLevel"`
}
