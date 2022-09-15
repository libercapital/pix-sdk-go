package charge

import (
	"time"
)

type ChargeStatus string

const (
	Active       ChargeStatus = "ATIVA"
	InProcessing              = "EM_PROCESSAMENTO"
	Completed                 = "CONCLUIDA"
	Denied                    = "NEGADA"

	RemovedFromCreator = "REMOVIDA_PELO_USUARIO_RECEBEDOR"
	RemovedFromPSP = "REMOVIDA_PELO_PSP"
)

type CreateCharge []*CreateChargeOption

func (c CreateCharge) Apply(charge *Charge) {
	for _, option := range c {
		option.ApplyFunc(charge)
	}
}

type FindChargeRequest struct {
	Revision        int64           `url:"revisao"`
}

type ListCharges []*ListChargesOption

func (c ListCharges) Apply(request *ListChargesRequest) {
	for _, option := range c {
		option.ApplyFunc(request)
	}
}

type ListChargesRequest struct {
	Start           string          `url:"inicio,omitempty"`
	End             string          `url:"fim,omitempty"`
	CPF             string          `url:"cpf,omitempty"`
	Status          string          `url:"status,omitempty"`
	Page            int32           `url:"paginacao.paginaAtual,omitempty"`
	ItemsPerPage    int32           `url:"paginacao.itensPorPagina,omitempty"`
}

type CalendarCharge struct {
	CreatedAt       *time.Time  `json:"criacao,omitempty"`
	Expiration      int32       `json:"expiracao,omitempty"`
	DueDate         string      `json:"dataDeVencimento,omitempty"`
	MaxOverdueDays  int32       `json:"validadeAposVencimento,omitempty"`
}

type DebtorCharge struct {
	CNPJ        string      `json:"cnpj,omitempty"`
	CPF         string      `json:"cpf,omitempty"`
	Name        string      `json:"nome,omitempty"`
	Street      string      `json:"logradouro,omitempty"`
	City        string      `json:"city,omitempty"`
	UF          string      `json:"uf,omitempty"`
	PostalCode  string      `json:"cep,omitempty"`
}

type PenaltyValue struct {
	Modality            string      `json:"modalidade,omitempty"`
	PercentageValue     string      `json:"valorPerc,omitempty"`
}

type FeesValue struct {
	Modality            string      `json:"modalidade,omitempty"`
	PercentageValue     string      `json:"valorPerc,omitempty"`
}

type FixedDateDiscount struct {
	Date                *time.Time          `json:"data,omitempty"`
	PercentageValue     string              `json:"valorPerc,omitempty"`
}

type DiscountValue struct {
	Modality            string               `json:"modalidade,omitempty"`
	FixedDateDiscount   []*FixedDateDiscount `json:"descontoDataFixa,omitempty"`
}

type ValueCharge struct {
	Origin               string         `json:"original,omitempty"`
	ModalityModification int            `json:"modalidadeAlteracao,omitempty"`
	Penalty              *PenaltyValue  `json:"multa,omitempty"`
	Fees                 *FeesValue     `json:"juros,omitempty"`
	Discount             *DiscountValue `json:"desconto,omitempty"`
}

type AdditionalInfoCharge struct {
	Name        string      `json:"nome"`
	Value       string      `json:"valor"`
}

type LocationCharge struct {
	Id          int64       `json:"id,omitempty"`
	Location    string      `json:"location,omitempty"`
	ChargeType  string      `json:"cob,omitempty"`
}

type Charge struct {
	Calendar       *CalendarCharge         `json:"calendario,omitempty"`
	TxId           string                  `json:"txid,omitempty"`
	Revision       int                     `json:"revisao,omitempty"`
	Loc            *LocationCharge         `json:"loc,omitempty"`
	Location       string                  `json:"location,omitempty"`
	Status         ChargeStatus            `json:"status,omitempty"`
	Debtor         *DebtorCharge           `json:"devedor,omitempty"`
	Value          *ValueCharge            `json:"valor,omitempty"`
	Key            string                  `json:"chave,omitempty"`
	Description    string                  `json:"solicitacaoPagador"`
	AdditionalInfo []*AdditionalInfoCharge `json:"infoAdicionais,omitempty"`
}

type PaginationParameter struct {
	AtPage          int64       `json:"paginaAtual"`
	ItemsPerPage    int64       `json:"itensPorPagina"`
	TotalPages      int64       `json:"quantidadeDePagina"`
	TotalItems      int64       `json:"quantidadeDeItens"`
}

type ListChargesParameters struct {
	StartDate       *time.Time              `json:"incio,omitempty"`
	EndDate         *time.Time              `json:"fim,omitempty"`
	Pagination      *PaginationParameter    `json:"paginacao"`
}

type Charges struct {
	Parameters      *ListChargesParameters      `json:"parametros"`
	Charges         []*Charge                   `json:"cobs"`
}