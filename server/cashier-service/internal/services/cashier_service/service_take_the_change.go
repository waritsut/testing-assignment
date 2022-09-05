package cashier_service

import (
	"cashier-service/internal/constants"
	"cashier-service/internal/models"
	"cashier-service/internal/structs"
)

func (s *CashierService) TakeTheChange(theChange structs.CashStruct) (cashInDrawer structs.CashStruct, err error) {
	var model []models.CashDrawer
	err = s.Find(&model).Error
	if err != nil {
		return cashInDrawer, err
	}

	for index := range model {
		switch model[index].Money_Value {
		case constants.OneThousandNote:
			model[index].Amount -= theChange.OneThousand
			cashInDrawer.OneThousand = model[index].Amount
		case constants.FiveHundredNote:
			model[index].Amount -= theChange.FiveHundred
			cashInDrawer.FiveHundred = model[index].Amount
		case constants.OneHundredNote:
			model[index].Amount -= theChange.OneHundred
			cashInDrawer.OneHundred = model[index].Amount
		case constants.FiftyNote:
			model[index].Amount -= theChange.Fifty
			cashInDrawer.Fifty = model[index].Amount
		case constants.TwentyNote:
			model[index].Amount -= theChange.Twenty
			cashInDrawer.Twenty = model[index].Amount
		case constants.TenCoin:
			model[index].Amount -= theChange.Ten
			cashInDrawer.Ten = model[index].Amount
		case constants.FiveCoin:
			model[index].Amount -= theChange.Five
			cashInDrawer.Five = model[index].Amount
		case constants.OneCoin:
			model[index].Amount -= theChange.One
			cashInDrawer.One = model[index].Amount
		case constants.TwentyFiveSatang:
			model[index].Amount -= theChange.TwentyFiveSatang
			cashInDrawer.TwentyFiveSatang = model[index].Amount
		}
	}

	err = s.Save(&model).Error
	if err != nil {
		return cashInDrawer, err
	}

	return cashInDrawer, nil
}
