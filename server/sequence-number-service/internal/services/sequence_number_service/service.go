package sequence_number_service

type Service interface {
	CalTheDataSet(size uint64) (output []int)
}

type SequenceNumberService struct {
}

func NewSequenceNumberService() Service {
	return &SequenceNumberService{}
}
