package services

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	"gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk.git/bank"
	"gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk.git/common"
)

type BaseService interface {
	CreateRequest(url string, method string, body io.Reader) (*http.Request, error)
	Execute(r *http.Request, options ...common.RequestOption) (*http.Response, error)
	SetAuthorizer(authorizer ServiceAuthorizer)
}

type ServiceAuthorizer interface {
	Authorize() (bank.Authorization, error)
}

type Context struct {
	Bank              bank.Bank
	ServiceAuthorizer ServiceAuthorizer
	Client            http.Client
}

func (c *Context) Execute(r *http.Request, options ...common.RequestOption) (*http.Response, error) {
	var config common.RequestConfig
	for _, option := range options {
		option.ApplyFunc(r, &config)
	}
	if !config.NoAuth && c.ServiceAuthorizer != nil {
		auth, err := c.ServiceAuthorizer.Authorize()
		if err != nil {
			return nil, err
		}
		r.Header.Set("Authorization", fmt.Sprintf("%s %s", auth.TokenType, auth.AccessToken))
	}
	response, err := c.Client.Do(r)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Context) CreateRequest(path string, method string, body io.Reader) (*http.Request, error) {
	var err error
	var url = path
	if !common.IsURL(url) {
		url = fmt.Sprintf("%s/%s", c.Bank.GetUrl(), path)
	}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (c *Context) SetAuthorizer(authorizer ServiceAuthorizer) {
	c.ServiceAuthorizer = authorizer
}

func NewBaseService(bank bank.Bank) BaseService {
	var httpClient http.Client
	if bank.GetCertificate() != nil {
		httpClient = http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					Certificates: []tls.Certificate{*bank.GetCertificate()},
				},
			},
		}
	}
	return &Context{Bank: bank, Client: httpClient}
}
