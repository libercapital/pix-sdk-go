package charge

import (
	"fmt"

	"github.com/google/go-querystring/query"
	"gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk/common"
	"gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk/services"
)

type Service interface {
	CreateCharge(create CreateCharge) (*Charge, error)
	FindCharge(txId string, options ...*FindChargeOption) (*Charge, error)
	FindCharges(charges ListCharges) (*Charges, error)
}

type Context struct {
	BaseService services.BaseService
}

func (c *Context) CreateCharge(create CreateCharge) (*Charge, error) {
	var charge Charge
	create.Apply(&charge)
	var txId = charge.TxId
	charge.TxId = ""
	var url = "cob"
	var method = "POST"
	if charge.Calendar.DueDate != "" {
		url = "cobv"
	}
	if txId != "" {
		url = fmt.Sprintf("%s/%s", url, txId)
		method = "PUT"
	}
	body, err := common.ToJSON(charge)
	if err != nil {
		return nil, err
	}
	request, err := c.BaseService.CreateRequest(url, method, body)
	if err != nil {
		return nil, err
	}
	response, err := c.BaseService.Execute(request, common.WithJSON())
	if err != nil {
		return nil, err
	}
	var responseCharge Charge
	err = common.ParseBody(response.Body, &responseCharge)
	if err != nil {
		return nil, err
	}
	return &responseCharge, nil
}

func (c *Context) FindCharge(txId string, options ...*FindChargeOption) (*Charge, error) {
	var findRequest FindChargeRequest
	for _, option := range options {
		option.ApplyFunc(&findRequest)
	}
	queryValue, _ := query.Values(findRequest)

	var url = fmt.Sprintf("v2/cob/%s?%s", txId, queryValue.Encode())
	request, err := c.BaseService.CreateRequest(url, "GET", nil)
	if err != nil {
		return nil, err
	}
	response, err := c.BaseService.Execute(request, common.WithJSON())
	if err != nil {
		return nil, err
	}
	var charge Charge
	err = common.ParseBody(response.Body, &charge)
	if err != nil {
		return nil, err
	}
	return &charge, nil
}

func (c *Context) FindCharges(listCharges ListCharges) (*Charges, error) {
	var listRequest ListChargesRequest
	listCharges.Apply(&listRequest)
	queryValue, _ := query.Values(listRequest)
	var url = fmt.Sprintf("/cob?%s", queryValue.Encode())
	request, err := c.BaseService.CreateRequest(url, "GET", nil)
	if err != nil {
		return nil, err
	}
	response, err := c.BaseService.Execute(request)
	if err != nil {
		return nil, err
	}
	var charges Charges
	err = common.ParseBody(response.Body, &charges)
	if err != nil {
		return nil, err
	}
	return &charges, nil
}

func NewChargeService(baseService services.BaseService) Service {
	return &Context{BaseService: baseService}

}
