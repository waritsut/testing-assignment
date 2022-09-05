package sequence_number_handler

import (
	"net/http"
	"sequence-number-service/internal/pkg/router"
	"sequence-number-service/internal/response"
)

type ResponseGetTheValueVariable struct {
	DataSet []int `json:"dataSet"`
	XValue  int   `json:"xValue"`
	YValue  int   `json:"yValue"`
	ZValue  int   `json:"zValue"`
}

func (ctrl Controller) GetTheValueVariable(c router.Context) {
	dataSet := ctrl.sequenceNumberService.CalTheDataSet(8)
	res := ResponseGetTheValueVariable{
		DataSet: dataSet,
		XValue:  dataSet[1],
		YValue:  dataSet[4],
		ZValue:  dataSet[5],
	}

	response.WriteJson(c, http.StatusOK, "success", res)
}
