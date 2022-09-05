package routes

import (
	"sequence-number-service/internal/controllers/sequence_number_handler"
	"sequence-number-service/internal/pkg/router"
	"sequence-number-service/internal/services/sequence_number_service"
)

func NewRoute() router.Mux {
	r := router.NewMyRouter()

	sequenceNumberService := sequence_number_service.NewSequenceNumberService()
	ctrl := sequence_number_handler.NewController(sequenceNumberService)
	sequenceNumberGroup := r.Group("sequenceNumbers")
	{
		sequenceNumberGroup.GET("", ctrl.GetTheValueVariable)
		sequenceNumberGroup.GET("/:index", ctrl.GetTheValueByIndex)
	}

	return r

}
