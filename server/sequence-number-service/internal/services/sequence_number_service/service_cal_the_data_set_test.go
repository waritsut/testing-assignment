package sequence_number_service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalTheDataSet(t *testing.T) {
	testCases := []struct {
		name          string
		sizeInput     uint64
		checkResponse func(t *testing.T, dataSet []int)
	}{
		{
			name:      "questionOnesolve",
			sizeInput: 8,
			checkResponse: func(t *testing.T, dataSet []int) {
				require.Equal(t, []int{1, 3, 8, 17, 31, 51, 78, 113}, dataSet)
			},
		},
		{
			name:      "10InputSize",
			sizeInput: 10,
			checkResponse: func(t *testing.T, dataSet []int) {
				require.Equal(t, []int{1, 3, 8, 17, 31, 51, 78, 113, 157, 211}, dataSet)
			},
		},
		{
			name:      "20InputSize",
			sizeInput: 20,
			checkResponse: func(t *testing.T, dataSet []int) {
				require.Equal(t, []int{1, 3, 8, 17, 31, 51, 78, 113, 157, 211, 276, 353,
					443, 547, 666, 801, 953, 1123, 1312, 1521}, dataSet)
			},
		},
	}

	for _, elem := range testCases {
		tc := elem
		t.Run(tc.name, func(t *testing.T) {
			sequenceNumberService := NewSequenceNumberService()
			dataSet := sequenceNumberService.CalTheDataSet(tc.sizeInput)
			tc.checkResponse(t, dataSet)
		})

	}
}
