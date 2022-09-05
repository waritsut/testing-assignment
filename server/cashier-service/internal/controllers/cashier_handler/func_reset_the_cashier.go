package cashier_handler

import (
	"cashier-service/internal/pkg/router"
	"cashier-service/internal/response"
	"net/http"
)

func (ctrl Controller) ResetTheCashier(c router.Context) {
	err := ctrl.cashierService.ResetTheCashier()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	response.WriteJson(c, http.StatusOK, "success", nil)

}
