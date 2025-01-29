package smartid

import (
	"context"

	"github.com/tab/smartid/internal/models"
	"github.com/tab/smartid/internal/requests"
)

// Authenticate sends an authentication request to the Smart-ID provider
func (c *Client) Authenticate(context context.Context, nationalIdentityNumber string) (*models.AuthenticationResponse, error) {
	result, err := requests.Call(context, c.httpClient, c.config, nationalIdentityNumber)
	if err != nil {
		return nil, err
	}

	return result, nil
}
