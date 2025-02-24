package utils

import (
	"crypto/x509"
	"encoding/base64"
	"strings"

	"github.com/tab/smartid/internal/errors"
	"github.com/tab/smartid/internal/identity"
)

type Person struct {
	IdentityNumber string
	PersonalCode   string
	FirstName      string
	LastName       string
}

func Extract(encodedCert string) (*Person, error) {
	certBytes, err := base64.StdEncoding.DecodeString(encodedCert)
	if err != nil {
		return nil, errors.ErrFailedToDecodeCertificate
	}

	cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		return nil, errors.ErrFailedToParseCertificate
	}

	parts := strings.Split(cert.Subject.CommonName, ",")
	if len(parts) < 2 {
		return nil, errors.ErrInvalidCertificate
	}

	result, err := identity.Parse(cert.Subject.SerialNumber)
	if err != nil {
		return nil, err
	}
	firstName := strings.TrimSpace(parts[0])
	lastName := strings.TrimSpace(parts[1])

	return &Person{
		IdentityNumber: cert.Subject.SerialNumber,
		PersonalCode:   result.ID,
		FirstName:      firstName,
		LastName:       lastName,
	}, nil
}
