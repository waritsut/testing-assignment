package cashier_handler

import (
	"cashier-service/internal/structs"
	mock_transaction "cashier-service/internal/transactions/cashier_transaction/mock"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetTheChangeMoney(t *testing.T) {
	testCases := []struct {
		name          string
		reqInput      structs.RequestCalTheChangeMoney
		buildStubs    func(mock *mock_transaction.MockTransaction)
		checkResponse func(t *testing.T, httpStatus int, body interface{})
	}{
		{
			name: "OK",
			reqInput: structs.RequestCalTheChangeMoney{
				ItemPrice:    500,
				ReceivedCash: 1000,
				Cash: structs.CashStruct{
					OneThousand: 1,
				},
			},
			buildStubs: func(mock *mock_transaction.MockTransaction) {
				mock.EXPECT().CalTheChangeTransaction(gomock.Any()).Times(1).Return(
					structs.CashStruct{FiveHundred: 1}, structs.CashStruct{OneThousand: 10}, float64(10000), nil)
			},
			checkResponse: func(t *testing.T, httpStatus int, body interface{}) {
				require.Equal(t, http.StatusOK, httpStatus)
				requireGetTheChangeBodyMatch(t, body, structs.CashStruct{OneThousand: 10},
					structs.CashStruct{FiveHundred: 1}, float64(10000), 500)
			},
		},

		{
			name: "BadRequestFieldRequired",
			reqInput: structs.RequestCalTheChangeMoney{
				ItemPrice:    0,
				ReceivedCash: 1000,
				Cash: structs.CashStruct{
					OneThousand: 1,
				},
			},
			buildStubs: func(mock *mock_transaction.MockTransaction) {
				mock.EXPECT().CalTheChangeTransaction(gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, httpStatus int, body interface{}) {
				err := body.(map[string]interface{})
				require.Equal(t, err["message"], "ItemPrice field is required")
				require.Equal(t, http.StatusBadRequest, httpStatus)
			},
		},

		{
			name: "BadRequestFieldMin",
			reqInput: structs.RequestCalTheChangeMoney{
				ItemPrice:    10,
				ReceivedCash: -5,
				Cash: structs.CashStruct{
					OneThousand: 1,
				},
			},
			buildStubs: func(mock *mock_transaction.MockTransaction) {
				mock.EXPECT().CalTheChangeTransaction(gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, httpStatus int, body interface{}) {
				err := body.(map[string]interface{})
				require.Equal(t, err["message"], "ReceivedCash field must greater than 0.25")
				require.Equal(t, http.StatusBadRequest, httpStatus)
			},
		},

		{
			name: "BadRequestInvalidThaiBaht",
			reqInput: structs.RequestCalTheChangeMoney{
				ItemPrice:    1000.99,
				ReceivedCash: 500,
				Cash: structs.CashStruct{
					OneThousand: 1,
				},
			},
			buildStubs: func(mock *mock_transaction.MockTransaction) {
				mock.EXPECT().CalTheChangeTransaction(gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, httpStatus int, body interface{}) {
				err := body.(map[string]interface{})
				require.Equal(t, err["message"], "ItemPrice field must be Thai Baht value")
				require.Equal(t, http.StatusBadRequest, httpStatus)
			},
		},

		{
			name: "BadRequestNotMatchCash",
			reqInput: structs.RequestCalTheChangeMoney{
				ItemPrice:    1000,
				ReceivedCash: 500,
				Cash: structs.CashStruct{
					OneThousand: 1,
				},
			},
			buildStubs: func(mock *mock_transaction.MockTransaction) {
				mock.EXPECT().CalTheChangeTransaction(gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, httpStatus int, body interface{}) {
				err := body.(map[string]interface{})
				require.Equal(t, err["message"], "ReceivedCash is not match with Cash")
				require.Equal(t, http.StatusBadRequest, httpStatus)
			},
		},

		{
			name: "BadRequestReceivedCashNotGreaterItem",
			reqInput: structs.RequestCalTheChangeMoney{
				ItemPrice:    1000,
				ReceivedCash: 500,
				Cash: structs.CashStruct{
					FiveHundred: 1,
				},
			},
			buildStubs: func(mock *mock_transaction.MockTransaction) {
				mock.EXPECT().CalTheChangeTransaction(gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, httpStatus int, body interface{}) {
				err := body.(map[string]interface{})
				require.Equal(t, err["message"], "ReceivedCash must greater than or equal to ItemPrice")
				require.Equal(t, http.StatusBadRequest, httpStatus)
			},
		},
	}

	for _, elem := range testCases {
		tc := elem
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := mock_transaction.NewMockTransaction(ctrl)
			tc.buildStubs(mock)

			cashierHandler := NewController(mock, nil)
			c := &TestGetTheChangeContext{
				reqBody: tc.reqInput,
			}
			cashierHandler.GetTheChangeMoney(c)
			tc.checkResponse(t, c.httpStatus, c.v)
		})
	}
}

func requireGetTheChangeBodyMatch(t *testing.T, body interface{},
	availableCash, changeCash structs.CashStruct, balance, change float64) {
	expectedRes := structs.ResponseCalTheChangeMoney{
		Change:        change,
		ChangeCash:    changeCash,
		AvailableCash: availableCash,
		Balance:       balance,
	}

	asserBody := body.(map[string]interface{})
	var resGot structs.ResponseCalTheChangeMoney
	bodyBytes, err := json.Marshal(asserBody["data"])
	require.NoError(t, err)
	json.Unmarshal(bodyBytes, &resGot)
	require.Equal(t, true, reflect.DeepEqual(resGot, expectedRes))
}

type TestGetTheChangeContext struct {
	httpStatus int
	reqBody    structs.RequestCalTheChangeMoney
	v          interface{}
}

func (t *TestGetTheChangeContext) BindJson(v interface{}) error {
	*v.(*structs.RequestCalTheChangeMoney) = t.reqBody
	return nil
}

func (t *TestGetTheChangeContext) JSON(code int, v interface{}) {
	t.httpStatus = code
	t.v = v
}

func (t *TestGetTheChangeContext) Param(key string) string {
	return ""
}

func (t *TestGetTheChangeContext) BindForm(obj interface{}) error {
	return nil
}
