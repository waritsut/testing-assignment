package cashier_service

import (
	"cashier-service/internal/structs"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalTheChangeMoney(t *testing.T) {
	testCases := []struct {
		name          string
		itemPrice     float64
		receivedCash  float64
		cashInDrawer  structs.CashStruct
		checkResponse func(t *testing.T, theChange structs.CashStruct, err error)
	}{
		{
			name:         "NormalCase",
			itemPrice:    200,
			receivedCash: 1000,
			cashInDrawer: structs.CashStruct{FiveHundred: 10, OneHundred: 10},
			checkResponse: func(t *testing.T, theChange structs.CashStruct, err error) {
				require.Equal(t, structs.CashStruct{FiveHundred: 1, OneHundred: 3}, theChange)
				require.NoError(t, err)
			},
		},
		{
			name:         "OneHundredNoteNotEnoughCase",
			itemPrice:    200,
			receivedCash: 1000,
			cashInDrawer: structs.CashStruct{OneHundred: 5, Twenty: 5, Ten: 100},
			checkResponse: func(t *testing.T, theChange structs.CashStruct, err error) {
				require.Equal(t, structs.CashStruct{OneHundred: 5, Twenty: 5, Ten: 20}, theChange)
				require.NoError(t, err)
			},
		},
		{
			name:         "NoteAndCoinNotEnoughForTheChangeCase",
			itemPrice:    10,
			receivedCash: 1000,
			cashInDrawer: structs.CashStruct{Ten: 10, One: 100},
			checkResponse: func(t *testing.T, theChange structs.CashStruct, err error) {
				require.Equal(t, errors.New("banknote and coin in the cash drawer are insufficient for the change"), err)
			},
		},
		{
			name:         "CanNotCalculate",
			itemPrice:    500.55,
			receivedCash: 100.99,
			cashInDrawer: structs.CashStruct{OneThousand: 100},
			checkResponse: func(t *testing.T, theChange structs.CashStruct, err error) {
				require.Equal(t, errors.New("cloud not calculate the change"), err)
			},
		},
	}

	for _, elem := range testCases {
		tc := elem
		t.Run(tc.name, func(t *testing.T) {
			cashierService := NewCashierService(nil)
			theChange, err := cashierService.CalTheChangeMoney(tc.itemPrice, tc.receivedCash, tc.cashInDrawer)
			tc.checkResponse(t, theChange, err)
		})
	}
}
