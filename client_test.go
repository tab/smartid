package smartid

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/tab/smartid/internal/config"
	"github.com/tab/smartid/internal/errors"
)

func Test_NewClient(t *testing.T) {
	type result struct {
		config *config.Config
	}

	tests := []struct {
		name     string
		before   func(c Client)
		expected result
	}{
		{
			name: "Success",
			before: func(c Client) {
				c.
					WithRelyingPartyName("DEMO").
					WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000").
					WithCertificateLevel("QUALIFIED").
					WithHashType("SHA512").
					WithInteractionType("displayTextAndPIN").
					WithDisplayText60("Enter PIN1").
					WithDisplayText200("Confirm the authentication request and enter PIN1").
					WithURL("https://sid.demo.sk.ee/smart-id-rp/v2").
					WithTimeout(60 * time.Second)
			},
			expected: result{
				config: &config.Config{
					RelyingPartyName: "DEMO",
					RelyingPartyUUID: "00000000-0000-0000-0000-000000000000",
					CertificateLevel: "QUALIFIED",
					HashType:         "SHA512",
					InteractionType:  "displayTextAndPIN",
					DisplayText60:    "Enter PIN1",
					DisplayText200:   "Confirm the authentication request and enter PIN1",
					URL:              "https://sid.demo.sk.ee/smart-id-rp/v2",
					Timeout:          60 * time.Second,
					TLSConfig:        nil,
				},
			},
		},
		{
			name: "Default values",
			before: func(c Client) {
				c.
					WithRelyingPartyName("DEMO").
					WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000")
			},
			expected: result{
				config: &config.Config{
					RelyingPartyName: "DEMO",
					RelyingPartyUUID: "00000000-0000-0000-0000-000000000000",
					CertificateLevel: "QUALIFIED",
					HashType:         "SHA512",
					InteractionType:  "displayTextAndPIN",
					DisplayText60:    "Enter PIN1",
					DisplayText200:   "Confirm the authentication request and enter PIN1",
					URL:              "https://sid.demo.sk.ee/smart-id-rp/v2",
					Timeout:          60 * time.Second,
					TLSConfig:        nil,
				},
			},
		},
		{
			name: "Error: Missing relying party name",
			before: func(c Client) {
				c.WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000")
			},
			expected: result{
				config: &config.Config{
					RelyingPartyName: "",
					RelyingPartyUUID: "00000000-0000-0000-0000-000000000000",
					CertificateLevel: "QUALIFIED",
					HashType:         "SHA512",
					InteractionType:  "displayTextAndPIN",
					DisplayText60:    "Enter PIN1",
					DisplayText200:   "Confirm the authentication request and enter PIN1",
					URL:              "https://sid.demo.sk.ee/smart-id-rp/v2",
					Timeout:          60 * time.Second,
					TLSConfig:        nil,
				},
			},
		},
		{
			name: "Error: Missing relying party UUID",
			before: func(c Client) {
				c.WithRelyingPartyName("DEMO")
			},
			expected: result{
				config: &config.Config{
					RelyingPartyName: "DEMO",
					RelyingPartyUUID: "",
					CertificateLevel: "QUALIFIED",
					HashType:         "SHA512",
					InteractionType:  "displayTextAndPIN",
					DisplayText60:    "Enter PIN1",
					DisplayText200:   "Confirm the authentication request and enter PIN1",
					URL:              "https://sid.demo.sk.ee/smart-id-rp/v2",
					Timeout:          60 * time.Second,
					TLSConfig:        nil,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			tt.before(c)

			clientImpl := c.(*client)

			assert.NotNil(t, clientImpl)
			assert.Equal(t, tt.expected.config, clientImpl.config)
		})
	}
}

func Test_WithRelyingPartyName(t *testing.T) {
	c := NewClient()

	tests := []struct {
		name     string
		param    string
		expected string
	}{
		{
			name:     "Success",
			param:    "DEMO",
			expected: "DEMO",
		},
		{
			name:     "Empty",
			param:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c = c.WithRelyingPartyName(tt.param)
			clientImpl := c.(*client)
			assert.Equal(t, tt.expected, clientImpl.config.RelyingPartyName)
		})
	}
}

