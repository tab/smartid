package requests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/tab/smartid/config"
	"github.com/tab/smartid/internal/errors"
	"github.com/tab/smartid/internal/models"
	"github.com/tab/smartid/internal/utils"
)

func Call(context context.Context, client *resty.Client, config *config.Config, identity string) (*models.AuthenticationResponse, error) {
	result, err := createAuthenticationSession(context, client, config, identity)
	if err != nil {
		return nil, err
	}

	return fetchAuthenticationSession(context, client, config, result.SessionID)
}

func createAuthenticationSession(
	_ context.Context,
	httpClient *resty.Client,
	cfg *config.Config,
	identity string,
) (*models.AuthenticationSessionResponse, error) {
	hash, err := utils.GenerateHash(cfg.HashType)
	if err != nil {
		return nil, err
	}

	body := models.AuthenticationSessionRequest{
		RelyingPartyName:       cfg.RelyingPartyName,
		RelyingPartyUUID:       cfg.RelyingPartyUUID,
		NationalIdentityNumber: identity,
		CertificateLevel:       cfg.CertificateLevel,
		Hash:                   hash,
		HashType:               cfg.HashType,
		AllowedInteractionsOrder: []models.AllowedInteraction{
			{
				Type:          cfg.InteractionType,
				DisplayText60: cfg.Text,
			},
		},
	}

	endpoint := fmt.Sprintf("%s/authentication/etsi/%s", cfg.URL, identity)
	response, err := httpClient.R().SetBody(body).Post(endpoint)
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusOK {
		return nil, errors.ErrSmartIdProviderError
	}

	var result models.AuthenticationSessionResponse
	if err = json.Unmarshal(response.Body(), &result); err != nil {
		return nil, err
	}

	code, err := utils.VerificationCode(hash)
	if err != nil {
		return nil, err
	}

	return &models.AuthenticationSessionResponse{
		SessionID: result.SessionID,
		Code:      code,
	}, nil
}

func fetchAuthenticationSession(
	_ context.Context,
	httpClient *resty.Client,
	cfg *config.Config,
	sessionId string,
) (*models.AuthenticationResponse, error) {
	endpoint := fmt.Sprintf("%s/session/%s", cfg.URL, sessionId)

	response, err := httpClient.R().Get(endpoint)
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusOK {
		return nil, errors.ErrSmartIdProviderError
	}

	var result models.AuthenticationResponse
	if err = json.Unmarshal(response.Body(), &result); err != nil {
		return nil, err
	}

	return &result, nil
}
