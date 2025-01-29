package client

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/tab/smart-id/internal/errors"
	"github.com/tab/smart-id/internal/models"
)

// Authenticate sends an authentication request to the Smart-ID provider
func (c *Client) Authenticate(nationalIdentityNumber string) (*models.AuthenticateResponse, error) {
	hash, err := generateHash()
	if err != nil {
		return nil, err
	}

	response, err := c.createSession(nationalIdentityNumber, hash)
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, errors.ErrSmartIdProviderError
	}

	var result models.AuthenticateResponse
	if err = json.Unmarshal(response.Body(), &result); err != nil {
		return nil, err
	}

	code, err := generateCode(hash)
	if err != nil {
		return nil, err
	}

	return &models.AuthenticateResponse{
		SessionID: result.SessionID,
		Code:      code,
	}, nil
}

func (c *Client) createSession(nationalIdentityNumber, hash string) (*resty.Response, error) {
	body := c.requestPayload(nationalIdentityNumber, hash)
	endpoint := fmt.Sprintf("%s/authentication/etsi/%s", c.config.URL, nationalIdentityNumber)

	response, err := c.httpClient.R().SetBody(body).Post(endpoint)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) requestPayload(nationalIdentityNumber, hash string) models.AuthenticateRequest {
	return models.AuthenticateRequest{
		RelyingPartyName:       c.config.RelyingPartyName,
		RelyingPartyUUID:       c.config.RelyingPartyUUID,
		NationalIdentityNumber: nationalIdentityNumber,
		CertificateLevel:       c.config.CertificateLevel,
		Hash:                   hash,
		HashType:               c.config.HashType,
		AllowedInteractionsOrder: []models.AllowedInteraction{
			{
				Type:          c.config.InteractionType,
				DisplayText60: c.config.Text,
			},
		},
	}
}

func generateHash() (string, error) {
	randBytes := make([]byte, 64)
	_, err := rand.Read(randBytes)
	if err != nil {
		return "", err
	}

	hash := sha512.Sum512(randBytes)
	encoded := base64.StdEncoding.EncodeToString(hash[:])

	return encoded, nil
}

func generateCode(hash string) (string, error) {
	decodedHash, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return "", err
	}

	sha256Hash := sha256.Sum256(decodedHash)
	lastTwoBytes := sha256Hash[len(sha256Hash)-2:]
	codeInt := binary.BigEndian.Uint16(lastTwoBytes)
	vc := codeInt % 10000
	code := fmt.Sprintf("%04d", vc)

	return code, nil
}
