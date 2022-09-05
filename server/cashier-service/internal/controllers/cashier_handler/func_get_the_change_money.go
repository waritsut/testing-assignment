package cashier_handler

import (
	"cashier-service/internal/constants"
	"cashier-service/internal/pkg/router"
	"cashier-service/internal/response"
	"cashier-service/internal/structs"
	"errors"
	"fmt"
	"math"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func (ctrl Controller) GetTheChangeMoney(c router.Context) {
	var req structs.RequestCalTheChangeMoney
	validate = validator.New()
	validate.RegisterValidation("is-thai-baht", validateThaiBaht)
	validate.RegisterStructValidation(validationCashStructLevel, structs.RequestCalTheChangeMoney{})

	c.BindJson(&req)
	err := validateStruct(req)
	if err != nil {
		response.ErrorJson(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = checkSumIsMatchReceivedCash(req)
	if err != nil {
		response.ErrorJson(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	isOK := checkReceivedCashGreaterItem(req)
	if !isOK {
		response.ErrorJson(c, http.StatusBadRequest, "ReceivedCash must greater than or equal to ItemPrice", nil)
		return
	}

	theChange, availableCash, balance, err := ctrl.cashierTransaction.CalTheChangeTransaction(req)
	if err != nil {
		response.ErrorJson(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	res := structs.ResponseCalTheChangeMoney{
		AvailableCash: availableCash,
		Balance:       balance,
		Change:        math.Abs(req.ItemPrice - req.ReceivedCash),
		ChangeCash:    theChange,
	}

	response.WriteJson(c, http.StatusOK, "success", res)
}

var validate *validator.Validate

func validateStruct(req structs.RequestCalTheChangeMoney) error {
	err := validate.Struct(req)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return err
		}
		var errorText string
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				errorText = fmt.Sprintf("%s field is required", err.Field())
			case "min":
				errorText = fmt.Sprintf("%s field must greater than %s", err.Field(), err.Param())
			case "is-thai-baht":
				errorText = fmt.Sprintf("%s field must be Thai Baht value", err.Field())
			case "cash":
				errorText = fmt.Sprintf("money in %s struct field must valid", err.Field())
			}
		}
		return errors.New(errorText)
	}
	return nil
}

func validationCashStructLevel(sl validator.StructLevel) {
	cash := sl.Current().Interface().(structs.RequestCalTheChangeMoney)
	sumCash := cash.Cash.OneThousand + cash.Cash.FiveHundred + cash.Cash.OneHundred +
		cash.Cash.Fifty + cash.Cash.Twenty + cash.Cash.Ten + cash.Cash.Five +
		cash.Cash.One + cash.Cash.TwentyFiveSatang
	if sumCash == 0 {
		sl.ReportError(cash.Cash, "cash", "Cash", "cash", "")
	}
}

func validateThaiBaht(fl validator.FieldLevel) bool {
	return math.Mod(fl.Field().Float(), float64(0.25)) == 0
}

func checkSumIsMatchReceivedCash(req structs.RequestCalTheChangeMoney) (err error) {
	sumCash := (float64(req.Cash.OneThousand) * constants.OneThousandNote) + (float64(req.Cash.FiveHundred) * constants.FiveHundredNote) +
		(float64(req.Cash.OneHundred) * constants.OneHundredNote) + (float64(req.Cash.Fifty) * constants.FiftyNote) +
		(float64(req.Cash.Twenty) * constants.TwentyNote) + (float64(req.Cash.Ten) * constants.TenCoin) +
		(float64(req.Cash.Five) * constants.FiveCoin) + (float64(req.Cash.One) * constants.OneCoin) +
		(float64(req.Cash.TwentyFiveSatang) * constants.TwentyFiveSatang)
	if sumCash != req.ReceivedCash {
		return errors.New("ReceivedCash is not match with Cash")
	}
	return nil
}

func checkReceivedCashGreaterItem(req structs.RequestCalTheChangeMoney) bool {
	return req.ReceivedCash >= req.ItemPrice
}
