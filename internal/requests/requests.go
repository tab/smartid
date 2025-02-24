package requests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/tab/smartid/internal/config"
	"github.com/tab/smartid/internal/errors"
	"github.com/tab/smartid/internal/models"
	"github.com/tab/smartid/internal/utils"
)

const (
	MaxIdleConnections        = 10000
	MaxIdleConnectionsPerHost = 10000
	IdleConnTimeout           = 90 * time.Second
	TLSHandshakeTimeout       = 10 * time.Second
	Timeout                   = 60 * time.Second

	StatusNoSuitableAccount = 471
	StatusViewSmartIdApp    = 472
	StatusClientTooOld      = 480
	StatusSystemMaintenance = 580
)

type Response struct {
	Id   string `json:"sessionID"`
	Code string `json:"code"`
}

func CreateAuthenticationSession(
	ctx context.Context,
	cfg *config.Config,
	identity string,
) (*Response, error) {
	hash, err := utils.GenerateHash(cfg.HashType)
	if err != nil {
		return nil, err
	}

	body := models.AuthenticationRequest{
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
	response, err := httpClient(cfg).R().SetContext(ctx).SetBody(body).Post(endpoint)
	if err != nil {
		return nil, err
	}

	if response.IsSuccess() {
		var result Response
		if err = json.Unmarshal(response.Body(), &result); err != nil {
			return nil, err
		}

		code, err := utils.GenerateVerificationCode(hash)
		if err != nil {
			return nil, err
		}

		return &Response{
			Id:   result.Id,
			Code: code,
		}, nil
	}

	switch response.StatusCode() {
	case StatusNoSuitableAccount:
		return nil, errors.ErrSmartIdNoSuitableAccount
	case StatusViewSmartIdApp:
		return nil, errors.ErrSmartIdViewApp
	case StatusClientTooOld:
		return nil, errors.ErrSmartIdClientTooOld
	case StatusSystemMaintenance:
		return nil, errors.ErrSmartIdMaintenance
	default:
		return nil, errors.ErrSmartIdProviderError
	}
}

func FetchAuthenticationSession(
	ctx context.Context,
	cfg *config.Config,
	sessionId string,
) (*models.AuthenticationResponse, error) {
	endpoint := fmt.Sprintf("%s/session/%s", cfg.URL, sessionId)

	response, err := httpClient(cfg).R().SetContext(ctx).Get(endpoint)
	if err != nil {
		return nil, err
	}

	if response.IsSuccess() {
		var result models.AuthenticationResponse
		if err = json.Unmarshal(response.Body(), &result); err != nil {
			return nil, err
		}

		return &result, nil
	}

	switch response.StatusCode() {
	case http.StatusNotFound:
		return nil, errors.ErrSmartIdSessionNotFound
	case StatusNoSuitableAccount:
		return nil, errors.ErrSmartIdNoSuitableAccount
	case StatusViewSmartIdApp:
		return nil, errors.ErrSmartIdViewApp
	case StatusClientTooOld:
		return nil, errors.ErrSmartIdClientTooOld
	case StatusSystemMaintenance:
		return nil, errors.ErrSmartIdMaintenance
	default:
		return nil, errors.ErrSmartIdProviderError
	}
}

func httpClient(cfg *config.Config) *resty.Client {
	transport := &http.Transport{
		MaxIdleConns:        MaxIdleConnections,
		MaxIdleConnsPerHost: MaxIdleConnectionsPerHost,
		IdleConnTimeout:     IdleConnTimeout,
		TLSHandshakeTimeout: TLSHandshakeTimeout,
		TLSClientConfig:     cfg.TLSConfig,
	}

	client := resty.NewWithClient(&http.Client{
		Transport: transport,
		Timeout:   Timeout,
	})

	client.
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json")

	return client
}
