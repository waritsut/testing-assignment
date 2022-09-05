package main

import (
	"broker-service/event"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestPayload struct {
	Action        string               `json:"action"`
	SeqNumByIndex SeqNumByIndexPayload `json:"seqNum,omitempty"`
	Cashier       CashierPayload       `json:"cashier,omitempty"`
	Log           LogPayload           `json:"log,omitempty"`
}

type SeqNumByIndexPayload struct {
	Index int `json:"index"`
}

type CashierPayload struct {
	ItemPrice    float64 `json:"itemPrice"`
	ReceivedCash float64 `json:"receivedCash"`
	Cash         struct {
		OneThousand      uint `json:"oneThousandNote"`
		FiveHundred      uint `json:"fiveHundredNote"`
		OneHundred       uint `json:"oneHundredNote"`
		Fifty            uint `json:"fiftyNote"`
		Twenty           uint `json:"twentyNote"`
		Ten              uint `json:"tenCoin"`
		Five             uint `json:"fiveCoin"`
		One              uint `json:"oneCoin"`
		TwentyFiveSatang uint `json:"twentyFiveSatang"`
	} `json:"cash"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err.Error())
		return
	}

	app.logEventViaRabbit(requestPayload.Action)
	switch requestPayload.Action {
	case "seq-num":
		app.seqNum(w)
	case "seq-num-by-index":
		app.seqNumByIndex(w, requestPayload.SeqNumByIndex)
	case "cashier-get-balance":
		app.cashierGetBalance(w)
	case "cashier-cal-the-change":
		app.cashierCalTheChange(w, requestPayload.Cashier)
	case "cashier-reset":
		app.cashierReset(w)
	default:
		app.errorJSON(w, "unexpected action")
	}

}

func (app *Config) seqNum(w http.ResponseWriter) {
	jsonData, _ := json.MarshalIndent(nil, "", "\t")

	request, err := http.NewRequest("GET", "http://sequence-number-service/sequenceNumbers", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		log.Println("sequence-number-service OK")
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, jsonFromService.Message)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, jsonFromService.Message, response.StatusCode)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = jsonFromService.Message
	payload.Data = jsonFromService.Data
	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) seqNumByIndex(w http.ResponseWriter, s SeqNumByIndexPayload) {
	jsonData, _ := json.MarshalIndent(nil, "", "\t")

	request, err := http.NewRequest("GET", "http://sequence-number-service/sequenceNumbers/"+fmt.Sprint(s.Index), bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		log.Println("sequence-number-service OK")
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, jsonFromService.Message)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, jsonFromService.Message, response.StatusCode)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = jsonFromService.Message
	payload.Data = jsonFromService.Data
	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) cashierGetBalance(w http.ResponseWriter) {
	jsonData, _ := json.MarshalIndent(nil, "", "\t")

	request, err := http.NewRequest("GET", "http://cashier-service/cashiers", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		log.Println("cashier-service OK")
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, jsonFromService.Message)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, jsonFromService.Message, response.StatusCode)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = jsonFromService.Message
	payload.Data = jsonFromService.Data
	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) cashierCalTheChange(w http.ResponseWriter, c CashierPayload) {
	jsonData, _ := json.MarshalIndent(c, "", "\t")

	request, err := http.NewRequest("PATCH", "http://cashier-service/cashiers/changes", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		log.Println("cashier-service OK")
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, jsonFromService.Message)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, jsonFromService.Message, response.StatusCode)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = jsonFromService.Message
	payload.Data = jsonFromService.Data
	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) cashierReset(w http.ResponseWriter) {
	jsonData, _ := json.MarshalIndent(nil, "", "\t")

	request, err := http.NewRequest("PUT", "http://cashier-service/cashiers/resettings", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		log.Println("cashier-service OK")
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, jsonFromService.Message)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, jsonFromService.Message, response.StatusCode)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = jsonFromService.Message
	payload.Data = jsonFromService.Data
	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) logEventViaRabbit(action string) {
	logPayLoad := LogPayload{
		Name: action,
		Data: "access broker service",
	}
	err := app.pushToQueue(logPayLoad.Name, logPayLoad.Data)
	if err != nil {
		return
	}
}

func (app *Config) pushToQueue(name, msg string) error {
	emitter, err := event.NewEventEmitter(app.Rabbit)
	if err != nil {
		return err
	}

	payload := LogPayload{
		Name: name,
		Data: msg,
	}

	j, _ := json.MarshalIndent(&payload, "", "\t")
	err = emitter.Push(string(j), "log.INFO")
	if err != nil {
		return err
	}
	return nil
}
