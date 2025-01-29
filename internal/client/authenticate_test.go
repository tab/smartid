package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"smartid/internal/config"
	"smartid/internal/models"
)

func Test_Authenticate(t *testing.T) {
	tests := []struct {
		name     string
		before   func(w http.ResponseWriter, r *http.Request)
		param    string
		expected *models.AuthenticateResponse
		error    bool
	}{
		{
			name: "Success",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"sessionID": "8fdb516d-1a82-43ba-b82d-be63df569b86", "code": "1234"}`))
			},
			param: "PNOEE-30303039914",
			expected: &models.AuthenticateResponse{
				SessionID: "8fdb516d-1a82-43ba-b82d-be63df569b86",
				Code:      "1234",
			},
			error: false,
		},
		{
			name:     "Error",
			before:   func(w http.ResponseWriter, r *http.Request) {},
			param:    "not-a-personal-code",
			expected: &models.AuthenticateResponse{},
			error:    true,
		},
		{
			name: "Not found",
			before: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(`{"title": "Not Found", "status": 404}`))
			},
			param:    "PNOEE-30303039914",
			expected: &models.AuthenticateResponse{},
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
			expected: &models.AuthenticateResponse{},
			error:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testServer := httptest.NewServer(http.HandlerFunc(tt.before))
			defer testServer.Close()

			client, err := NewClient(
				config.WithRelyingPartyName("DEMO"),
				config.WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000"),
				config.WithURL(testServer.URL),
			)
			assert.NoError(t, err)

			response, err := client.Authenticate(tt.param)

			if tt.error {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.SessionID, response.SessionID)
			}
		})
	}
}
