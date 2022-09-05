package cashier_service

import (
	"cashier-service/internal/constants"
	"cashier-service/internal/structs"
	"errors"
	"fmt"
	"math"
)

func (s *CashierService) CalTheChangeMoney(itemPrice, receivedCash float64,
	cashInDrawer structs.CashStruct) (theChange structs.CashStruct, err error) {
	noteExistStatus := map[string]interface{}{
		"OneThousandNote":  true,
		"FiveHundredNote":  true,
		"OneHundredNote":   true,
		"FiftyNote":        true,
		"TwentyNote":       true,
		"TenCoin":          true,
		"FiveCoin":         true,
		"OneCoin":          true,
		"TwentyFiveSatang": true,
	}

	changeValue := math.Abs(itemPrice - receivedCash)
	var isOk bool

	for !isOk {
		var diff float64
		if fmt.Sprintf("%v", noteExistStatus["OneThousandNote"]) == "true" && cashInDrawer.OneThousand != 0 {
			theChange.OneThousand = uint(changeValue) / uint(constants.OneThousandNote)
			diff = math.Mod(changeValue, constants.OneThousandNote)
		} else if fmt.Sprintf("%T", noteExistStatus["OneThousandNote"]) == "uint" {
			theChange.OneThousand = noteExistStatus["OneThousandNote"].(uint)
			if fmt.Sprintf("%v", noteExistStatus["OneThousandNote"]) != "0" {
				diff = diff - (float64(noteExistStatus["OneThousandNote"].(uint)) * constants.OneThousandNote)
			}
		} else {
			diff = changeValue
		}

		if diff != 0 && fmt.Sprintf("%v", noteExistStatus["FiveHundredNote"]) == "true" {
			theChange.FiveHundred = uint(diff) / uint(constants.FiveHundredNote)
			diff = math.Mod(diff, constants.FiveHundredNote)
		} else if fmt.Sprintf("%T", noteExistStatus["FiveHundredNote"]) == "uint" {
			theChange.FiveHundred = noteExistStatus["FiveHundredNote"].(uint)
			if fmt.Sprintf("%v", noteExistStatus["FiveHundredNote"]) != "0" {
				diff = diff - (float64(noteExistStatus["FiveHundredNote"].(uint)) * constants.FiveHundredNote)
			}
		}

		if diff != 0 && fmt.Sprintf("%v", noteExistStatus["OneHundredNote"]) == "true" {
			theChange.OneHundred = uint(diff) / uint(constants.OneHundredNote)
			diff = math.Mod(diff, constants.OneHundredNote)
		} else if fmt.Sprintf("%T", noteExistStatus["OneHundredNote"]) == "uint" {
			theChange.OneHundred = noteExistStatus["OneHundredNote"].(uint)
			if fmt.Sprintf("%v", noteExistStatus["OneHundredNote"]) != "0" {
				diff = diff - (float64(noteExistStatus["OneHundredNote"].(uint)) * constants.OneHundredNote)
			}
		}

		if diff != 0 && fmt.Sprintf("%v", noteExistStatus["FiftyNote"]) == "true" {
			theChange.Fifty = uint(diff) / uint(constants.FiftyNote)
			diff = math.Mod(diff, constants.FiftyNote)
		} else if fmt.Sprintf("%T", noteExistStatus["FiftyNote"]) == "uint" {
			theChange.Fifty = noteExistStatus["FiftyNote"].(uint)
			if fmt.Sprintf("%v", noteExistStatus["FiftyNote"]) != "0" {
				diff = diff - (float64(noteExistStatus["FiftyNote"].(uint)) * constants.FiftyNote)
			}
		}

		if diff != 0 && fmt.Sprintf("%v", noteExistStatus["TwentyNote"]) == "true" {
			theChange.Twenty = uint(diff) / uint(constants.TwentyNote)
			diff = math.Mod(diff, constants.TwentyNote)
		} else if fmt.Sprintf("%T", noteExistStatus["TwentyNote"]) == "uint" {
			theChange.Twenty = noteExistStatus["TwentyNote"].(uint)
			if fmt.Sprintf("%v", noteExistStatus["TwentyNote"]) != "0" {
				diff = diff - (float64(noteExistStatus["TwentyNote"].(uint)) * constants.TwentyNote)
			}
		}

		if diff != 0 && fmt.Sprintf("%v", noteExistStatus["TenCoin"]) == "true" {
			theChange.Ten = uint(diff) / uint(constants.TenCoin)
			diff = math.Mod(diff, constants.TenCoin)
		} else if fmt.Sprintf("%T", noteExistStatus["TenCoin"]) == "uint" {
			theChange.Ten = noteExistStatus["TenCoin"].(uint)
			if fmt.Sprintf("%v", noteExistStatus["TenCoin"]) != "0" {
				diff = diff - (float64(noteExistStatus["TenCoin"].(uint)) * constants.TenCoin)
			}
		}

		if diff != 0 && fmt.Sprintf("%v", noteExistStatus["FiveCoin"]) == "true" {
			theChange.Five = uint(diff) / uint(constants.FiveCoin)
			diff = math.Mod(diff, constants.FiveCoin)
		} else if fmt.Sprintf("%T", noteExistStatus["FiveCoin"]) == "uint" {
			theChange.Five = noteExistStatus["FiveCoin"].(uint)
			if fmt.Sprintf("%v", noteExistStatus["FiveCoin"]) != "0" {
				diff = diff - (float64(noteExistStatus["FiveCoin"].(uint)) * constants.FiveCoin)
			}
		}

		if diff != 0 && fmt.Sprintf("%v", noteExistStatus["OneCoin"]) == "true" {
			theChange.One = uint(diff) / uint(constants.OneCoin)
			diff = math.Mod(diff, constants.OneCoin)
		} else if fmt.Sprintf("%T", noteExistStatus["OneCoin"]) == "uint" {
			theChange.One = noteExistStatus["OneCoin"].(uint)
			if fmt.Sprintf("%v", noteExistStatus["OneCoin"]) != "0" {
				diff = diff - (float64(noteExistStatus["OneCoin"].(uint)) * constants.OneCoin)
			}
		}

		if diff != 0 && fmt.Sprintf("%v", noteExistStatus["TwentyFiveSatang"]) == "true" {
			satangCoin := diff / (constants.TwentyFiveSatang)
			theChange.TwentyFiveSatang = uint(satangCoin)
			diff = math.Mod(diff, constants.TwentyFiveSatang)
		} else if fmt.Sprintf("%T", noteExistStatus["TwentyFiveSatang"]) == "uint" {
			return theChange, errors.New("banknote and coin in the cash drawer are insufficient for the change")
		}

		if diff != 0 {
			return theChange, errors.New("cloud not calculate the change")
		}

		isOk = checkCashIsEnough(theChange, cashInDrawer, noteExistStatus)
	}

	return theChange, nil
}