func Test_WithRelyingPartyUUID(t *testing.T) {
	c := NewClient()

	tests := []struct {
		name     string
		param    string
		expected string
	}{
		{
			name:     "Success",
			param:    "00000000-0000-0000-0000-000000000000",
			expected: "00000000-0000-0000-0000-000000000000",
		},
		{
			name:     "Empty",
			param:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c = c.WithRelyingPartyUUID(tt.param)
			clientImpl := c.(*client)
			assert.Equal(t, tt.expected, clientImpl.config.RelyingPartyUUID)
		})
	}
}

func Test_WithCertificateLevel(t *testing.T) {
	c := NewClient()

	tests := []struct {
		name     string
		param    string
		expected string
	}{
		{
			name:     "Success",
			param:    "QUALIFIED",
			expected: "QUALIFIED",
		},
		{
			name:     "Empty",
			param:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c = c.WithCertificateLevel(tt.param)
			clientImpl := c.(*client)
			assert.Equal(t, tt.expected, clientImpl.config.CertificateLevel)
		})
	}
}

func Test_WithHashType(t *testing.T) {
	c := NewClient()

	tests := []struct {
		name     string
		param    string
		expected string
	}{
		{
			name:     "Success",
			param:    "SHA512",
			expected: "SHA512",
		},
		{
			name:     "Empty",
			param:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c = c.WithHashType(tt.param)
			clientImpl := c.(*client)
			assert.Equal(t, tt.expected, clientImpl.config.HashType)
		})
	}
}

func Test_WithNonce(t *testing.T) {
	c := NewClient()

	tests := []struct {
		name     string
		param    string
		expected string
	}{
		{
			name:     "Success",
			param:    "1234567890",
			expected: "1234567890",
		},
		{
			name:     "Empty",
			param:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c = c.WithNonce(tt.param)
			clientImpl := c.(*client)
			assert.Equal(t, tt.expected, clientImpl.config.Nonce)
		})
	}
}

func Test_WithInteractionType(t *testing.T) {
	c := NewClient()

	tests := []struct {
		name     string
		param    string
		expected string
	}{
		{
			name:     "Success",
			param:    "displayTextAndPIN",
			expected: "displayTextAndPIN",
		},
		{
			name:     "Empty",
			param:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c = c.WithInteractionType(tt.param)
			clientImpl := c.(*client)
			assert.Equal(t, tt.expected, clientImpl.config.InteractionType)
		})
	}
}

func TestClient_WithDisplayText60(t *testing.T) {
	c := NewClient()

	tests := []struct {
		name     string
		param    string
		expected string
	}{
		{
			name:     "Success",
			param:    "Please enter PIN1",
			expected: "Please enter PIN1",
		},
		{
			name:     "Empty",
			param:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c = c.WithDisplayText60(tt.param)
			clientImpl := c.(*client)
			assert.Equal(t, tt.expected, clientImpl.config.DisplayText60)
		})
	}
}

func TestClient_WithDisplayText200(t *testing.T) {
	c := NewClient()

	tests := []struct {
		name     string
		param    string
		expected string
	}{
		{
			name:     "Success",
			param:    "Please confirm the authentication request and enter PIN1",
			expected: "Please confirm the authentication request and enter PIN1",
		},
		{
			name:     "Empty",
			param:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c = c.WithDisplayText200(tt.param)
			clientImpl := c.(*client)
			assert.Equal(t, tt.expected, clientImpl.config.DisplayText200)
		})
	}
}

func Test_WithURL(t *testing.T) {
	c := NewClient()

	tests := []struct {
		name     string
		param    string
		expected string
	}{
		{
			name:     "Success",
			param:    "https://sid.demo.sk.ee/smart-id-rp/v2",
			expected: "https://sid.demo.sk.ee/smart-id-rp/v2",
		},
		{
			name:     "Empty",
			param:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c = c.WithURL(tt.param)
			clientImpl := c.(*client)
			assert.Equal(t, tt.expected, clientImpl.config.URL)
		})
	}
}

