package bank

import (
	"fmt"
	"time"
)

var autorizationCache = make(map[string]Authorization)

type Bank interface {
	GetCredentials() Credentials
	GetAuthorization() Authorization
	SetAuthorization(authorization Authorization)
	GetCertificate() *Certificate
	GetUrl() string
	GetAuthUrl() string
}

type Context struct {
	BaseURL       string
	AuthURL       string
	Credentials   Credentials
	Authorization Authorization
	Certificate   *Certificate
}

func (c *Context) GetCertificate() *Certificate {
	return c.Certificate
}

func (c *Context) SetAuthorization(authorization Authorization) {
	c.Authorization = Authorization{
		AccessToken: authorization.AccessToken,
		TokenType:   authorization.TokenType,
		ExpiresIn:   authorization.ExpiresIn,
		Scope:       authorization.Scope,
		ExpireDate:  time.Now().Add(time.Second * time.Duration(authorization.ExpiresIn)).Truncate(time.Second),
	}
	autorizationCache[fmt.Sprintf("%s:%s@%s", c.Credentials.ClientId, c.Credentials.ClientSecret, c.BaseURL)] = c.Authorization
}

func (c *Context) GetUrl() string {
	return c.BaseURL
}

func (c *Context) GetAuthorization() Authorization {
	return c.Authorization
}

func (c *Context) GetCredentials() Credentials {
	return c.Credentials
}

func (c *Context) GetAuthUrl() string {
	return c.AuthURL
}

func NewBank(baseURL string, authURL string, cert *Certificate, credentials Credentials) Bank {
	var bank = Context{Credentials: credentials, BaseURL: baseURL, AuthURL: authURL, Certificate: cert}
	if auth, exists := autorizationCache[fmt.Sprintf("%s:%s@%s", credentials.ClientId, credentials.ClientSecret, baseURL)]; exists {
		bank.Authorization = auth
	}
	return &bank
}
