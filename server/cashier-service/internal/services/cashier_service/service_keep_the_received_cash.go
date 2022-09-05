package cashier_service

import (
	"cashier-service/internal/constants"
	"cashier-service/internal/models"
	"cashier-service/internal/structs"
)

func (s *CashierService) KeepTheReceivedCash(receivedCash structs.CashStruct) (cashInDrawer structs.CashStruct, err error) {
	var model []models.CashDrawer
	err = s.Find(&model).Error
	if err != nil {
		return cashInDrawer, err
	}

	for index := range model {
		switch model[index].Money_Value {
		case constants.OneThousandNote:
			model[index].Amount += receivedCash.OneThousand
			cashInDrawer.OneThousand = model[index].Amount
		case constants.FiveHundredNote:
			model[index].Amount += receivedCash.FiveHundred
			cashInDrawer.FiveHundred = model[index].Amount
		case constants.OneHundredNote:
			model[index].Amount += receivedCash.OneHundred
			cashInDrawer.OneHundred = model[index].Amount
		case constants.FiftyNote:
			model[index].Amount += receivedCash.Fifty
			cashInDrawer.Fifty = model[index].Amount
		case constants.TwentyNote:
			model[index].Amount += receivedCash.Twenty
			cashInDrawer.Twenty = model[index].Amount
		case constants.TenCoin:
			model[index].Amount += receivedCash.Ten
			cashInDrawer.Ten = model[index].Amount
		case constants.FiveCoin:
			model[index].Amount += receivedCash.Five
			cashInDrawer.Five = model[index].Amount
		case constants.OneCoin:
			model[index].Amount += receivedCash.One
			cashInDrawer.One = model[index].Amount
		case constants.TwentyFiveSatang:
			model[index].Amount += receivedCash.TwentyFiveSatang
			cashInDrawer.TwentyFiveSatang = model[index].Amount
		}

	}

	err = s.Save(&model).Error
	if err != nil {
		return cashInDrawer, err
	}

	return cashInDrawer, nil
}
