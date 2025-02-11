package smartid

import (
	"context"
	"fmt"

	"github.com/tab/smartid/internal/errors"
	"github.com/tab/smartid/internal/models"
	"github.com/tab/smartid/internal/requests"
	"github.com/tab/smartid/internal/utils"
)

const (
	Running  = "RUNNING"
	Complete = "COMPLETE"

	OK                                              = "OK"
	USER_REFUSED                                    = "USER_REFUSED"
	USER_REFUSED_DISPLAYTEXTANDPIN                  = "USER_REFUSED_DISPLAYTEXTANDPIN"
	USER_REFUSED_VC_CHOICE                          = "USER_REFUSED_VC_CHOICE"
	USER_REFUSED_CONFIRMATIONMESSAGE                = "USER_REFUSED_CONFIRMATIONMESSAGE"
	USER_REFUSED_CONFIRMATIONMESSAGE_WITH_VC_CHOICE = "USER_REFUSED_CONFIRMATIONMESSAGE_WITH_VC_CHOICE"
	USER_REFUSED_CERT_CHOICE                        = "USER_REFUSED_CERT_CHOICE"
	WRONG_VC                                        = "WRONG_VC"
	TIMEOUT                                         = "TIMEOUT"
)

// Error represents an error from the Smart-ID provider
type Error struct {
	Code string
}

// Error returns the error message
func (e *Error) Error() string {
	return fmt.Sprintf("authentication failed: %s", e.Code)
}

// CreateSession creates authentication session with the Smart-ID provider
func (c *Client) CreateSession(ctx context.Context, nationalIdentityNumber string) (*models.Session, error) {
	session, err := requests.CreateAuthenticationSession(ctx, c.config, nationalIdentityNumber)
	if err != nil {
		return nil, err
	}

	return session, nil
}

// FetchSession fetches the authentication session from the Smart-ID provider
func (c *Client) FetchSession(ctx context.Context, sessionId string) (*Person, error) {
	response, err := requests.FetchAuthenticationSession(ctx, c.config, sessionId)
	if err != nil {
		return nil, err
	}

	switch response.State {
	case Running:
		return nil, errors.ErrAuthenticationIsRunning
	case Complete:
		switch response.Result.EndResult {
		case OK:
			person, err := utils.Extract(response.Cert.Value)
			if err != nil {
				return nil, err
			}

			return &Person{
				IdentityNumber: person.IdentityNumber,
				PersonalCode:   person.PersonalCode,
				FirstName:      person.FirstName,
				LastName:       person.LastName,
			}, nil
		case USER_REFUSED,
			USER_REFUSED_DISPLAYTEXTANDPIN,
			USER_REFUSED_VC_CHOICE,
			USER_REFUSED_CONFIRMATIONMESSAGE,
			USER_REFUSED_CONFIRMATIONMESSAGE_WITH_VC_CHOICE,
			USER_REFUSED_CERT_CHOICE,
			WRONG_VC,
			TIMEOUT:
			return nil, &Error{Code: response.Result.EndResult}
		}
	default:
		return nil, errors.ErrUnsupportedState
	}

	return nil, errors.ErrUnsupportedResult
}
