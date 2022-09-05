package logger

import (
	"cashier-service/event"
	"cashier-service/internal/globals"
	"encoding/json"
)

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func LogEventViaRabbit(payload string) {
	logPayLoad := LogPayload{
		Name: "sequence-number-service",
		Data: payload,
	}
	err := pushToQueue(logPayLoad.Name, logPayLoad.Data)
	if err != nil {
		return
	}
}

func pushToQueue(name, msg string) error {
	emitter, err := event.NewEventEmitter(globals.Rabbit)
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
