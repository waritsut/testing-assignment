package cashier_transaction

import (
	"cashier-service/internal/services/cashier_service"
	"cashier-service/internal/structs"
)

type Transaction interface {
	CalTheChangeTransaction(req structs.RequestCalTheChangeMoney) (theChange, availableCash structs.CashStruct,
		balance float64, err error)
}

type CashierTransaction struct {
	cashierService cashier_service.Service
}

func NewCashierTransaction(cashierService cashier_service.Service) Transaction {
	return &CashierTransaction{
		cashierService: cashierService,
	}
}
