package pix

import (
	"fmt"

	"gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk.git/common"
	"gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk.git/errors"
	"gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk.git/services"
)

type Service interface {
	ListPix(listPix ListPix) (ListPixResponse, error)
	FindPix(e2eId string) (Pix, error)
}

type pixServiceContext struct {
	BaseService services.BaseService
}

func (p pixServiceContext) ListPix(listPix ListPix) (ListPixResponse, error) {
	var listPixParameter ListPixParameter
	listPix.Apply(&listPixParameter)
	var url = fmt.Sprintf("pix?%s", listPixParameter.ToQueryString())
	request, err := p.BaseService.CreateRequest(url, "GET", nil)
	if err != nil {
		return ListPixResponse{}, err
	}
	response, err := p.BaseService.Execute(request)
	if err != nil {
		return ListPixResponse{}, err
	}
	if response.StatusCode != 200 {
		switch response.StatusCode {
		case 403:
			return ListPixResponse{}, errors.ErrNotAllowed
		case 404:
			return ListPixResponse{}, errors.ErrEntityNotFound
		default:
			return ListPixResponse{}, errors.ErrUnknown
		}
	}
	var listPixResponse ListPixResponse
	if err := common.ParseBody(response.Body, &listPixResponse); err != nil {
		return ListPixResponse{}, err
	}
	return listPixResponse, nil
}

func (p pixServiceContext) FindPix(e2eId string) (Pix, error) {
	request, err := p.BaseService.CreateRequest(fmt.Sprintf("pix/%s", e2eId), "GET", nil)
	if err != nil {
		return Pix{}, err
	}
	response, err := p.BaseService.Execute(request)
	if err != nil {
		return Pix{}, err
	}
	if response.StatusCode != 200 {
		switch response.StatusCode {
		case 404:
			return Pix{}, errors.ErrPixNotFound
		case 403:
			return Pix{}, errors.ErrNotAllowed
		default:
			return Pix{}, errors.ErrUnknown
		}
	}
	var pix Pix
	if err := common.ParseBody(response.Body, &pix); err != nil {
		return Pix{}, err
	}
	return pix, nil
}

func NewPixService(baseService services.BaseService) Service {
	return &pixServiceContext{
		BaseService: baseService,
	}
}
