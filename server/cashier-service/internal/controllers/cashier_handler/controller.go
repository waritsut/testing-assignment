package cashier_handler

import (
	"cashier-service/internal/services/cashier_service"
	"cashier-service/internal/transactions/cashier_transaction"
)

type Controller struct {
	cashierTransaction cashier_transaction.Transaction
	cashierService     cashier_service.Service
}

func NewController(cashierTransaction cashier_transaction.Transaction,
	cashierService cashier_service.Service) *Controller {
	return &Controller{
		cashierTransaction: cashierTransaction,
		cashierService:     cashierService,
	}
}
