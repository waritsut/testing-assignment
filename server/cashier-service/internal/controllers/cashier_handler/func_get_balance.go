package cashier_handler

import (
	"cashier-service/internal/pkg/router"
	"cashier-service/internal/response"
	"cashier-service/internal/structs"
	"net/http"
)

func (ctrl Controller) GetBalance(c router.Context) {
	balance, err := ctrl.cashierService.GetBalance()
	if err != nil {
		response.ErrorJson(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	cashInDrawer, err := ctrl.cashierService.GetNoteCoinAmount()
	if err != nil {
		response.ErrorJson(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	res := structs.ResponseGetBalance{
		Balance: balance,
		Cash:    cashInDrawer,
	}

	response.WriteJson(c, http.StatusOK, "success", res)

}
