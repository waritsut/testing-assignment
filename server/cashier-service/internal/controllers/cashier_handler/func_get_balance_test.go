package cashier_handler

import (
	mock_service "cashier-service/internal/services/cashier_service/mock"
	"cashier-service/internal/structs"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetBalance(t *testing.T) {
	testCases := []struct {
		name          string
		buildStubs    func(mock *mock_service.MockService)
		checkResponse func(t *testing.T, httpStatus int, body interface{})
	}{
		{
			name: "OK",
			buildStubs: func(mock *mock_service.MockService) {
				mock.EXPECT().GetBalance().Times(1).Return(float64(1000), nil)
				mock.EXPECT().GetNoteCoinAmount().Times(1).Return(structs.CashStruct{OneHundred: 10}, nil)
			},
			checkResponse: func(t *testing.T, httpStatus int, body interface{}) {
				require.Equal(t, http.StatusOK, httpStatus)
				requireGetBalanceBodyMatch(t, body, structs.CashStruct{OneHundred: 10}, float64(1000))
			},
		},
	}

	for _, elem := range testCases {
		tc := elem
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := mock_service.NewMockService(ctrl)
			tc.buildStubs(mock)

			cashierHandler := NewController(nil, mock)
			c := &TestGetBalanceContext{}
			cashierHandler.GetBalance(c)
			tc.checkResponse(t, c.httpStatus, c.v)
		})
	}
}

func requireGetBalanceBodyMatch(t *testing.T, body interface{},
	cash structs.CashStruct, balance float64) {
	expectedRes := structs.ResponseGetBalance{
		Balance: balance,
		Cash:    cash,
	}

	asserBody := body.(map[string]interface{})
	var resGot structs.ResponseGetBalance
	bodyBytes, err := json.Marshal(asserBody["data"])
	require.NoError(t, err)
	json.Unmarshal(bodyBytes, &resGot)
	require.Equal(t, true, reflect.DeepEqual(resGot, expectedRes))
}

type TestGetBalanceContext struct {
	httpStatus int

	v interface{}
}

func (t *TestGetBalanceContext) BindJson(v interface{}) error {
	return nil
}

func (t *TestGetBalanceContext) JSON(code int, v interface{}) {
	t.httpStatus = code
	t.v = v
}

func (t *TestGetBalanceContext) Param(key string) string {
	return ""
}

func (t *TestGetBalanceContext) BindForm(obj interface{}) error {
	return nil
}
