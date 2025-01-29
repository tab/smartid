package config

import "testing"

import "github.com/stretchr/testify/assert"

func Test_WithRelyingPartyName(t *testing.T) {
	c := &Config{}

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
			WithRelyingPartyName(tt.param)(c)
			assert.Equal(t, tt.expected, c.RelyingPartyName)
		})
	}
}

func Test_WithRelyingPartyUUID(t *testing.T) {
	c := &Config{}

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
			WithRelyingPartyUUID(tt.param)(c)
			assert.Equal(t, tt.expected, c.RelyingPartyUUID)
		})
	}
}

func Test_WithCertificateLevel(t *testing.T) {
	c := &Config{}

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
			WithCertificateLevel(tt.param)(c)
			assert.Equal(t, tt.expected, c.CertificateLevel)
		})
	}
}

func Test_WithHashType(t *testing.T) {
	c := &Config{}

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
			WithHashType(tt.param)(c)
			assert.Equal(t, tt.expected, c.HashType)
		})
	}
}

func Test_WithInteractionType(t *testing.T) {
	c := &Config{}

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
			WithInteractionType(tt.param)(c)
			assert.Equal(t, tt.expected, c.InteractionType)
		})
	}
}

func Test_WithText(t *testing.T) {
	c := &Config{}

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
			WithText(tt.param)(c)
			assert.Equal(t, tt.expected, c.Text)
		})
	}
}

func Test_WithURL(t *testing.T) {
	c := &Config{}

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
			WithURL(tt.param)(c)
			assert.Equal(t, tt.expected, c.URL)
		})
	}
}

func Test_WithTimeout(t *testing.T) {
	c := &Config{}

	tests := []struct {
		name     string
		param    int
		expected int
	}{
		{
			name:     "Success",
			param:    60,
			expected: 60,
		},
		{
			name:     "Zero",
			param:    0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WithTimeout(tt.param)(c)
			assert.Equal(t, tt.expected, c.Timeout)
		})
	}
}
