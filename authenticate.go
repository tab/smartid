package smartid

import (
	"context"

	"github.com/tab/smartid/internal/errors"
	"github.com/tab/smartid/internal/models"
	"github.com/tab/smartid/internal/requests"
	"github.com/tab/smartid/internal/utils"
)

// Result holds the authentication result
type Result struct {
	Person *models.Person
	Err    error
}

// Authenticate makes an authentication with the Smart-ID provider
func (c *Client) Authenticate(ctx context.Context, nationalIdentityNumber string) <-chan Result {
	result := make(chan Result, 1)

	go func() {
		session, err := requests.CreateAuthenticationSession(ctx, c.config, nationalIdentityNumber)
		if err != nil {
			result <- Result{nil, err}
			return
		}

		response, err := requests.FetchAuthenticationSession(ctx, c.config, session.SessionID)
		if err != nil {
			result <- Result{nil, err}
			return
		}

		person, err := utils.Extract(response.Cert.Value)
		if err != nil {
			result <- Result{nil, err}
			return
		}

		result <- Result{person, nil}
	}()

	return result
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
