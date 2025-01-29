package client

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tab/smart-id/internal/config"
)

func Test_NewClient(t *testing.T) {
	tests := []struct {
		name     string
		options  []config.Option
		expected *Client
		error    bool
	}{
		{
			name: "Success",
			options: []config.Option{
				config.WithRelyingPartyName("DEMO"),
				config.WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000"),
				config.WithCertificateLevel("QUALIFIED"),
				config.WithHashType("SHA512"),
				config.WithInteractionType("displayTextAndPIN"),
				config.WithText("Enter PIN1"),
				config.WithURL("https://sid.demo.sk.ee/smart-id-rp/v2"),
				config.WithTimeout(60),
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
					Timeout:          60,
				},
			},
			error: false,
		},
		{
			name: "Default values",
			options: []config.Option{
				config.WithRelyingPartyName("DEMO"),
				config.WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000"),
			},
			expected: &Client{
				config: &config.Config{
					RelyingPartyName: "DEMO",
					RelyingPartyUUID: "00000000-0000-0000-0000-000000000000",
					CertificateLevel: "QUALIFIED",
					HashType:         "SHA512",
					InteractionType:  "displayTextAndPIN",
					URL:              "https://sid.demo.sk.ee/smart-id-rp/v2",
					Timeout:          60,
				},
			},
			error: false,
		},
		{
			name: "Error: Missing relying party name",
			options: []config.Option{
				config.WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000"),
				config.WithCertificateLevel("QUALIFIED"),
				config.WithHashType("SHA512"),
				config.WithInteractionType("displayTextAndPIN"),
				config.WithText("Enter PIN1"),
				config.WithURL("https://sid.demo.sk.ee/smart-id-rp/v2"),
				config.WithTimeout(60),
			},
			expected: nil,
			error:    true,
		},
		{
			name: "Error: Missing relying party UUID",
			options: []config.Option{
				config.WithRelyingPartyName("DEMO"),
				config.WithCertificateLevel("QUALIFIED"),
				config.WithHashType("SHA512"),
				config.WithInteractionType("displayTextAndPIN"),
				config.WithText("Enter PIN1"),
				config.WithURL("https://sid.demo.sk.ee/smart-id-rp/v2"),
				config.WithTimeout(60),
			},
			expected: nil,
			error:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.options...)

			if tt.error {
				assert.Error(t, err)
				assert.Nil(t, client)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, client)

				assert.Equal(t, tt.expected.config, client.config)
			}
		})
	}
}
