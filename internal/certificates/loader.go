package certificates

import (
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"

	"github.com/tab/smartid/internal/errors"
)

const CertExtension = ".pem"

func LoadFromFile(path string) (*x509.Certificate, error) {
	certPEM, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.ErrFailedToReadCertificateFile
	}

	block, _ := pem.Decode(certPEM)
	if block == nil {
		return nil, errors.ErrFailedToDecodeCertificateFile
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, errors.ErrFailedToParseCertificateFile
	}

	return cert, nil
}

func LoadFromDir(dir string) ([]*x509.Certificate, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, errors.ErrFailedToReadCertificateFile
	}

	pemCount := 0
	for _, file := range files {
		if filepath.Ext(file.Name()) == CertExtension {
			pemCount++
		}
	}
	certs := make([]*x509.Certificate, 0, pemCount)

	for _, file := range files {
		if filepath.Ext(file.Name()) != CertExtension {
			continue
		}

		cert, err := LoadFromFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}
		certs = append(certs, cert)
	}

	return certs, nil
}
