package structs

type CashStruct struct {
	OneThousand      uint `json:"oneThousandNote"`
	FiveHundred      uint `json:"fiveHundredNote"`
	OneHundred       uint `json:"oneHundredNote"`
	Fifty            uint `json:"fiftyNote"`
	Twenty           uint `json:"twentyNote"`
	Ten              uint `json:"tenCoin"`
	Five             uint `json:"fiveCoin"`
	One              uint `json:"oneCoin"`
	TwentyFiveSatang uint `json:"twentyFiveSatang"`
}
