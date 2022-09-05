package cashier_service

import (
	"cashier-service/internal/constants"
	"cashier-service/internal/models"
	"cashier-service/internal/structs"
)

func (s *CashierService) GetNoteCoinAmount() (cashInDrawer structs.CashStruct, err error) {
	var model []models.CashDrawer
	err = s.Find(&model).Error
	if err != nil {
		return cashInDrawer, err
	}

	for index := range model {
		switch model[index].Money_Value {
		case constants.OneThousandNote:
			cashInDrawer.OneThousand = model[index].Amount
		case constants.FiveHundredNote:
			cashInDrawer.FiveHundred = model[index].Amount
		case constants.OneHundredNote:
			cashInDrawer.OneHundred = model[index].Amount
		case constants.FiftyNote:
			cashInDrawer.Fifty = model[index].Amount
		case constants.TwentyNote:
			cashInDrawer.Twenty = model[index].Amount
		case constants.TenCoin:
			cashInDrawer.Ten = model[index].Amount
		case constants.FiveCoin:
			cashInDrawer.Five = model[index].Amount
		case constants.OneCoin:
			cashInDrawer.One = model[index].Amount
		case constants.TwentyFiveSatang:
			cashInDrawer.TwentyFiveSatang = model[index].Amount
		}
	}

	return cashInDrawer, nil
}
