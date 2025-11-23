package recaptcha

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const verifyURL = "https://www.google.com/recaptcha/api/siteverify"

// Verifier defines the interface for verifying reCAPTCHA tokens
//
//go:generate mockgen -destination=../../internal/mocks/recaptcha/verifier.go -package=mocks backend/pkg/recaptcha Verifier
type Verifier interface {
	Verify(ctx context.Context, token string) (bool, error)
}

type GoogleVerifier struct {
	secretKey string
	client    *http.Client
	disabled  bool
}

type siteVerifyResponse struct {
	Success     bool     `json:"success"`
	ChallengeTS string   `json:"challenge_ts"` // Timestamp of the challenge load
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes"`
}

func NewVerifier(secretKey string, disabled bool) *GoogleVerifier {
	return &GoogleVerifier{
		secretKey: secretKey,
		client:    &http.Client{Timeout: 10 * time.Second},
		disabled:  disabled,
	}
}

func (v *GoogleVerifier) Verify(ctx context.Context, token string) (bool, error) {
	if v.disabled {
		return true, nil
	}
	if token == "" {
		return false, nil
	}

	data := url.Values{}
	data.Set("secret", v.secretKey)
	data.Set("response", token)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, verifyURL, nil)
	if err != nil {
		return false, err
	}
	req.URL.RawQuery = data.Encode()

	resp, err := v.client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var result siteVerifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	return result.Success, nil
}

