package smartid

import (
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"

	"github.com/tab/smartid/internal/certificates"
	"github.com/tab/smartid/internal/errors"
)

type Manager struct {
	certificates []*x509.Certificate
}

// NewCertificateManager creates a new certificate manager instance
func NewCertificateManager(certsDir string) (*Manager, error) {
	certs, err := certificates.LoadFromDir(certsDir)
	if err != nil {
		return nil, err
	}

	return &Manager{
		certificates: certs,
	}, nil
}

// TLSConfig returns a new tls.Config instance with the certificate pinning
func (p *Manager) TLSConfig() *tls.Config {
	return &tls.Config{
		VerifyPeerCertificate: p.VerifyPeerCertificate,
		MinVersion:            tls.VersionTLS12,
	}
}

// VerifyPeerCertificate verifies the peer certificate against the pinned certificates
func (p *Manager) VerifyPeerCertificate(rawCerts [][]byte, _ [][]*x509.Certificate) error {
	for _, rawCert := range rawCerts {
		cert, err := x509.ParseCertificate(rawCert)
		if err != nil {
			continue
		}

		publicKeyHash := sha256.Sum256(cert.RawSubjectPublicKeyInfo)
		actualPin := base64.StdEncoding.EncodeToString(publicKeyHash[:])

		for _, expectedCert := range p.certificates {
			expectedHash := sha256.Sum256(expectedCert.RawSubjectPublicKeyInfo)
			expectedPin := base64.StdEncoding.EncodeToString(expectedHash[:])

			if actualPin == expectedPin {
				return nil
			}
		}
	}

	return errors.ErrFailedToVerifyCertificate
}
