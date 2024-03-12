package services

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"

	liberlogger "github.com/libercapital/liber-logger-go"
	"github.com/libercapital/liber-logger-go/tracing"
	"github.com/libercapital/pix-sdk-go/bank"
	"github.com/libercapital/pix-sdk-go/common"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
)

const (
	defaultMaxIdleConns        = 100
	defaultMaxConnsPerHost     = 100
	defaultMaxIdleConnsPerHost = 100
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

func buildHttpClient() *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = defaultMaxIdleConns
	t.MaxConnsPerHost = defaultMaxConnsPerHost
	t.MaxIdleConnsPerHost = defaultMaxIdleConnsPerHost

	httpClient := &http.Client{
		Transport: liberlogger.HttpClient{
			Proxied:      t,
			RedactedKeys: liberlogger.DefaultKeys,
		},
	}

	config := tracing.HttpTraceConfig{
		OperationName: "bank-request",
		SpanType:      ext.SpanTypeHTTP,
		ResourceName: func(req *http.Request) string {
			return fmt.Sprintf("%s %s", req.Method, req.URL.String())
		},
	}

	return tracing.HttpTrace(httpClient, config)
}

func applyCertificate(rootCAs *x509.CertPool, cert tls.Certificate) http.RoundTripper {
	liberlogger.Info(context.TODO()).Interface("rootCas", &rootCAs).Interface("cert", cert).Msg("Applying certificate to the request")

	transport := liberlogger.HttpClient{
		Proxied: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:                rootCAs,
				Certificates:           []tls.Certificate{cert},
				InsecureSkipVerify:     true,
				SessionTicketsDisabled: true,
			},
		},
		RedactedKeys: liberlogger.DefaultKeys,
	}

	return transport
}

func NewBaseService(bank bank.Bank) BaseService {
	httpClient := buildHttpClient()

	if bank.GetCertificate() != nil {
		rootCAs, cert := bank.GetCertificate().Load()

		httpClient.Transport = applyCertificate(rootCAs, cert)
	}

	return &Context{Bank: bank, Client: *httpClient}
}
