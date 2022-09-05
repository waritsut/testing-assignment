package sequence_number_handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"reflect"
	mock_sequence_number_service "sequence-number-service/internal/services/sequence_number_service/mock"
	"strconv"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func genByIndexDataSet(size uint64) (dataSet []int) {
	rand.Seed(time.Now().UnixNano())
	for index := 0; index <= int(size); index++ {
		dataSet = append(dataSet, rand.Intn(10))
	}
	return dataSet
}

func TestGetTheValueByIndex(t *testing.T) {
	testCases := []struct {
		name          string
		indexInput    string
		buildStubs    func(mock *mock_sequence_number_service.MockService, index int, mockDataSet []int)
		checkResponse func(t *testing.T, httpStatus int, body interface{}, expectedIndexValue int, mockDataSet []int)
	}{
		{
			name:       "OK",
			indexInput: "5",
			buildStubs: func(mock *mock_sequence_number_service.MockService, index int, mockDataSet []int) {
				mock.EXPECT().CalTheDataSet(gomock.Eq(uint64(index + 1))).Times(1).Return(mockDataSet)
			},
			checkResponse: func(t *testing.T, httpStatus int, body interface{}, expectedIndexValue int, mockDataSet []int) {
				require.Equal(t, http.StatusOK, httpStatus)
				requireByIndexBodyMatch(t, body, mockDataSet, expectedIndexValue)
			},
		},
		{
			name:       "BadRequest",
			indexInput: "textString",
			buildStubs: func(mock *mock_sequence_number_service.MockService, index int, mockDataSet []int) {
				mock.EXPECT().CalTheDataSet(gomock.Eq(uint64(index + 1))).Times(0)
			},
			checkResponse: func(t *testing.T, httpStatus int, body interface{}, expectedIndexValue int, mockDataSet []int) {
				require.Equal(t, http.StatusBadRequest, httpStatus)
			},
		},
	}

	for _, elem := range testCases {
		tc := elem
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			intIndex, err := strconv.Atoi(tc.indexInput)
			if err != nil {
				intIndex = 0
			}

			mockDataSet := genByIndexDataSet(uint64(intIndex))
			mock := mock_sequence_number_service.NewMockService(ctrl)
			tc.buildStubs(mock, intIndex, mockDataSet)

			sequenceNumberHandler := NewController(mock)
			c := &TestByIndexContext{index: tc.indexInput}
			sequenceNumberHandler.GetTheValueByIndex(c)

			tc.checkResponse(t, c.httpStatus, c.v, intIndex, mockDataSet)
		})
	}
}

func requireByIndexBodyMatch(t *testing.T, body interface{}, expectedDataSet []int, expectedIndexValue int) {
	expectedRes := ResponseGetTheValueByIndex{
		DataSet: expectedDataSet,
		Value:   expectedDataSet[expectedIndexValue],
	}

	asserBody := body.(map[string]interface{})
	var resGot ResponseGetTheValueByIndex
	bodyBytes, err := json.Marshal(asserBody["data"])
	require.NoError(t, err)
	json.Unmarshal(bodyBytes, &resGot)

	require.Equal(t, true, reflect.DeepEqual(resGot, expectedRes))
}

type TestByIndexContext struct {
	httpStatus int
	v          interface{}

	index string
}

func (t *TestByIndexContext) JSON(code int, v interface{}) {
	t.httpStatus = code
	t.v = v
}

func (t *TestByIndexContext) Param(key string) string {
	return t.index
}

func (t *TestByIndexContext) BindJson(v interface{}) error {
	return nil
}

func (t *TestByIndexContext) BindForm(obj interface{}) error {
	return nil
}
