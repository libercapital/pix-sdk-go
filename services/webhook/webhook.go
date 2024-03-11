package webhook

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	liberlogger "github.com/libercapital/liber-logger-go"
	"github.com/libercapital/pix-sdk-go/common"
	pixErrors "github.com/libercapital/pix-sdk-go/errors"
	"github.com/libercapital/pix-sdk-go/services"
)

type Service interface {
	CreateWebhook(ctx context.Context, key string, webhook CreateWebhook) error
	FindWebhook(ctx context.Context, key string) (Webhook, error)
	DeleteWebhook(ctx context.Context, key string) error
}

type webhookService struct {
	baseService services.BaseService
}

type violation struct {
	Reason   string `json:"razao,omitempty"`
	Property string `json:"propriedade,omitempty"`
}

type response struct {
	Type       string      `json:"type,omitempty"`
	Title      string      `json:"title,omitempty"`
	Status     uint16      `json:"status,omitempty"`
	Detail     string      `json:"detail,omitempty"`
	Violations []violation `json:"violacoes,omitempty"`
}

func NewWebhookService(baseService services.BaseService) Service {
	return &webhookService{baseService: baseService}
}

func (w webhookService) checkHttpStatus(statusCode int) error {
	if statusCode != http.StatusOK {
		switch statusCode {
		case http.StatusBadRequest:
			return pixErrors.ErrBadRequest
		case http.StatusForbidden:
			return pixErrors.ErrNotAllowed
		case http.StatusNotFound:
			return pixErrors.ErrEntityNotFound
		default:
			return pixErrors.ErrUnknown
		}
	}

	return nil
}

func (w webhookService) parseErrorResponse(httpResponse *http.Response) (response, error) {
	var parsedResponse response

	if err := common.ParseBody(httpResponse.Body, &parsedResponse); err != nil {
		return response{}, err
	}

	return parsedResponse, nil
}

func (w webhookService) CreateWebhook(ctx context.Context, key string, webhook CreateWebhook) error {
	url := fmt.Sprintf("webhook/%s", key)

	body, err := json.Marshal(webhook)

	if err != nil {
		return err
	}

	request, err := w.baseService.CreateRequest(url, http.MethodPut, bytes.NewReader(body))

	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := w.baseService.Execute(request)

	if err != nil {
		return err
	}

	err = w.checkHttpStatus(response.StatusCode)

	if err != nil {
		parsedResponse, errResponse := w.parseErrorResponse(response)

		if errResponse == nil {
			liberlogger.
				Warn(ctx).
				Interface("response", parsedResponse).
				Msg("error to create webhook")
		}
	}

	return err
}

func (w webhookService) FindWebhook(ctx context.Context, key string) (Webhook, error) {
	url := fmt.Sprintf("webhook/%s", key)

	request, err := w.baseService.CreateRequest(url, http.MethodGet, nil)

	if err != nil {
		return Webhook{}, err
	}

	response, err := w.baseService.Execute(request)

	if err != nil {
		return Webhook{}, err
	}

	err = w.checkHttpStatus(response.StatusCode)

	if err != nil {
		parsedResponse, errResponse := w.parseErrorResponse(response)

		if errResponse == nil {
			liberlogger.
				Warn(ctx).
				Interface("response", parsedResponse).
				Msg("error to find webhook")
		}

		return Webhook{}, err
	}

	var webhookResponse Webhook

	if err := common.ParseBody(response.Body, &webhookResponse); err != nil {
		return Webhook{}, err
	}

	return webhookResponse, nil
}

func (w webhookService) DeleteWebhook(ctx context.Context, key string) error {
	url := fmt.Sprintf("webhook/%s", key)

	request, err := w.baseService.CreateRequest(url, http.MethodDelete, nil)

	if err != nil {
		return err
	}

	response, err := w.baseService.Execute(request)

	if err != nil {
		return err
	}

	err = w.checkHttpStatus(response.StatusCode)

	if err != nil {
		parsedResponse, errResponse := w.parseErrorResponse(response)

		if errResponse == nil {
			liberlogger.
				Warn(ctx).
				Interface("response", parsedResponse).
				Msg("error to delete webhook")
		}
	}

	return err
}
