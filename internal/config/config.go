package config

import (
	"crypto/tls"
	"time"
)

// Config is a struct holds the client configuration options
type Config struct {
	RelyingPartyName string
	RelyingPartyUUID string
	CertificateLevel string
	HashType         string
	InteractionType  string
	DisplayText60    string
	DisplayText200   string
	URL              string
	Timeout          time.Duration
	TLSConfig        *tls.Config
}
