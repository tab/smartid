package certificates

import (
	"crypto/x509"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tab/smartid/internal/errors"
)

func Test_NewCertificatePinner(t *testing.T) {
	tests := []struct {
		name string
		dir  string
		err  error
	}{
		{
			name: "Success",
			dir:  "testdata/valid",
			err:  nil,
		},
		{
			name: "Error: Failed to read certificate file",
			dir:  "testdata/missing",
			err:  errors.ErrFailedToReadCertificateFile,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewCertificatePinner(tt.dir)

			if tt.err != nil {
				assert.Equal(t, tt.err, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func Test_Pinner_VerifyPeerCertificate(t *testing.T) {
	certPEM, err := LoadFromFile("testdata/valid/cert.pem")
	assert.NoError(t, err)

	tests := []struct {
		name     string
		certs    []*x509.Certificate
		rawCerts [][]byte
		err      error
	}{
		{
			name:     "Success",
			certs:    []*x509.Certificate{certPEM},
			rawCerts: [][]byte{certPEM.Raw},
			err:      nil,
		},
		{
			name:     "Error: No matching certificate",
			certs:    []*x509.Certificate{certPEM},
			rawCerts: [][]byte{},
			err:      errors.ErrFailedToVerifyCertificate,
		},
		{
			name:     "Error: Invalid certificate",
			certs:    []*x509.Certificate{certPEM},
			rawCerts: [][]byte{[]byte("invalid")},
			err:      errors.ErrFailedToVerifyCertificate,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pinner{
				certificates: tt.certs,
			}

			err := p.VerifyPeerCertificate(tt.rawCerts, nil)
			if tt.err != nil {
				assert.Equal(t, tt.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_Pinner_TLSConfig(t *testing.T) {
	p := &Pinner{}
	config := p.TLSConfig()

	assert.NotNil(t, config)
	assert.NotNil(t, config.VerifyPeerCertificate)
	assert.Equal(t, uint16(0x0303), config.MinVersion) // TLS 1.2
}
