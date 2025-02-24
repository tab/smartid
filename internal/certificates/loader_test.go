package certificates

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tab/smartid/internal/errors"
)

func Test_Certificates_LoadFromFile(t *testing.T) {
	tests := []struct {
		name string
		path string
		err  error
	}{
		{
			name: "Success",
			path: "testdata/valid/cert.pem",
			err:  nil,
		},
		{
			name: "Failed to read certificate file",
			path: "testdata/missing.pem",
			err:  errors.ErrFailedToReadCertificateFile,
		},
		{
			name: "Failed to decode certificate file",
			path: "testdata/invalid/invalid_decode.pem",
			err:  errors.ErrFailedToDecodeCertificateFile,
		},
		{
			name: "Failed to parse certificate file",
			path: "testdata/invalid/invalid_parse.pem",
			err:  errors.ErrFailedToParseCertificateFile,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoadFromFile(tt.path)

			if tt.err != nil {
				assert.Equal(t, tt.err, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func Test_Certificates_LoadFromDir(t *testing.T) {
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
			name: "Failed to read directory",
			dir:  "testdata/missing",
			err:  errors.ErrFailedToReadCertificateFile,
		},
		{
			name: "Failed to decode certificate file",
			dir:  "testdata/invalid",
			err:  errors.ErrFailedToDecodeCertificateFile,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoadFromDir(tt.dir)

			if tt.err != nil {
				assert.Equal(t, tt.err, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
