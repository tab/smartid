package requests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

	MinSmartIdTimeout = 1000
	MaxSmartIdTimeout = 120000

	CertificateLevelQUALIFIED = "QUALIFIED"

	InteractionTypeDisplayTextAndPIN                            = "displayTextAndPIN"
	InteractionTypeVerificationCodeChoice                       = "verificationCodeChoice"
	InteractionTypeConfirmationMessage                          = "confirmationMessage"
	InteractionTypeConfirmationMessageAndVerificationCodeChoice = "confirmationMessageAndVerificationCodeChoice"

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

	interaction := models.AllowedInteraction{
		Type: cfg.InteractionType,
	}

	switch cfg.InteractionType {
	case InteractionTypeDisplayTextAndPIN, InteractionTypeVerificationCodeChoice:
		interaction.DisplayText60 = cfg.DisplayText60
	case InteractionTypeConfirmationMessage, InteractionTypeConfirmationMessageAndVerificationCodeChoice:
		interaction.DisplayText200 = cfg.DisplayText200
	}

	body := models.AuthenticationRequest{
		RelyingPartyName:       cfg.RelyingPartyName,
		RelyingPartyUUID:       cfg.RelyingPartyUUID,
		NationalIdentityNumber: identity,
		CertificateLevel:       cfg.CertificateLevel,
		Hash:                   hash,
		HashType:               cfg.HashType,
		AllowedInteractionsOrder: []models.AllowedInteraction{
			interaction,
		},
	}

	endpoint := fmt.Sprintf("%s/authentication/etsi/%s", cfg.URL, identity)
	response, err := httpClient(cfg).R().SetContext(ctx).SetBody(body).Post(endpoint)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode() {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted:
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
	case http.StatusForbidden:
		return nil, errors.ErrSmartIdAccessForbidden
	case StatusNoSuitableAccount, http.StatusNotFound:
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

	timeout := int(cfg.Timeout.Milliseconds())

	switch {
	case timeout < MinSmartIdTimeout:
		timeout = MinSmartIdTimeout
	case timeout > MaxSmartIdTimeout:
		timeout = MaxSmartIdTimeout
	}

	response, err := httpClient(cfg).R().
		SetContext(ctx).
		SetQueryParam("timeoutMs", strconv.Itoa(timeout)).
		Get(endpoint)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode() {
	case http.StatusOK:
		var result models.AuthenticationResponse
		if err = json.Unmarshal(response.Body(), &result); err != nil {
			return nil, err
		}
		return &result, nil
	case http.StatusForbidden:
		return nil, errors.ErrSmartIdAccessForbidden
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
