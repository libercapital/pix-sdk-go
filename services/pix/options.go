package pix

import "time"

type ListPixOption struct {
	ApplyFunc func(parameter *ListPixParameter)
}

func StartDate(date time.Time) ListPixOption {
	return ListPixOption{
		ApplyFunc: func(parameter *ListPixParameter) {
			parameter.StartDate = date
		},
	}
}

func EndDate(date time.Time) ListPixOption {
	return ListPixOption{
		ApplyFunc: func(parameter *ListPixParameter) {
			parameter.EndDate = date
		},
	}
}

func TxId(txId string) ListPixOption {
	return ListPixOption{
		ApplyFunc: func(parameter *ListPixParameter) {
			parameter.TxId = txId
		},
	}
}

func HasTxId(hasTxId bool) ListPixOption {
	return ListPixOption{
		ApplyFunc: func(parameter *ListPixParameter) {
			parameter.HasTxId = hasTxId
		},
	}
}

func HasDevolution(devolution bool) ListPixOption {
	return ListPixOption{
		ApplyFunc: func(parameter *ListPixParameter) {
			parameter.HasDevolution = devolution
		},
	}
}

func CPF(cpf string) ListPixOption {
	return ListPixOption{
		ApplyFunc: func(parameter *ListPixParameter) {
			parameter.Cpf = cpf
		},
	}
}

func CNPJ(cnpj string) ListPixOption {
	return ListPixOption{
		ApplyFunc: func(parameter *ListPixParameter) {
			parameter.Cnpj = cnpj
		},
	}
}

func ActualPage(page int32) ListPixOption {
	return ListPixOption{
		ApplyFunc: func(parameter *ListPixParameter) {
			parameter.ActualPage = page
		},
	}
}

func ItensPerPage(itens int32) ListPixOption {
	return ListPixOption{
		ApplyFunc: func(parameter *ListPixParameter) {
			parameter.ItensPerPage = itens
		},
	}
}
