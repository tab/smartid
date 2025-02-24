package smartid

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/tab/smartid/internal/config"
	"github.com/tab/smartid/internal/errors"
	"github.com/tab/smartid/internal/utils"
)

const (
	CertificateLevel = "QUALIFIED"
	InteractionType  = "displayTextAndPIN"
	Text             = "Enter PIN1"
	Timeout          = 60 * time.Second
	URL              = "https://sid.demo.sk.ee/smart-id-rp/v2"
)

type Client interface {
	CreateSession(ctx context.Context, nationalIdentityNumber string) (*Session, error)
	FetchSession(ctx context.Context, sessionId string) (*Person, error)

	WithRelyingPartyName(name string) Client
	WithRelyingPartyUUID(id string) Client
	WithCertificateLevel(level string) Client
	WithHashType(hashType string) Client
	WithInteractionType(interactionType string) Client
	WithText(text string) Client
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
		Text:             Text,
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

func (c *client) WithInteractionType(interactionType string) Client {
	c.config.InteractionType = interactionType
	return c
}

func (c *client) WithText(text string) Client {
	c.config.Text = text
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
