package pixsdk

import (
	"context"
	"crypto/tls"
	"gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk.git/services/webhook"

	"gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk.git/bank"
	"gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk.git/services"
	"gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk.git/services/auth"
	"gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk.git/services/pix"
)

type Service interface {
	pix.Service
	webhook.Service
	SetConfig(config Config)
}

type Context struct {
	Config         Config
	Bank           bank.Bank
	BaseService    services.BaseService
	AuthService    auth.Service
	PixService     pix.Service
	WebhookService webhook.Service
}

type Config struct {
	AuthURL      string
	BaseURL      string
	ClientId     string
	ClientSecret string
	Certificate  *tls.Certificate
}

func (c *Context) CreateWebhook(ctx context.Context, key string, webhook webhook.CreateWebhook) error {
	return c.WebhookService.CreateWebhook(ctx, key, webhook)
}

func (c *Context) FindWebhook(ctx context.Context, key string) (webhook.Webhook, error) {
	return c.WebhookService.FindWebhook(ctx, key)
}

func (c *Context) DeleteWebhook(ctx context.Context, key string) error {
	return c.WebhookService.DeleteWebhook(ctx, key)
}

func (c *Context) FindPix(endToEndId string) (pix.Pix, error) {
	return c.PixService.FindPix(endToEndId)
}

func (c *Context) ListPix(listPix pix.ListPix) (pix.ListPixResponse, error) {
	return c.PixService.ListPix(listPix)
}

func (c *Context) setupConfig(config Config) {
	var bank = bank.NewBank(config.BaseURL, config.AuthURL, config.Certificate, bank.NewCredentials(config.ClientId, config.ClientSecret))
	var baseService = services.NewBaseService(bank)
	var authService = auth.NewAuthService(bank, baseService)
	baseService.SetAuthorizer(authService)
	var pixService = pix.NewPixService(baseService)
	var webhookService = webhook.NewWebhookService(baseService)

	c.Config = config
	c.Bank = bank
	c.BaseService = baseService
	c.AuthService = authService
	c.PixService = pixService
	c.WebhookService = webhookService
}

func (c *Context) SetConfig(config Config) {
	c.setupConfig(config)
}

func NewService(config Config) Service {
	var ctx = &Context{}
	ctx.setupConfig(config)
	return ctx
}
