package cashier_handler

import (
	mock_service "cashier-service/internal/services/cashier_service/mock"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestResetTheCashier(t *testing.T) {
	testCases := []struct {
		name          string
		buildStubs    func(mock *mock_service.MockService)
		checkResponse func(t *testing.T, httpStatus int, body interface{})
	}{
		{
			name: "OK",
			buildStubs: func(mock *mock_service.MockService) {
				mock.EXPECT().ResetTheCashier().Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, httpStatus int, body interface{}) {
				require.Equal(t, http.StatusOK, httpStatus)
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
			c := &TestResetTheCashierContext{}
			cashierHandler.ResetTheCashier(c)
			tc.checkResponse(t, c.httpStatus, c.v)
		})
	}

}

type TestResetTheCashierContext struct {
	httpStatus int

	v interface{}
}

func (t *TestResetTheCashierContext) BindJson(v interface{}) error {
	return nil
}

func (t *TestResetTheCashierContext) JSON(code int, v interface{}) {
	t.httpStatus = code
	t.v = v
}

func (t *TestResetTheCashierContext) Param(key string) string {
	return ""
}

func (t *TestResetTheCashierContext) BindForm(obj interface{}) error {
	return nil
}
