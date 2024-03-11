package pix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/libercapital/pix-sdk-go/common"
	"github.com/libercapital/pix-sdk-go/errors"
	servicesMock "github.com/libercapital/pix-sdk-go/mocks/services"
	authMock "github.com/libercapital/pix-sdk-go/mocks/services/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type JSON map[string]any

func makeJson(data JSON) io.ReadCloser {
	content, _ := json.Marshal(data)
	return io.NopCloser(bytes.NewBuffer([]byte(content)))
}

func Test_Context_FindPix(t *testing.T) {

	type args struct {
		e2eId string
	}

	type fields struct {
		baseService *servicesMock.BaseService
		authService *authMock.Service
	}

	var now = time.Now().Truncate(time.Second).UTC()

	var tests = []struct {
		name         string
		fields       fields
		args         args
		wantErr      error
		want         Pix
		mockBehavior func(fields, *args)
	}{
		{
			name: "sucessful find pix",
			fields: fields{
				baseService: &servicesMock.BaseService{},
				authService: &authMock.Service{},
			},
			args: args{
				e2eId: "E7268236923237655FX723",
			},
			wantErr: nil,
			want: Pix{
				E2EId: "E7268236923237655FX723",
				TxId:  "TESTE0001",
				Value: "1",
				Time:  now,
				Key:   "pix@bavabank.com.br",
			},
			mockBehavior: func(f fields, a *args) {
				f.baseService.On("SetAuthorizer", mock.Anything)

				var url = fmt.Sprintf("pix/%s", a.e2eId)
				request, _ := http.NewRequest("GET", url, nil)
				f.baseService.On("CreateRequest", url, "GET", nil).
					Return(request, nil)

				f.baseService.On("Execute", request).
					Return(&http.Response{StatusCode: http.StatusOK, Body: makeJson(JSON{
						"endToEndId": a.e2eId,
						"txId":       "TESTE0001",
						"valor":      1,
						"horario":    now.Format(time.RFC3339),
						"chave":      "pix@bavabank.com.br",
					})}, nil)
			},
		},
		{
			name: "not found pix 404",
			fields: fields{
				baseService: &servicesMock.BaseService{},
				authService: &authMock.Service{},
			},
			args: args{
				e2eId: "E7268236923237655FX723",
			},
			wantErr: errors.ErrPixNotFound,
			want:    Pix{},
			mockBehavior: func(f fields, a *args) {
				f.baseService.On("SetAuthorizer", mock.Anything)

				var url = fmt.Sprintf("pix/%s", a.e2eId)
				request, _ := http.NewRequest("GET", url, nil)
				f.baseService.On("CreateRequest", url, "GET", nil).
					Return(request, nil)
				f.baseService.On("Execute", request).
					Return(&http.Response{
						StatusCode: http.StatusNotFound,
					}, nil)
			},
		},
		{
			name: "not allowed pix 403",
			args: args{
				e2eId: "E7268236923237655FX723",
			},
			fields: fields{
				baseService: &servicesMock.BaseService{},
				authService: &authMock.Service{},
			},
			wantErr: errors.ErrNotAllowed,
			want:    Pix{},
			mockBehavior: func(f fields, a *args) {
				f.baseService.On("SetAuthorizer", mock.Anything)

				var url = fmt.Sprintf("pix/%s", a.e2eId)
				request, _ := http.NewRequest("GET", url, nil)
				f.baseService.On("CreateRequest", url, "GET", nil).
					Return(request, nil)

				f.baseService.On("Execute", request).
					Return(&http.Response{
						StatusCode: http.StatusForbidden,
					}, nil)
			},
		},
		{
			name: "unknown error 500",
			args: args{
				e2eId: "E7268236923237655FX723",
			},
			fields: fields{
				baseService: &servicesMock.BaseService{},
				authService: &authMock.Service{},
			},
			wantErr: errors.ErrUnknown,
			want:    Pix{},
			mockBehavior: func(f fields, a *args) {
				f.baseService.On("SetAuthorizer", mock.Anything)

				var url = fmt.Sprintf("pix/%s", a.e2eId)
				request, _ := http.NewRequest("GET", url, nil)
				f.baseService.On("CreateRequest", url, "GET", nil).
					Return(request, nil)

				f.baseService.On("Execute", request).
					Return(&http.Response{
						StatusCode: http.StatusInternalServerError,
					}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pixService := NewPixService(tt.fields.baseService)
			if tt.mockBehavior != nil {
				tt.mockBehavior(tt.fields, &tt.args)
			}
			tt.fields.baseService.SetAuthorizer(tt.fields.authService)
			pix, err := pixService.FindPix(tt.args.e2eId)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, pix)
		})
	}
}

