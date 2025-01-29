package models

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

type SessionResponse struct {
	State               string      `json:"state"`
	Result              Result      `json:"result"`
	Signature           Signature   `json:"signature"`
	Cert                Certificate `json:"cert"`
	InteractionFlowUsed string      `json:"interactionFlowUsed"`
}
