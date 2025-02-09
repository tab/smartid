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

// Result holds the authentication resultCh
type Result struct {
	Person *models.Person
	Err    error
}

// Error represents an error from the Smart-ID provider
type Error struct {
	Code string
}

// Error returns the error message
func (e *Error) Error() string {
	return fmt.Sprintf("authentication failed: %s", e.Code)
}

// Authenticate makes an authentication with the Smart-ID provider
func (c *Client) Authenticate(
	ctx context.Context,
	nationalIdentityNumber string,
) (*models.Session, <-chan Result) {
	resultCh := make(chan Result, 1)

	session, err := requests.CreateAuthenticationSession(ctx, c.config, nationalIdentityNumber)
	if err != nil {
		resultCh <- Result{nil, err}
		return nil, resultCh
	}

	go func() {
		response, err := requests.FetchAuthenticationSession(ctx, c.config, session.Id)
		if err != nil {
			resultCh <- Result{nil, err}
			return
		}

		switch response.State {
		case Running:
			resultCh <- Result{nil, errors.ErrAuthenticationIsRunning}
		case Complete:
			switch response.Result.EndResult {
			case OK:
				person, err := utils.Extract(response.Cert.Value)
				if err != nil {
					resultCh <- Result{nil, err}
					return
				}

				resultCh <- Result{person, nil}
			case USER_REFUSED,
				USER_REFUSED_DISPLAYTEXTANDPIN,
				USER_REFUSED_VC_CHOICE,
				USER_REFUSED_CONFIRMATIONMESSAGE,
				USER_REFUSED_CONFIRMATIONMESSAGE_WITH_VC_CHOICE,
				USER_REFUSED_CERT_CHOICE,
				WRONG_VC,
				TIMEOUT:

				resultCh <- Result{nil, &Error{Code: response.Result.EndResult}}
			default:
				resultCh <- Result{nil, errors.ErrUnsupportedResult}
			}
		default:
			resultCh <- Result{nil, errors.ErrUnsupportedState}
		}
	}()

	return session, resultCh
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
