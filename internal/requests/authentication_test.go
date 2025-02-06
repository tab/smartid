package requests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"

	"github.com/tab/smartid/internal/config"
	"github.com/tab/smartid/internal/models"
)

func Test_Call(t *testing.T) {
	ctx := context.Background()
	cfg := &config.Config{
		RelyingPartyName: "DEMO",
		RelyingPartyUUID: "00000000-0000-0000-0000-000000000000",
		CertificateLevel: "QUALIFIED",
		InteractionType:  "displayTextAndPIN",
		Text:             "Enter PIN1",
		HashType:         "SHA512",
		Timeout:          10 * time.Second,
	}
	httpClient := resty.New().
		SetTimeout(cfg.Timeout).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json")

	tests := []struct {
		name     string
		before   func(w http.ResponseWriter, r *http.Request)
		param    string
		expected *models.AuthenticationResponse
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
			expected: &models.AuthenticationResponse{
				State: "COMPLETE",
				Result: models.Result{
					EndResult:      "OK",
					DocumentNumber: "PNOEE-30303039914",
				},
				Signature: models.Signature{
					Value:     "NUqfcuvZrQdJy9yAIpmDaOi22uWnCL7rbF+5Vx/th2pa7VK3+xRBM7S9CqEn9pRWsoOddqFDUfbz7w6Vd/AvaduLH3UEWs+LCTlQ+9liCGUcY4N97xMhlVwv1MnybBbDKKk7e+xXAdFGV7T+2lE5PwP9h4YyCl/1Jg1lXcuNWJEcu2E1bcJtOI6yDO+3PYEDuc/NNsj/1SZFvg+ffLhJOMKEOJe+Jxf6hsn6NoAFyBYBDvKGeAX92FHej5BFQbvk/sWJee9ENC+Mmjsr+rUiJI0iKh+WN0fiQYzdtv0TsowcGF0vqrRnlbDEc301xetowJBcefko8DcroqtvgzXQ3W0ruEeYKbehzEmB/iEI1iBjQi3hxrfaXD1cgZRzWcIurSzgv+rB5QE1xWV7GPRu9gV5b/yKRkfIbPclOa6OTpjlKTu+EG6qM7z1H9+UMp/Lx62Amin57W+oH0kiDm5zAMTETEDkRE0WXpKOlETvOUDS26hsXa9KlTnapisSpfSc0s2dCXjqYQ1Faw18gKKnZBhG5WkrqaaHpMGFqHVfPSVIe8uALVVaCBHzQ/Nly8dhE26YJXSY+BoIjrX4znXVCE38hpWPeYGMu+4Y/gm+STVkwQKVlXYIjG5nWpnkNI5ivvfHhusiLJf9bPKxMPSRtjme3g69vU4NHMpGnAJp2ER+i/S1DigULSEUQscrFNYFzu8Ha67a0SnRc1ozu2VCatyx+eFoVSuoriiVnOnb1/mXXsS3keBa2PygazHnMI2Hd7aKnhpM41fbKoa9awZbCy8Udw6gTfSyqLZNZ07fB1x9wEVnV7ZMZn8NcsBWYTR9v9DpZkYXO/7GgWTBCKtTwL0/TOWKmEkibPhmtGNqzA/+LeJoGCXgIvqBRZjVHLNZu90CtLCSddaSJR6MgMyfL4eGIWTybZULll8UO7G2XkRl46HsPVjv1CILmy80V6VHhRCBUEOpn9TWn6Q+hGelshbQozy4/hfidF9JkdXi0y/3fHBLyhsAcLsZ2+n6qoxs",
					Algorithm: "sha256WithRSAEncryption",
				},
				Cert: models.Certificate{
					Value:            "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",
					CertificateLevel: "QUALIFIED",
				},
				InteractionFlowUsed: "displayTextAndPIN",
			},
			error: false,
		},
		{
			name: "Error: USER_REFUSED",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"state": "COMPLETE", "result": {"endResult": "USER_REFUSED"}}`))
			},
			param: "PNOEE-30303039914",
			expected: &models.AuthenticationResponse{
				State: "COMPLETE",
				Result: models.Result{
					EndResult: "USER_REFUSED",
				},
			},
			error: false,
		},
		{
			name: "Error: TIMEOUT",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"state": "COMPLETE", "result": {"endResult": "TIMEOUT"}}`))
			},
			param: "PNOEE-30303039914",
			expected: &models.AuthenticationResponse{
				State: "COMPLETE",
				Result: models.Result{
					EndResult: "TIMEOUT",
				},
			},
			error: false,
		},
		{
			name: "Bad Request",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"title": "Bad Request", "status": 400}`))
			},
			param:    "PNOEE-30303039914",
			expected: &models.AuthenticationResponse{},
			error:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testServer := httptest.NewServer(http.HandlerFunc(tt.before))
			defer testServer.Close()

			cfg.URL = testServer.URL

			response, err := Call(ctx, httpClient, cfg, tt.param)

			if tt.error {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.State, response.State)
				assert.Equal(t, tt.expected.Result, response.Result)
			}
		})
	}
}

func Test_createAuthenticationSession(t *testing.T) {
	ctx := context.Background()
	cfg := &config.Config{
		RelyingPartyName: "DEMO",
		RelyingPartyUUID: "00000000-0000-0000-0000-000000000000",
		CertificateLevel: "QUALIFIED",
		InteractionType:  "displayTextAndPIN",
		Text:             "Enter PIN1",
		HashType:         "SHA512",
		Timeout:          10 * time.Second,
	}
	httpClient := resty.New().
		SetTimeout(cfg.Timeout).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json")

	tests := []struct {
		name     string
		before   func(w http.ResponseWriter, r *http.Request)
		identity string
		expected *models.AuthenticationSessionResponse
		error    bool
	}{
		{
			name: "Success",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"sessionID": "8fdb516d-1a82-43ba-b82d-be63df569b86", "code": "1234"}`))
			},
			identity: "PNOEE-30303039914",
			expected: &models.AuthenticationSessionResponse{
				SessionID: "8fdb516d-1a82-43ba-b82d-be63df569b86",
				Code:      "1234",
			},
			error: false,
		},
		{
			name:     "Error",
			before:   func(w http.ResponseWriter, r *http.Request) {},
			identity: "not-a-personal-code",
			expected: &models.AuthenticationSessionResponse{},
			error:    true,
		},
		{
			name: "Not found",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(`{"title": "Not Found", "status": 404}`))
			},
			identity: "PNOEE-30303039914",
			expected: &models.AuthenticationSessionResponse{},
			error:    true,
		},
		{
			name: "Bad Request",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"title": "Bad Request", "status": 400}`))
			},
			identity: "PNOEE-30303039914",
			expected: &models.AuthenticationSessionResponse{},
			error:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testServer := httptest.NewServer(http.HandlerFunc(tt.before))
			defer testServer.Close()

			cfg.URL = testServer.URL

			response, err := createAuthenticationSession(ctx, httpClient, cfg, tt.identity)

			if tt.error {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.SessionID, response.SessionID)
			}
		})
	}
}

func Test_fetchAuthenticationSession(t *testing.T) {
	ctx := context.Background()

	cfg := &config.Config{
		RelyingPartyName: "DEMO",
		RelyingPartyUUID: "00000000-0000-0000-0000-000000000000",
		CertificateLevel: "QUALIFIED",
		InteractionType:  "displayTextAndPIN",
		Text:             "Enter PIN1",
		HashType:         "SHA512",
		Timeout:          10 * time.Second,
	}
	httpClient := resty.New().
		SetTimeout(cfg.Timeout).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json")

	id := "8fdb516d-1a82-43ba-b82d-be63df569b86"

	tests := []struct {
		name     string
		before   func(w http.ResponseWriter, r *http.Request)
		id       string
		expected *models.AuthenticationResponse
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
			id: id,
			expected: &models.AuthenticationResponse{
				State: "COMPLETE",
				Result: models.Result{
					EndResult:      "OK",
					DocumentNumber: "PNOEE-30303039914",
				},
				Signature: models.Signature{
					Value:     "NUqfcuvZrQdJy9yAIpmDaOi22uWnCL7rbF+5Vx/th2pa7VK3+xRBM7S9CqEn9pRWsoOddqFDUfbz7w6Vd/AvaduLH3UEWs+LCTlQ+9liCGUcY4N97xMhlVwv1MnybBbDKKk7e+xXAdFGV7T+2lE5PwP9h4YyCl/1Jg1lXcuNWJEcu2E1bcJtOI6yDO+3PYEDuc/NNsj/1SZFvg+ffLhJOMKEOJe+Jxf6hsn6NoAFyBYBDvKGeAX92FHej5BFQbvk/sWJee9ENC+Mmjsr+rUiJI0iKh+WN0fiQYzdtv0TsowcGF0vqrRnlbDEc301xetowJBcefko8DcroqtvgzXQ3W0ruEeYKbehzEmB/iEI1iBjQi3hxrfaXD1cgZRzWcIurSzgv+rB5QE1xWV7GPRu9gV5b/yKRkfIbPclOa6OTpjlKTu+EG6qM7z1H9+UMp/Lx62Amin57W+oH0kiDm5zAMTETEDkRE0WXpKOlETvOUDS26hsXa9KlTnapisSpfSc0s2dCXjqYQ1Faw18gKKnZBhG5WkrqaaHpMGFqHVfPSVIe8uALVVaCBHzQ/Nly8dhE26YJXSY+BoIjrX4znXVCE38hpWPeYGMu+4Y/gm+STVkwQKVlXYIjG5nWpnkNI5ivvfHhusiLJf9bPKxMPSRtjme3g69vU4NHMpGnAJp2ER+i/S1DigULSEUQscrFNYFzu8Ha67a0SnRc1ozu2VCatyx+eFoVSuoriiVnOnb1/mXXsS3keBa2PygazHnMI2Hd7aKnhpM41fbKoa9awZbCy8Udw6gTfSyqLZNZ07fB1x9wEVnV7ZMZn8NcsBWYTR9v9DpZkYXO/7GgWTBCKtTwL0/TOWKmEkibPhmtGNqzA/+LeJoGCXgIvqBRZjVHLNZu90CtLCSddaSJR6MgMyfL4eGIWTybZULll8UO7G2XkRl46HsPVjv1CILmy80V6VHhRCBUEOpn9TWn6Q+hGelshbQozy4/hfidF9JkdXi0y/3fHBLyhsAcLsZ2+n6qoxs",
					Algorithm: "sha256WithRSAEncryption",
				},
				Cert: models.Certificate{
					Value:            "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",
					CertificateLevel: "QUALIFIED",
				},
				InteractionFlowUsed: "displayTextAndPIN",
			},
			error: false,
		},
		{
			name: "Error: USER_REFUSED",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"state": "COMPLETE", "result": {"endResult": "USER_REFUSED"}}`))
			},
			id: id,
			expected: &models.AuthenticationResponse{
				State: "COMPLETE",
				Result: models.Result{
					EndResult: "USER_REFUSED",
				},
			},
			error: false,
		},
		{
			name: "Error: TIMEOUT",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"state": "COMPLETE", "result": {"endResult": "TIMEOUT"}}`))
			},
			id: id,
			expected: &models.AuthenticationResponse{
				State: "COMPLETE",
				Result: models.Result{
					EndResult: "TIMEOUT",
				},
			},
			error: false,
		},
		{
			name: "Not found",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(`{"title": "Not Found", "status": 404}`))
			},
			id:       id,
			expected: &models.AuthenticationResponse{},
			error:    true,
		},
		{
			name: "Bad Request",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"title": "Bad Request", "status": 400}`))
			},
			id:       id,
			expected: &models.AuthenticationResponse{},
			error:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testServer := httptest.NewServer(http.HandlerFunc(tt.before))
			defer testServer.Close()

			cfg.URL = testServer.URL

			response, err := fetchAuthenticationSession(ctx, httpClient, cfg, id)

			if tt.error {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.State, response.State)
				assert.Equal(t, tt.expected.Result, response.Result)
			}
		})
	}
}
