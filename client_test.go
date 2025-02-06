package smartid

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/tab/smartid/internal/config"
)

func Test_NewClient(t *testing.T) {
	client := NewClient()

	tests := []struct {
		name     string
		before   func()
		expected *Client
	}{
		{
			name: "Success",
			before: func() {
				client.
					WithRelyingPartyName("DEMO").
					WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000").
					WithCertificateLevel("QUALIFIED").
					WithHashType("SHA512").
					WithInteractionType("displayTextAndPIN").
					WithText("Enter PIN1").
					WithURL("https://sid.demo.sk.ee/smart-id-rp/v2").
					WithTimeout(60 * time.Second)
			},
			expected: &Client{
				config: &config.Config{
					RelyingPartyName: "DEMO",
					RelyingPartyUUID: "00000000-0000-0000-0000-000000000000",
					CertificateLevel: "QUALIFIED",
					HashType:         "SHA512",
					InteractionType:  "displayTextAndPIN",
					Text:             "Enter PIN1",
					URL:              "https://sid.demo.sk.ee/smart-id-rp/v2",
					Timeout:          60 * time.Second,
				},
			},
		},
		{
			name: "Default values",
			before: func() {
				client.
					WithRelyingPartyName("DEMO").
					WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000")
			},
			expected: &Client{
				config: &config.Config{
					RelyingPartyName: "DEMO",
					RelyingPartyUUID: "00000000-0000-0000-0000-000000000000",
					CertificateLevel: "QUALIFIED",
					HashType:         "SHA512",
					InteractionType:  "displayTextAndPIN",
					Text:             "Enter PIN1",
					URL:              "https://sid.demo.sk.ee/smart-id-rp/v2",
					Timeout:          60 * time.Second,
				},
			},
		},
		{
			name: "Error: Missing relying party name",
			before: func() {
				client.
					WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000")
			},
			expected: &Client{
				config: &config.Config{
					RelyingPartyName: "DEMO",
					RelyingPartyUUID: "00000000-0000-0000-0000-000000000000",
					CertificateLevel: "QUALIFIED",
					HashType:         "SHA512",
					InteractionType:  "displayTextAndPIN",
					Text:             "Enter PIN1",
					URL:              "https://sid.demo.sk.ee/smart-id-rp/v2",
					Timeout:          60 * time.Second,
				},
			},
		},
		{
			name: "Error: Missing relying party UUID",
			before: func() {
				client.
					WithRelyingPartyName("DEMO")
			},
			expected: &Client{
				config: &config.Config{
					RelyingPartyName: "DEMO",
					RelyingPartyUUID: "00000000-0000-0000-0000-000000000000",
					CertificateLevel: "QUALIFIED",
					HashType:         "SHA512",
					InteractionType:  "displayTextAndPIN",
					Text:             "Enter PIN1",
					URL:              "https://sid.demo.sk.ee/smart-id-rp/v2",
					Timeout:          60 * time.Second,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before()

			assert.NotNil(t, client)
			assert.Equal(t, tt.expected.config, client.config)
		})
	}
}

func Test_WithRelyingPartyName(t *testing.T) {
	client := NewClient()

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
			client = client.WithRelyingPartyName(tt.param)
			assert.Equal(t, tt.expected, client.config.RelyingPartyName)
		})
	}
}

func Test_WithRelyingPartyUUID(t *testing.T) {
	client := NewClient()

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
			client = client.WithRelyingPartyUUID(tt.param)
			assert.Equal(t, tt.expected, client.config.RelyingPartyUUID)
		})
	}
}

func Test_WithCertificateLevel(t *testing.T) {
	client := NewClient()

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
			client = client.WithCertificateLevel(tt.param)
			assert.Equal(t, tt.expected, client.config.CertificateLevel)
		})
	}
}

func Test_WithHashType(t *testing.T) {
	client := NewClient()

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
			client = client.WithHashType(tt.param)
			assert.Equal(t, tt.expected, client.config.HashType)
		})
	}
}

func Test_WithInteractionType(t *testing.T) {
	client := NewClient()

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
			client = client.WithInteractionType(tt.param)
			assert.Equal(t, tt.expected, client.config.InteractionType)
		})
	}
}

func Test_WithText(t *testing.T) {
	client := NewClient()

	tests := []struct {
		name     string
		param    string
		expected string
	}{
		{
			name:     "Success",
			param:    "Enter PIN1",
			expected: "Enter PIN1",
		},
		{
			name:     "Empty",
			param:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client = client.WithText(tt.param)
			assert.Equal(t, tt.expected, client.config.Text)
		})
	}
}

func Test_WithURL(t *testing.T) {
	client := NewClient()

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
			client = client.WithURL(tt.param)
			assert.Equal(t, tt.expected, client.config.URL)
		})
	}
}

func Test_WithTimeout(t *testing.T) {
	client := NewClient()

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
			client = client.WithTimeout(tt.param)
			assert.Equal(t, tt.expected, client.config.Timeout)
		})
	}
}