func Test_Context_ListPix(t *testing.T) {

	type args struct {
		startDate time.Time
		endDate   time.Time
	}

	type fields struct {
		authService *authMock.Service
		baseService *servicesMock.BaseService
	}
	var now = time.Now().Round(time.Second).UTC()

	var dateLayout = "2006-01-02T15:04:05Z"
	var dateLayout2 = "2006-01-02T15:04:05"

	var tests = []struct {
		name         string
		args         args
		fields       fields
		wantErr      error
		want         ListPixResponse
		mockBehavior func(fields, *args)
	}{
		{
			name: "successful list pix",
			args: args{
				startDate: now.AddDate(0, -1, 0),
				endDate:   now,
			},
			fields: fields{
				authService: &authMock.Service{},
				baseService: &servicesMock.BaseService{},
			},
			wantErr: nil,
			want: ListPixResponse{
				Parameters: &ListPixParameterResponse{
					StartDate: common.PixTime{Time: now.AddDate(0, -1, 0)},
					EndDate:   common.PixTime{Time: now},
					Pagination: ListPixParameterPaginationResponse{
						ActualPage:   0,
						ItemsPerPage: 100,
						TotalPages:   1,
						TotalItems:   1,
					},
				},
				Pix: []Pix{
					{
						E2EId: "E7268236923237655FX723",
						TxId:  "TESTE0001",
						Value: "1",
						Time:  now,
						Key:   "pix@bavabank.com.br",
					},
				},
			},
			mockBehavior: func(f fields, a *args) {
				f.baseService.On("SetAuthorizer", mock.Anything)

				var url = fmt.Sprintf("pix?fim=%s&inicio=%s", url.QueryEscape(a.endDate.Format(dateLayout)), url.QueryEscape(a.startDate.Format(dateLayout)))
				request, _ := http.NewRequest("GET", url, nil)
				f.baseService.On("CreateRequest", url, "GET", nil).
					Return(request, nil)

				f.baseService.On("Execute", request).
					Return(&http.Response{
						StatusCode: http.StatusOK,
						Body: makeJson(JSON{
							"parametros": JSON{
								"inicio": a.startDate.Format(dateLayout2),
								"fim":    a.startDate.Format(dateLayout2),
								"paginacao": JSON{
									"paginaAtual":            0,
									"itensPorPagina":         100,
									"quantidadeDePaginas":    1,
									"quantidadeTotalDeItens": 1,
								},
							},
							"pix": []JSON{
								{
									"pix": JSON{
										"endToEndId": "E7268236923237655FX723",
										"txId":       "TESTE0001",
										"valor":      1,
										"horario":    now.Format(dateLayout),
										"chave":      "pix@bavabank.com.br",
									},
								},
							},
						}),
					}, nil)
			},
		},
		{
			name: "not allowed list pix 403",
			args: args{
				startDate: now.AddDate(0, -1, 0),
				endDate:   now,
			},
			fields: fields{
				authService: &authMock.Service{},
				baseService: &servicesMock.BaseService{},
			},
			wantErr: errors.ErrNotAllowed,
			want:    ListPixResponse{},
			mockBehavior: func(f fields, a *args) {
				f.baseService.On("SetAuthorizer", mock.Anything)

				var url = fmt.Sprintf("pix?fim=%s&inicio=%s", url.QueryEscape(a.endDate.Format(dateLayout)), url.QueryEscape(a.startDate.Format(dateLayout)))
				request, _ := http.NewRequest("GET", url, nil)
				f.baseService.On("CreateRequest", url, "GET", nil).
					Return(request, nil)

				f.baseService.On("Execute", request).
					Return(&http.Response{
						StatusCode: 403,
					}, nil)

			},
		},
		{
			name: "unknown error list pix 500",
			args: args{
				startDate: now.AddDate(0, -1, 0),
				endDate:   now,
			},
			fields: fields{
				authService: &authMock.Service{},
				baseService: &servicesMock.BaseService{},
			},
			wantErr: errors.ErrUnknown,
			want:    ListPixResponse{},
			mockBehavior: func(f fields, a *args) {
				f.baseService.On("SetAuthorizer", mock.Anything)

				var url = fmt.Sprintf("pix?fim=%s&inicio=%s", url.QueryEscape(a.endDate.Format(dateLayout)), url.QueryEscape(a.startDate.Format(dateLayout)))
				request, _ := http.NewRequest("GET", url, nil)
				f.baseService.On("CreateRequest", url, "GET", nil).
					Return(request, nil)

				f.baseService.On("Execute", request).
					Return(&http.Response{
						StatusCode: 500,
					}, nil)

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pixService := NewPixService(tt.fields.baseService)
			if tt.mockBehavior != nil {
				tt.mockBehavior(tt.fields, &tt.args)
			}
			tt.fields.baseService.SetAuthorizer(tt.fields.authService)
			listPix, err := pixService.ListPix(ListPix{
				StartDate(tt.args.startDate),
				EndDate(tt.args.endDate),
			})
			if listPix.Parameters != nil {
				listPix.Parameters.StartDate = common.PixTime{Time: now.AddDate(0, -1, 0)}
				listPix.Parameters.EndDate = common.PixTime{Time: now}
			}
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, listPix)
		})
	}
}
