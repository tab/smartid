package certificates

import (
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"

	"github.com/tab/smartid/internal/errors"
)

type Pinner struct {
	certificates []*x509.Certificate
}

// NewCertificatePinner creates a new certificate pinner instance
func NewCertificatePinner(certsDir string) (*Pinner, error) {
	certs, err := LoadFromDir(certsDir)
	if err != nil {
		return nil, err
	}

	return &Pinner{
		certificates: certs,
	}, nil
}

// TLSConfig returns a new tls.Config instance with the certificate pinner
func (p *Pinner) TLSConfig() *tls.Config {
	return &tls.Config{
		VerifyPeerCertificate: p.VerifyPeerCertificate,
		MinVersion:            tls.VersionTLS12,
	}
}

// VerifyPeerCertificate verifies the peer certificate against the pinned certificates
func (p *Pinner) VerifyPeerCertificate(rawCerts [][]byte, _ [][]*x509.Certificate) error {
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
