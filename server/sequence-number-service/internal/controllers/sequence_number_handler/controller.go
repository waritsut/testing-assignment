package sequence_number_handler

import "sequence-number-service/internal/services/sequence_number_service"

type Controller struct {
	sequenceNumberService sequence_number_service.Service
}

func NewController(sequenceNumberService sequence_number_service.Service) *Controller {
	return &Controller{
		sequenceNumberService: sequenceNumberService,
	}
}
