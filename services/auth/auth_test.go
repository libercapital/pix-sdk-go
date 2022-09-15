package auth

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	_ "embed"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk.git/bank"
	"gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk.git/errors"
	bankMocks "gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk.git/mocks/bank"
	servicesMocks "gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk.git/mocks/services"
)

func Test_Context_Authorize(t *testing.T) {

	type fields struct {
		bank        *bankMocks.Bank
		baseService *servicesMocks.BaseService
	}

	type args struct {
		authUrl      string
		clientId     string
		clientSecret string
	}

	tests := []struct {
		name         string
		fields       fields
		args         args
		wantErr      error
		want         bank.Authorization
		mockBehavior func(fields, *args)
	}{
		{
			name: "successful authorization",
			fields: fields{
				bank:        &bankMocks.Bank{},
				baseService: &servicesMocks.BaseService{},
			},
			args: args{
				authUrl:      "https://api.pix.com.br/oauth/token",
				clientId:     "pix-client-id",
				clientSecret: "pix-client-secret",
			},
			wantErr: nil,
			want: bank.Authorization{
				AccessToken: "access-token",
				TokenType:   "bearer",
				Scope:       "email profile",
				ExpiresIn:   1000,
				ExpireDate:  time.Now().Add(time.Second * 1000).Truncate(time.Second),
			},
			mockBehavior: func(f fields, a *args) {
				f.baseService.On("SetAuthorizer", mock.Anything)
				f.bank.On("GetAuthorization").
					Return(bank.Authorization{})
				f.bank.On("GetAuthUrl").
					Return(a.authUrl)
				f.bank.On("GetCredentials").
					Return(bank.Credentials{ClientId: a.clientId, ClientSecret: a.clientSecret})

				var bodyRequest = url.Values{"grant_type": {"client_credentials"}, "scope": {"pix.read"}}
				request, _ := http.NewRequest("POST", a.authUrl, bytes.NewBuffer([]byte(bodyRequest.Encode())))
				request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				request.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", a.clientId, a.clientSecret)))))
				f.baseService.On("CreateRequest", a.authUrl, "POST", bytes.NewBuffer([]byte(bodyRequest.Encode()))).
					Return(request, nil)

				f.baseService.On("Execute", request, mock.Anything, mock.Anything, mock.Anything).
					Return(&http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBuffer([]byte(`{"access_token": "access-token", "token_type": "bearer", "expires_in": 1000, "scope": "email profile"}`)))}, nil)

				f.bank.On("SetAuthorization", bank.Authorization{
					AccessToken: "access-token",
					TokenType:   "bearer",
					Scope:       "email profile",
					ExpiresIn:   1000,
				})
			},
		},
		{
			name: "successful with cache",
			fields: fields{
				bank:        &bankMocks.Bank{},
				baseService: &servicesMocks.BaseService{},
			},
			args:    args{},
			wantErr: nil,
			want: bank.Authorization{
				AccessToken: "access-token",
				TokenType:   "bearer",
				ExpiresIn:   1000,
				Scope:       "email profile",
				ExpireDate:  time.Now().Add(time.Second * 1000).Truncate(time.Second),
			},
			mockBehavior: func(f fields, a *args) {
				f.baseService.On("SetAuthorizer", mock.Anything)
				f.bank.On("GetAuthorization").
					Return(bank.Authorization{
						AccessToken: "access-token",
						TokenType:   "bearer",
						ExpiresIn:   1000,
						Scope:       "email profile",
						ExpireDate:  time.Now().Add(time.Second * 1000).Truncate(time.Second),
					})
			},
		},
		{
			name: "unauthorized response 401",
			fields: fields{
				bank:        &bankMocks.Bank{},
				baseService: &servicesMocks.BaseService{},
			},
			args: args{
				authUrl:      "https://api.pix.com.br/oauth/token",
				clientId:     "pix-client-id",
				clientSecret: "pix-client-secret",
			},
			wantErr: errors.ErrNotAuthorized,
			want:    bank.Authorization{},
			mockBehavior: func(f fields, a *args) {
				f.baseService.On("SetAuthorizer", mock.Anything)
				f.bank.On("GetAuthorization").
					Return(bank.Authorization{})
				f.bank.On("GetAuthUrl").
					Return(a.authUrl)
				f.bank.On("GetCredentials").
					Return(bank.Credentials{ClientId: a.clientId, ClientSecret: a.clientSecret})

				var bodyRequest = url.Values{"grant_type": {"client_credentials"}, "scope": {"pix.read"}}
				request, _ := http.NewRequest("POST", a.authUrl, bytes.NewBuffer([]byte(bodyRequest.Encode())))
				request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				request.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", a.clientId, a.clientSecret)))))
				f.baseService.On("CreateRequest", a.authUrl, "POST", bytes.NewBuffer([]byte(bodyRequest.Encode()))).
					Return(request, nil)

				f.baseService.On("Execute", request, mock.Anything, mock.Anything, mock.Anything).
					Return(&http.Response{StatusCode: http.StatusUnauthorized}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewAuthService(tt.fields.bank, tt.fields.baseService)
			if tt.mockBehavior != nil {
				tt.mockBehavior(tt.fields, &tt.args)
			}
			tt.fields.baseService.SetAuthorizer(s)
			authorization, err := s.Authorize()
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, authorization)

		})
	}
}
