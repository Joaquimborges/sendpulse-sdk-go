package sendpulse_sdk_go

import (
	"context"
	"fmt"
	"net/http"
)

// ValidatorService is a service to validate email addresses
type ValidatorService struct {
	client *Client
}

// newValidatorService creates ValidatorService
func newValidatorService(cl *Client) *ValidatorService {
	return &ValidatorService{client: cl}
}

// ValidateEmail verifies one email address
func (service *ValidatorService) ValidateEmail(ctx context.Context, email string) error {
	path := "/verifier-service/send-single-to-verify/"
	var response struct {
		Result bool `json:"result"`
	}
	type bodyFormat struct {
		Email string `json:"email"`
	}
	body := bodyFormat{Email: email}
	_, err := service.client.newRequest(ctx, http.MethodPost, path, body, &response, true)
	return err
}

// EmailValidationResult describes a result of a verification of specific email
type EmailValidationResult struct {
	Email  string `json:"email"`
	Checks struct {
		Status      int    `json:"status"`
		ValidFormat int    `json:"valid_format"`
		Disposable  int    `json:"disposable"`
		Webmail     int    `json:"webmail"`
		Gibberish   int    `json:"gibberish"`
		StatusText  string `json:"status_text"`
	} `json:"checks"`
}

// GetEmailValidationResult returns the results of a verification of specific email
func (service *ValidatorService) GetEmailValidationResult(ctx context.Context, email string) (*EmailValidationResult, error) {
	path := fmt.Sprintf("/verifier-service/get-single-result/?email=%s", email)
	var response struct {
		Result bool                   `json:"result"`
		Data   *EmailValidationResult `json:"data"`
	}
	_, err := service.client.newRequest(ctx, http.MethodGet, path, nil, &response, true)
	return response.Data, err
}

// DeleteEmailValidationResult removes the result of checking one address
func (service *ValidatorService) DeleteEmailValidationResult(ctx context.Context, email string) error {
	path := "/verifier-service/delete-single-result"
	var response struct {
		Result bool `json:"result"`
	}
	type bodyFormat struct {
		Email string `json:"email"`
	}
	body := bodyFormat{Email: email}
	_, err := service.client.newRequest(ctx, http.MethodGet, path, body, &response, true)
	return err
}
