package smartid

import (
	"context"

	"github.com/tab/smartid/internal/errors"
	"github.com/tab/smartid/internal/models"
	"github.com/tab/smartid/internal/requests"
	"github.com/tab/smartid/internal/utils"
)

// Authenticate sends an authentication request to the Smart-ID provider
func (c *Client) Authenticate(context context.Context, nationalIdentityNumber string) (*models.Person, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}

	result, err := requests.Call(context, c.httpClient, c.config, nationalIdentityNumber)
	if err != nil {
		return nil, err
	}

	person, err := utils.Extract(result.Cert.Value)
	if err != nil {
		return nil, err
	}

	return person, nil
}

func (c *Client) Validate() error {
	if c.config.RelyingPartyName == "" {
		return errors.ErrMissingRelyingPartyName
	}

	if c.config.RelyingPartyUUID == "" {
		return errors.ErrMissingRelyingPartyUUID
	}

	return nil
}
