package sequence_number_service

func (s *SequenceNumberService) CalTheDataSet(size uint64) (output []int) {
	diffLayerOne := 2
	diffLayerTwo := 3

	for index := 0; index < int(size); index++ {
		if index == 0 {
			output = append(output, 1)
			continue
		}

		if index > 1 {
			diffLayerOne += diffLayerTwo
			diffLayerTwo++
		}

		value := output[index-1] + diffLayerOne
		output = append(output, value)
	}

	return output

}
