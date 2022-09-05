package structs

type RequestCalTheChangeMoney struct {
	ItemPrice    float64    `json:"itemPrice" validate:"required,min=0.25,is-thai-baht"`
	ReceivedCash float64    `json:"receivedCash" validate:"required,min=0.25,is-thai-baht"`
	Cash         CashStruct `json:"cash"`
}

type ResponseCalTheChangeMoney struct {
	Change        float64    `json:"change"`
	ChangeCash    CashStruct `json:"changeCash"`
	AvailableCash CashStruct `json:"availableCash"`
	Balance       float64    `json:"balance"`
}