func Test_WithTimeout(t *testing.T) {
	c := NewClient()

	tests := []struct {
		name     string
		param    time.Duration
		expected time.Duration
	}{
		{
			name:     "Success",
			param:    60 * time.Second,
			expected: 60 * time.Second,
		},
		{
			name:     "Zero",
			param:    0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c = c.WithTimeout(tt.param)
			clientImpl := c.(*client)
			assert.Equal(t, tt.expected, clientImpl.config.Timeout)
		})
	}
}

func TestClient_WithTLSConfig(t *testing.T) {
	manager, err := NewCertificateManager("./certs")
	assert.NoError(t, err)

	tlsConfig := manager.TLSConfig()

	type result struct {
		config *config.Config
	}

	tests := []struct {
		name     string
		before   func(c Client)
		expected result
	}{
		{
			name: "Success",
			before: func(c Client) {
				c.
					WithRelyingPartyName("DEMO").
					WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000").
					WithCertificateLevel("QUALIFIED").
					WithHashType("SHA512").
					WithInteractionType("displayTextAndPIN").
					WithDisplayText60("Enter PIN1").
					WithDisplayText200("Confirm the authentication request and enter PIN1").
					WithURL("https://sid.demo.sk.ee/smart-id-rp/v2").
					WithTimeout(60 * time.Second).
					WithTLSConfig(tlsConfig)
			},
			expected: result{
				config: &config.Config{
					RelyingPartyName: "DEMO",
					RelyingPartyUUID: "00000000-0000-0000-0000-000000000000",
					CertificateLevel: "QUALIFIED",
					HashType:         "SHA512",
					InteractionType:  "displayTextAndPIN",
					DisplayText60:    "Enter PIN1",
					DisplayText200:   "Confirm the authentication request and enter PIN1",
					URL:              "https://sid.demo.sk.ee/smart-id-rp/v2",
					Timeout:          60 * time.Second,
					TLSConfig:        tlsConfig,
				},
			},
		},
		{
			name: "Without TLS Config",
			before: func(c Client) {
				c.
					WithRelyingPartyName("DEMO").
					WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000")
			},
			expected: result{
				config: &config.Config{
					RelyingPartyName: "DEMO",
					RelyingPartyUUID: "00000000-0000-0000-0000-000000000000",
					CertificateLevel: "QUALIFIED",
					HashType:         "SHA512",
					InteractionType:  "displayTextAndPIN",
					DisplayText60:    "Enter PIN1",
					DisplayText200:   "Confirm the authentication request and enter PIN1",
					URL:              "https://sid.demo.sk.ee/smart-id-rp/v2",
					Timeout:          60 * time.Second,
					TLSConfig:        nil,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			tt.before(c)

			clientImpl := c.(*client)
			assert.Equal(t, tt.expected.config, clientImpl.config)
		})
	}
}

func Test_Validate(t *testing.T) {
	tests := []struct {
		name     string
		before   func(c Client)
		expected error
		error    bool
	}{
		{
			name: "Success",
			before: func(c Client) {
				c.
					WithRelyingPartyName("DEMO").
					WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000")
			},
			expected: nil,
			error:    false,
		},
		{
			name: "Error: Missing Relying Party Name",
			before: func(c Client) {
				c.
					WithRelyingPartyName("").
					WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000")
			},
			expected: errors.ErrMissingRelyingPartyName,
			error:    true,
		},
		{
			name: "Error: Missing Relying Party UUID",
			before: func(c Client) {
				c.
					WithRelyingPartyName("DEMO").
					WithRelyingPartyUUID("")
			},
			expected: errors.ErrMissingRelyingPartyUUID,
			error:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()

			tt.before(c)

			err := c.Validate()

			if tt.error {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
