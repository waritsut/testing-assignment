package sequence_number_handler

import (
	"net/http"
	"sequence-number-service/internal/pkg/router"
	"sequence-number-service/internal/response"
	"strconv"
)

type ResponseGetTheValueByIndex struct {
	DataSet []int `json:"dataSet"`
	Value   int   `json:"value"`
}

func (ctrl Controller) GetTheValueByIndex(c router.Context) {
	index := c.Param("index")
	intIndex, err := strconv.ParseUint(index, 10, 64)
	if err != nil {
		response.ErrorJson(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	dataSet := ctrl.sequenceNumberService.CalTheDataSet(intIndex + 1)
	res := ResponseGetTheValueByIndex{
		DataSet: dataSet,
		Value:   dataSet[intIndex],
	}
	response.WriteJson(c, http.StatusOK, "success", res)
}
