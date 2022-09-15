package common

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

type applyFunc func(r *http.Request)

type RequestConfig struct {
	NoAuth bool
}

type RequestOption struct {
	ApplyFunc func(r *http.Request, config *RequestConfig)
}

type Options []RequestOption

func WithJSON() RequestOption {
	return WithContentType("application/json")
}

func ToUrlForm(data any) (io.Reader, error) {
	queryValue, err := query.Values(data)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer([]byte(queryValue.Encode())), nil
}

func ToJSON(data any) (io.Reader, error) {
	dataMarshal, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(dataMarshal), nil
}

func WithUrlForm() RequestOption {
	return WithContentType("application/x-www-form-urlencoded")
}

func WithContentType(contentType string) RequestOption {
	return RequestOption{
		ApplyFunc: func(r *http.Request, config *RequestConfig) {
			r.Header.Add("Content-Type", contentType)
		},
	}
}

func WithAuthorization(value string) RequestOption {
	return RequestOption{
		ApplyFunc: func(r *http.Request, config *RequestConfig) {
			r.Header.Add("Authorization", value)
		},
	}
}

func WithBasicAuth(username string, password string) RequestOption {
	return RequestOption{
		ApplyFunc: func(r *http.Request, config *RequestConfig) {
			r.SetBasicAuth(username, password)
		},
	}
}

func RequestOptions(options Options) func(r *http.Request, config *RequestConfig) {
	return func(r *http.Request, config *RequestConfig) {
		for _, option := range options {
			option.ApplyFunc(r, config)
		}
	}
}

func NoAuth() RequestOption {
	return RequestOption{
		ApplyFunc: func(r *http.Request, config *RequestConfig) {
			config.NoAuth = true
		},
	}
}
