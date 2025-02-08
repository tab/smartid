package smartid

import (
	"time"

	"github.com/tab/smartid/internal/config"
	"github.com/tab/smartid/internal/utils"
)

const (
	CertificateLevel = "QUALIFIED"
	InteractionType  = "displayTextAndPIN"
	Timeout          = 60 * time.Second
	URL              = "https://sid.demo.sk.ee/smart-id-rp/v2"
)

// Client holds the client configuration and the HTTP client
type Client struct {
	config *config.Config
}

// NewClient creates a new Smart-ID client instance
func NewClient() *Client {
	cfg := &config.Config{
		CertificateLevel: CertificateLevel,
		HashType:         utils.HashTypeSHA512,
		InteractionType:  InteractionType,
		URL:              URL,
		Timeout:          Timeout,
	}

	return &Client{
		config: cfg,
	}
}

// WithRelyingPartyName is option to set the RelyingPartyName
func (c *Client) WithRelyingPartyName(name string) *Client {
	c.config.RelyingPartyName = name
	return c
}

// WithRelyingPartyUUID is option to set the RelyingPartyUUID
func (c *Client) WithRelyingPartyUUID(id string) *Client {
	c.config.RelyingPartyUUID = id
	return c
}

// WithCertificateLevel is option to set the certificate level
func (c *Client) WithCertificateLevel(level string) *Client {
	c.config.CertificateLevel = level
	return c
}

// WithHashType is option to set the hash type
func (c *Client) WithHashType(hashType string) *Client {
	c.config.HashType = hashType
	return c
}

// WithInteractionType is option to set the interaction type
func (c *Client) WithInteractionType(interactionType string) *Client {
	c.config.InteractionType = interactionType
	return c
}

// WithText is option to set the display text
func (c *Client) WithText(text string) *Client {
	c.config.Text = text
	return c
}

// WithURL is option to set the Smart-ID service URL
func (c *Client) WithURL(url string) *Client {
	c.config.URL = url
	return c
}

// WithTimeout is option to set the request timeout
func (c *Client) WithTimeout(timeout time.Duration) *Client {
	c.config.Timeout = timeout
	return c
}
