package auth

import (
	"time"

	"gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk/bank"
	"gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk/common"
	"gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk/errors"
	"gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk/services"
)

type Service interface {
	Authorize() (bank.Authorization, error)
}

type Context struct {
	BaseService services.BaseService
	Bank        bank.Bank
}

func (c *Context) Authorize() (bank.Authorization, error) {
	if time.Now().Before(c.Bank.GetAuthorization().ExpireDate) {
		return c.Bank.GetAuthorization(), nil
	}
	body, err := common.ToUrlForm(NewAuthorizationRequest(ClientCredentials, PixRead))
	if err != nil {
		return bank.Authorization{}, err
	}
	request, err := c.BaseService.CreateRequest(c.Bank.GetAuthUrl(), "POST", body)
	if err != nil {
		return bank.Authorization{}, err
	}
	var credentials = c.Bank.GetCredentials()
	response, err := c.BaseService.Execute(request, common.WithUrlForm(), common.NoAuth(), common.WithBasicAuth(credentials.ClientId, credentials.ClientSecret))
	if err != nil {
		return bank.Authorization{}, err
	}
	if response.StatusCode == 401 {
		return bank.Authorization{}, errors.ErrNotAuthorized
	}
	var authorization bank.Authorization
	err = common.ParseBody(response.Body, &authorization)
	if err != nil {
		return bank.Authorization{}, err
	}
	c.Bank.SetAuthorization(authorization)
	return bank.Authorization{
		AccessToken: authorization.AccessToken,
		TokenType:   authorization.TokenType,
		Scope:       authorization.Scope,
		ExpiresIn:   authorization.ExpiresIn,
		ExpireDate:  time.Now().Add(time.Second * time.Duration(authorization.ExpiresIn)).Truncate(time.Second),
	}, err
}

func NewAuthService(bank bank.Bank, baseService services.BaseService) Service {
	var authService = Context{Bank: bank, BaseService: baseService}
	return &authService

}
