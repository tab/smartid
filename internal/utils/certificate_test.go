package utils

import (
	"encoding/base64"
	"encoding/pem"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tab/smartid/internal/errors"
)

func Test_Certificate_Extract(t *testing.T) {
	tests := []struct {
		name     string
		cert     string
		expected *Person
		error    bool
		err      error
	}{
		{
			name: "Success",
			cert: readCertificate("testdata/valid_cert.pem"),
			expected: &Person{
				IdentityNumber: "PNOEE-30303039914",
				PersonalCode:   "30303039914",
				FirstName:      "TESTNUMBER",
				LastName:       "OK",
			},
			error: false,
		},
		{
			name:     "Error: failed to decode certificate",
			cert:     "invalid-base64-data",
			expected: nil,
			error:    true,
			err:      errors.ErrFailedToDecodeCertificate,
		},
		{
			name:     "Error: failed to parse certificate",
			cert:     readCertificate("testdata/failed_to_parse_cert.pem"),
			expected: nil,
			error:    true,
			err:      errors.ErrFailedToParseCertificate,
		},
		{
			name:     "Error: invalid serial number",
			cert:     readCertificate("testdata/invalid_identity_number.pem"),
			expected: nil,
			error:    true,
			err:      errors.ErrInvalidIdentityNumber,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Extract(tt.cert)

			if tt.error {
				assert.Error(t, err)
				assert.Equal(t, tt.err, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func readCertificate(filePath string) string {
	certPEM, err := os.ReadFile(filePath)
	if err != nil {
		return ""
	}

	block, _ := pem.Decode(certPEM)
	if block == nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(block.Bytes)
}
