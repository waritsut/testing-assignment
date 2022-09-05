package structs

type ResponseGetBalance struct {
	Balance float64    `json:"balance"`
	Cash    CashStruct `json:"cash"`
}
