package charge

import (
	"fmt"
	"time"
)

type CreateChargeOption struct {
	ApplyFunc  func(charge *Charge)
}

func TxId(txId string) *CreateChargeOption {
	return &CreateChargeOption{ApplyFunc: func(charge *Charge) {
		charge.TxId = txId
	}}
}

func Description(description string) *CreateChargeOption {
	return &CreateChargeOption{ApplyFunc: func(charge *Charge) {
		charge.Description = description
	}}
}

func Debtor(debtor DebtorCharge) *CreateChargeOption {
	return &CreateChargeOption{ApplyFunc: func(charge *Charge) {
		charge.Debtor = &debtor
	}}
}

func Value(value float64) *CreateChargeOption {
	return &CreateChargeOption{ApplyFunc: func(charge *Charge) {
		if charge.Value == nil {
			charge.Value = &ValueCharge{}
		}
		charge.Value.Origin = fmt.Sprintf("%.2f", value)
	}}
}

func Key(key string) *CreateChargeOption {
	return &CreateChargeOption{ApplyFunc: func(charge *Charge) {
		charge.Key = key
	}}
}

func AddInfo(name, value string) *CreateChargeOption {
	return &CreateChargeOption{ApplyFunc: func(charge *Charge) {
		charge.AdditionalInfo = append(charge.AdditionalInfo, &AdditionalInfoCharge{Value: value, Name: name})
	}}
}

func Validate(dueDate time.Time, maxOverdueDats int32) *CreateChargeOption{
	return &CreateChargeOption{ApplyFunc: func(charge *Charge) {
		if charge.Calendar == nil {
			charge.Calendar = &CalendarCharge{}
		}
		year, month, day := dueDate.Date()
		charge.Calendar.DueDate = fmt.Sprintf("%d-%d-%d", year, month, day)
		charge.Calendar.MaxOverdueDays = maxOverdueDats
	}}
}

func Expiration(seconds int32) *CreateChargeOption {
	return &CreateChargeOption{ApplyFunc: func(charge *Charge) {
		if charge.Calendar == nil {
			charge.Calendar = &CalendarCharge{}
		}
		charge.Calendar.Expiration = seconds
	}}
}

type FindChargeOption struct {
	ApplyFunc       func(request *FindChargeRequest)
}

func Revision(revision int64) *FindChargeOption {
	return &FindChargeOption{ApplyFunc: func(request *FindChargeRequest) {
		request.Revision = revision
	}}
}

type ListChargesOption struct {
	ApplyFunc       func(request *ListChargesRequest)
}

func Start(start time.Time) *ListChargesOption {
	return &ListChargesOption{ApplyFunc: func(request *ListChargesRequest) {
		request.Start = start.Format(time.RFC3339)
	}}
}

func End(end time.Time) *ListChargesOption {
	return &ListChargesOption{ApplyFunc: func(request *ListChargesRequest) {
		request.End = end.Format(time.RFC3339)
	}}
}

func CPF(cpf string) *ListChargesOption {
	return &ListChargesOption{ApplyFunc: func(request *ListChargesRequest) {
		request.CPF = cpf
	}}
}

func Status(status ChargeStatus) *ListChargesOption {
	return &ListChargesOption{ApplyFunc: func(request *ListChargesRequest) {
		request.Status = string(status)
	}}
}

func Page(page int32) *ListChargesOption {
	return &ListChargesOption{ApplyFunc: func(request *ListChargesRequest) {
		request.Page = page
	}}
}

func ItemsPerPage(itemsPerPage int32) *ListChargesOption {
	return &ListChargesOption{ApplyFunc: func(request *ListChargesRequest) {
		request.ItemsPerPage = itemsPerPage
	}}
}