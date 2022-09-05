package routes

import (
	"cashier-service/internal/controllers/cashier_handler"
	"cashier-service/internal/pkg/db_driver"
	"cashier-service/internal/pkg/router"
	"cashier-service/internal/services/cashier_service"
	"cashier-service/internal/transactions/cashier_transaction"
)

func NewRoute(db db_driver.Repo) router.Mux {
	r := router.NewMyRouter()

	cashierService := cashier_service.NewCashierService(db)
	cashierTransaction := cashier_transaction.NewCashierTransaction(cashierService)
	ctrl := cashier_handler.NewController(cashierTransaction, cashierService)
	cashierGroup := r.Group("cashiers")
	{
		cashierGroup.GET("", ctrl.GetBalance)
		cashierGroup.PATCH("/changes", ctrl.GetTheChangeMoney)
		cashierGroup.PUT("/resettings", ctrl.ResetTheCashier)
	}

	return r

}
