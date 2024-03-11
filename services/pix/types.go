package pix

import (
	"time"

	"github.com/google/go-querystring/query"
	"github.com/libercapital/pix-sdk-go/common"
)

type DevolutionTime struct {
	Request     common.PixTime `json:"solitacao"`
	Liquidation common.PixTime `json:"liquidacao"`
}

type Devolution struct {
	Id       string         `json:"id"`
	ReturnId string         `json:"rtrId"`
	Value    float64        `json:"valor"`
	Time     DevolutionTime `json:"horario"`
	Status   string         `json:"status"`
}

type Pix struct {
	E2EId      string       `json:"endToEndId"`
	TxId       string       `json:"txid"`
	Value      string       `json:"valor"`
	Time       time.Time    `json:"horario"`
	Key        string       `json:"chave"`
	PayerInfo  string       `json:"infoPagador"`
	Devolution []Devolution `json:"devolucoes"`
	Components *Component   `json:"componentesValor"`
}

type Component struct {
	Original  *ValueComponentValue `json:"original,omitempty"`
	Penalty   *ValueComponentValue `json:"multa,omitempty"`
	Discount  *ValueComponentValue `json:"desconto,omitempty"`
	Interest  *ValueComponentValue `json:"juros,omitempty"`
	Deduction *ValueComponentValue `json:"abatimento,omitempty"`
}

type ValueComponentValue struct {
	Amount string `json:"valor"`
}

type ListPixResponse struct {
	Parameters *ListPixParameterResponse `json:"parametros"`
	Pix        []Pix                     `json:"pix"`
}

type ListPixParameterResponse struct {
	StartDate  common.PixTime                     `json:"inicio"`
	EndDate    common.PixTime                     `json:"fim"`
	Pagination ListPixParameterPaginationResponse `json:"paginacao"`
}

type ListPixParameterPaginationResponse struct {
	ActualPage   int32 `json:"paginaAtual"`
	ItemsPerPage int32 `json:"itensPorPagina"`
	TotalPages   int32 `json:"quantidadeDePaginas"`
	TotalItems   int32 `json:"quantidadeTotalDeItens"`
}

type ListPixParameter struct {
	StartDate     time.Time `url:"inicio" layout:"2006-01-02T15:04:05Z"`
	EndDate       time.Time `url:"fim" layout:"2006-01-02T15:04:05Z"`
	TxId          string    `url:"txid,omitempty"`
	HasTxId       bool      `url:"txidPresente,omitempty"`
	HasDevolution bool      `url:"devolucaoPresente,omitempty"`
	Cpf           string    `url:"cpf,omitempty"`
	Cnpj          string    `url:"cnpj,omitempty"`
	ActualPage    int32     `url:"paginacao.paginaAtual,omitempty"`
	ItensPerPage  int32     `url:"pagina.itensPorPagina,omitempty"`
}

func (l ListPixParameter) ToQueryString() string {
	urlValues, err := query.Values(l)
	if err != nil {
		return ""
	}
	return urlValues.Encode()
}

type ListPix []ListPixOption

func (l ListPix) Apply(parameter *ListPixParameter) {
	for _, option := range l {
		option.ApplyFunc(parameter)
	}
}
