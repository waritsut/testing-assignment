package cashier_service

import (
	"cashier-service/internal/pkg/db_driver"
	"cashier-service/internal/structs"
)

type Service interface {
	CalTheChangeMoney(itemPrice, receivedCash float64,
		cashInDrawer structs.CashStruct) (theChange structs.CashStruct, err error)
	GetBalance() (balance float64, err error)
	GetNoteCoinAmount() (cashInDrawer structs.CashStruct, err error)
	KeepTheReceivedCash(receivedCash structs.CashStruct) (cashInDrawer structs.CashStruct, err error)
	ResetTheCashier() (err error)
	TakeTheChange(theChange structs.CashStruct) (cashInDrawer structs.CashStruct, err error)
}

type CashierService struct {
	db_driver.Repo
}

func NewCashierService(db db_driver.Repo) Service {
	return &CashierService{
		Repo: db,
	}
}
