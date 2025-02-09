package smartid

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tab/smartid/internal/errors"
	"github.com/tab/smartid/internal/models"
)

func Test_Authenticate(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name     string
		before   func(w http.ResponseWriter, r *http.Request)
		param    string
		expected *models.Person
		err      error
		error    bool
	}{
		{
			name: "Success",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`
{
	"state": "COMPLETE",
	"result": {
		"endResult": "OK",
		"documentNumber": "PNOEE-30303039914"
	},
	"signature": {
		"value": "NUqfcuvZrQdJy9yAIpmDaOi22uWnCL7rbF+5Vx/th2pa7VK3+xRBM7S9CqEn9pRWsoOddqFDUfbz7w6Vd/AvaduLH3UEWs+LCTlQ+9liCGUcY4N97xMhlVwv1MnybBbDKKk7e+xXAdFGV7T+2lE5PwP9h4YyCl/1Jg1lXcuNWJEcu2E1bcJtOI6yDO+3PYEDuc/NNsj/1SZFvg+ffLhJOMKEOJe+Jxf6hsn6NoAFyBYBDvKGeAX92FHej5BFQbvk/sWJee9ENC+Mmjsr+rUiJI0iKh+WN0fiQYzdtv0TsowcGF0vqrRnlbDEc301xetowJBcefko8DcroqtvgzXQ3W0ruEeYKbehzEmB/iEI1iBjQi3hxrfaXD1cgZRzWcIurSzgv+rB5QE1xWV7GPRu9gV5b/yKRkfIbPclOa6OTpjlKTu+EG6qM7z1H9+UMp/Lx62Amin57W+oH0kiDm5zAMTETEDkRE0WXpKOlETvOUDS26hsXa9KlTnapisSpfSc0s2dCXjqYQ1Faw18gKKnZBhG5WkrqaaHpMGFqHVfPSVIe8uALVVaCBHzQ/Nly8dhE26YJXSY+BoIjrX4znXVCE38hpWPeYGMu+4Y/gm+STVkwQKVlXYIjG5nWpnkNI5ivvfHhusiLJf9bPKxMPSRtjme3g69vU4NHMpGnAJp2ER+i/S1DigULSEUQscrFNYFzu8Ha67a0SnRc1ozu2VCatyx+eFoVSuoriiVnOnb1/mXXsS3keBa2PygazHnMI2Hd7aKnhpM41fbKoa9awZbCy8Udw6gTfSyqLZNZ07fB1x9wEVnV7ZMZn8NcsBWYTR9v9DpZkYXO/7GgWTBCKtTwL0/TOWKmEkibPhmtGNqzA/+LeJoGCXgIvqBRZjVHLNZu90CtLCSddaSJR6MgMyfL4eGIWTybZULll8UO7G2XkRl46HsPVjv1CILmy80V6VHhRCBUEOpn9TWn6Q+hGelshbQozy4/hfidF9JkdXi0y/3fHBLyhsAcLsZ2+n6qoxs",
		"algorithm": "sha256WithRSAEncryption"
	},
	"cert": {
		"value": "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",
		"certificateLevel": "QUALIFIED"
	},
	"interactionFlowUsed": "displayTextAndPIN"
}`))
			},
			param: "PNOEE-30303039914",
			expected: &models.Person{
				IdentityNumber: "PNOEE-30303039914",
				PersonalCode:   "30303039914",
				FirstName:      "TESTNUMBER",
				LastName:       "OK",
			},
			err:   nil,
			error: false,
		},
		{
			name: "Error: Invalid certificate",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`
{
	"state": "COMPLETE",
	"result": {
		"endResult": "OK",
		"documentNumber": "PNOEE-30303039914"
	},
	"signature": {
		"value": "invalid-signature",
		"algorithm": "sha256WithRSAEncryption"
	},
	"cert": {
		"value": "invalid-certificate",
		"certificateLevel": "QUALIFIED"
	},
	"interactionFlowUsed": "displayTextAndPIN"
}`))
			},
			param:    "PNOEE-30303039914",
			expected: &models.Person{},
			err:      errors.ErrFailedToDecodeCertificate,
			error:    true,
		},
		{
			name: "Error: Authentication is running",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"state": "RUNNING"}`))
			},
			param:    "PNOEE-30303039914",
			expected: &models.Person{},
			err:      errors.ErrAuthenticationIsRunning,
			error:    true,
		},
		{
			name: "Error: USER_REFUSED",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"state": "COMPLETE", "result": {"endResult": "USER_REFUSED"}}`))
			},
			param:    "PNOEE-30303039914",
			expected: &models.Person{},
			err:      &Error{Code: "USER_REFUSED"},
			error:    true,
		},
		{
			name: "Error: TIMEOUT",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"state": "COMPLETE", "result": {"endResult": "TIMEOUT"}}`))
			},
			param:    "PNOEE-30303039914",
			expected: &models.Person{},
			err:      &Error{Code: "TIMEOUT"},
			error:    true,
		},
		{
			name: "Error: result UNKNOWN",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"state": "COMPLETE", "result": {"endResult": "UNKNOWN"}}`))
			},
			param:    "PNOEE-30303039914",
			expected: &models.Person{},
			err:      errors.ErrUnsupportedResult,
			error:    true,
		},
		{
			name: "Error: state UNKNOWN",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"state": "UNKNOWN"}`))
			},
			param:    "PNOEE-30303039914",
			expected: &models.Person{},
			err:      errors.ErrUnsupportedState,
			error:    true,
		},
		{
			name: "Bad Request",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"title": "Bad Request", "status": 400}`))
			},
			param:    "PNOEE-30303039914",
			expected: &models.Person{},
			err:      errors.ErrSmartIdProviderError,
			error:    true,
		},
		{
			name: "Internal Server Error",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"title": "Internal Server Error", "status": 500}`))
			},
			param:    "PNOEE-30303039914",
			expected: &models.Person{},
			err:      errors.ErrSmartIdProviderError,
			error:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testServer := httptest.NewServer(http.HandlerFunc(tt.before))
			defer testServer.Close()

			client := NewClient()
			client.WithRelyingPartyName("DEMO").
				WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000").
				WithURL(testServer.URL)

			session, resultCh := client.Authenticate(ctx, tt.param)

			result := <-resultCh

			if tt.error {
				assert.Error(t, result.Err)
				assert.Nil(t, result.Person)
				assert.Equal(t, tt.err, result.Err)
			} else {
				assert.NotNil(t, session)
				assert.NoError(t, result.Err)
				assert.Equal(t, tt.expected, result.Person)
			}
		})
	}
}

func Test_Validate(t *testing.T) {
	client := NewClient()

	tests := []struct {
		name     string
		before   func()
		expected error
		error    bool
	}{
		{
			name: "Success",
			before: func() {
				client.
					WithRelyingPartyName("DEMO").
					WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000")
			},
			expected: nil,
			error:    false,
		},
		{
			name: "Error: Missing Relying Party Name",
			before: func() {
				client.
					WithRelyingPartyName("").
					WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000")
			},
			expected: errors.ErrMissingRelyingPartyName,
			error:    true,
		},
		{
			name: "Error: Missing Relying Party UUID",
			before: func() {
				client.
					WithRelyingPartyName("DEMO").
					WithRelyingPartyUUID("")
			},
			expected: errors.ErrMissingRelyingPartyUUID,
			error:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before()

			err := client.Validate()

			if tt.error {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_Error(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		expected string
	}{
		{
			name:     "Error: USER_REFUSED",
			code:     "USER_REFUSED",
			expected: "authentication failed: USER_REFUSED",
		},
		{
			name:     "Error: USER_REFUSED_DISPLAYTEXTANDPIN",
			code:     "USER_REFUSED_DISPLAYTEXTANDPIN",
			expected: "authentication failed: USER_REFUSED_DISPLAYTEXTANDPIN",
		},
		{
			name:     "Error: USER_REFUSED_VC_CHOICE",
			code:     "USER_REFUSED_VC_CHOICE",
			expected: "authentication failed: USER_REFUSED_VC_CHOICE",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &Error{Code: tt.code}
			assert.Equal(t, tt.expected, err.Error())
		})
	}
}