func checkCashIsEnough(expectedChange, cashInDrawer structs.CashStruct, noteExistStatus map[string]interface{}) bool {
	if expectedChange.OneThousand > cashInDrawer.OneThousand && fmt.Sprintf("%v", noteExistStatus["OneThousandNote"]) == "true" {
		noteExistStatus["OneThousandNote"] = cashInDrawer.OneThousand
		return false
	}
	if expectedChange.FiveHundred > cashInDrawer.FiveHundred && fmt.Sprintf("%v", noteExistStatus["FiveHundredNote"]) == "true" {
		noteExistStatus["FiveHundredNote"] = cashInDrawer.FiveHundred
		return false
	}
	if expectedChange.OneHundred > cashInDrawer.OneHundred && fmt.Sprintf("%v", noteExistStatus["OneHundredNote"]) == "true" {
		noteExistStatus["OneHundredNote"] = cashInDrawer.OneHundred
		return false
	}
	if expectedChange.Fifty > cashInDrawer.Fifty && fmt.Sprintf("%v", noteExistStatus["FiftyNote"]) == "true" {
		noteExistStatus["FiftyNote"] = cashInDrawer.Fifty
		return false
	}
	if expectedChange.Twenty > cashInDrawer.Twenty && fmt.Sprintf("%v", noteExistStatus["TwentyNote"]) == "true" {
		noteExistStatus["TwentyNote"] = cashInDrawer.Twenty
		return false
	}
	if expectedChange.Ten > cashInDrawer.Ten && fmt.Sprintf("%v", noteExistStatus["TenCoin"]) == "true" {
		noteExistStatus["TenCoin"] = cashInDrawer.Ten
		return false
	}
	if expectedChange.Five > cashInDrawer.Five && fmt.Sprintf("%v", noteExistStatus["FiveCoin"]) == "true" {
		noteExistStatus["FiveCoin"] = cashInDrawer.Five
		return false
	}
	if expectedChange.One > cashInDrawer.One && fmt.Sprintf("%v", noteExistStatus["OneCoin"]) == "true" {
		noteExistStatus["OneCoin"] = cashInDrawer.One
		return false
	}
	if expectedChange.TwentyFiveSatang > cashInDrawer.TwentyFiveSatang && fmt.Sprintf("%v", noteExistStatus["TwentyFiveSatang"]) == "true" {
		noteExistStatus["TwentyFiveSatang"] = cashInDrawer.TwentyFiveSatang
		return false
	}

	return true
}
