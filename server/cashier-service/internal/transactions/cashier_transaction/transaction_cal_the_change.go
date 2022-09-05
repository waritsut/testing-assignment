package cashier_transaction

import (
	"cashier-service/internal/globals"
	"cashier-service/internal/services/cashier_service"
	"cashier-service/internal/structs"

	"github.com/jinzhu/copier"
)

func (t *CashierTransaction) CalTheChangeTransaction(req structs.RequestCalTheChangeMoney) (theChange, availableCash structs.CashStruct,
	balance float64, err error) {
	var receivedCash structs.CashStruct
	copier.Copy(&receivedCash, req.Cash)

	tx := globals.DB.Begin()
	t.cashierService = cashier_service.NewCashierService(tx)

	cashInDrawer, err := t.cashierService.KeepTheReceivedCash(receivedCash)
	if err != nil {
		tx.Rollback()
		return theChange, availableCash, balance, err
	}

	theChange, err = t.cashierService.CalTheChangeMoney(req.ItemPrice, req.ReceivedCash, cashInDrawer)
	if err != nil {
		tx.Rollback()
		return theChange, availableCash, balance, err
	}

	availableCash, err = t.cashierService.TakeTheChange(theChange)
	if err != nil {
		tx.Rollback()
		return theChange, availableCash, balance, err
	}

	balance, err = t.cashierService.GetBalance()
	if err != nil {
		tx.Rollback()
		return theChange, availableCash, balance, err
	}
	tx.Commit()

	return theChange, availableCash, balance, nil
}
