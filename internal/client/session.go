package client

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"

	"smartid/internal/errors"
	"smartid/internal/models"
)

// Status fetches the status of a Smart-ID session
func (c *Client) Status(id string) (*models.SessionResponse, error) {
	response, err := c.fetchSession(id)
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, errors.ErrSmartIdProviderError
	}

	var result models.SessionResponse
	if err = json.Unmarshal(response.Body(), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) fetchSession(id string) (*resty.Response, error) {
	endpoint := fmt.Sprintf("%s/session/%s", c.config.URL, id)

	response, err := c.httpClient.R().Get(endpoint)
	if err != nil {
		return nil, err
	}

	return response, nil
}
