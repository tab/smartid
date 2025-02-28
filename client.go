package smartid

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/tab/smartid/internal/config"
	"github.com/tab/smartid/internal/errors"
	"github.com/tab/smartid/internal/requests"
	"github.com/tab/smartid/internal/utils"
)

const (
	CertificateLevel = requests.CertificateLevelQUALIFIED
	InteractionType  = requests.InteractionTypeDisplayTextAndPIN
	DisplayText60    = "Enter PIN1"
	DisplayText200   = "Confirm the authentication request and enter PIN1"
	Timeout          = requests.Timeout
	URL              = "https://sid.demo.sk.ee/smart-id-rp/v2"
)

type Client interface {
	CreateSession(ctx context.Context, nationalIdentityNumber string) (*Session, error)
	FetchSession(ctx context.Context, sessionId string) (*Person, error)

	WithRelyingPartyName(name string) Client
	WithRelyingPartyUUID(id string) Client
	WithCertificateLevel(level string) Client
	WithHashType(hashType string) Client
	WithNonce(nonce string) Client
	WithInteractionType(interactionType string) Client
	WithDisplayText60(text string) Client
	WithDisplayText200(text string) Client
	WithURL(url string) Client
	WithTimeout(timeout time.Duration) Client
	WithTLSConfig(tlsConfig *tls.Config) Client

	Validate() error
}

type client struct {
	config *config.Config
}

func NewClient() Client {
	cfg := &config.Config{
		CertificateLevel: CertificateLevel,
		HashType:         utils.HashTypeSHA512,
		InteractionType:  InteractionType,
		DisplayText60:    DisplayText60,
		DisplayText200:   DisplayText200,
		URL:              URL,
		Timeout:          Timeout,
	}

	return &client{
		config: cfg,
	}
}

func (c *client) WithRelyingPartyName(name string) Client {
	c.config.RelyingPartyName = name
	return c
}

func (c *client) WithRelyingPartyUUID(id string) Client {
	c.config.RelyingPartyUUID = id
	return c
}

func (c *client) WithCertificateLevel(level string) Client {
	c.config.CertificateLevel = level
	return c
}

func (c *client) WithHashType(hashType string) Client {
	c.config.HashType = hashType
	return c
}

func (c *client) WithNonce(nonce string) Client {
	c.config.Nonce = nonce
	return c
}

func (c *client) WithInteractionType(interactionType string) Client {
	c.config.InteractionType = interactionType
	return c
}

func (c *client) WithDisplayText60(text string) Client {
	c.config.DisplayText60 = text
	return c
}

func (c *client) WithDisplayText200(text string) Client {
	c.config.DisplayText200 = text
	return c
}

func (c *client) WithURL(url string) Client {
	c.config.URL = url
	return c
}

func (c *client) WithTimeout(timeout time.Duration) Client {
	c.config.Timeout = timeout
	return c
}

func (c *client) WithTLSConfig(tlsConfig *tls.Config) Client {
	c.config.TLSConfig = tlsConfig
	return c
}

func (c *client) Validate() error {
	if c.config.RelyingPartyName == "" {
		return errors.ErrMissingRelyingPartyName
	}

	if c.config.RelyingPartyUUID == "" {
		return errors.ErrMissingRelyingPartyUUID
	}

	return nil
}
