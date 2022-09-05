package sequence_number_handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"reflect"
	mock_sequence_number_service "sequence-number-service/internal/services/sequence_number_service/mock"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func genVariableDataSet() (dataSet []int) {
	rand.Seed(time.Now().UnixNano())
	for index := 0; index < 10; index++ {
		dataSet = append(dataSet, rand.Intn(10))
	}
	return dataSet
}

func TestGetTheValueVariable(t *testing.T) {
	mockDataSet := genVariableDataSet()
	testCases := []struct {
		name          string
		buildStubs    func(mock *mock_sequence_number_service.MockService)
		checkResponse func(t *testing.T, httpStatus int, body interface{})
	}{
		{
			name: "OK",
			buildStubs: func(mock *mock_sequence_number_service.MockService) {
				mock.EXPECT().CalTheDataSet(gomock.Eq(uint64(8))).Times(1).Return(mockDataSet)
			},
			checkResponse: func(t *testing.T, httpStatus int, body interface{}) {
				require.Equal(t, http.StatusOK, httpStatus)
				requireVariableBodyMatch(t, body, mockDataSet)
			},
		},
	}

	for _, elem := range testCases {
		tc := elem
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := mock_sequence_number_service.NewMockService(ctrl)
			tc.buildStubs(mock)

			sequenceNumberHandler := NewController(mock)
			c := &TestVariableContext{}
			sequenceNumberHandler.GetTheValueVariable(c)

			tc.checkResponse(t, c.httpStatus, c.v)
		})
	}

}

func requireVariableBodyMatch(t *testing.T, body interface{}, expectedDataSet []int) {
	expectedRes := ResponseGetTheValueVariable{
		DataSet: expectedDataSet,
		XValue:  expectedDataSet[1],
		YValue:  expectedDataSet[4],
		ZValue:  expectedDataSet[5],
	}

	asserBody := body.(map[string]interface{})
	var resGot ResponseGetTheValueVariable
	bodyBytes, err := json.Marshal(asserBody["data"])
	require.NoError(t, err)
	json.Unmarshal(bodyBytes, &resGot)
	require.Equal(t, true, reflect.DeepEqual(resGot, expectedRes))
}

type TestVariableContext struct {
	httpStatus int
	v          interface{}
}

func (t *TestVariableContext) JSON(code int, v interface{}) {
	t.httpStatus = code
	t.v = v
}

func (t *TestVariableContext) Param(key string) string {
	return ""
}

func (t *TestVariableContext) BindJson(v interface{}) error {
	return nil
}

func (t *TestVariableContext) BindForm(obj interface{}) error {
	return nil
}
