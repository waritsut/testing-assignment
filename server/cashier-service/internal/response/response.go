package response

import (
	"cashier-service/internal/logger"
	"cashier-service/internal/pkg/router"
	"encoding/json"
	"flag"
	"fmt"
)

func WriteJson(c router.Context, status int, message string, data interface{}) {
	payload := map[string]interface{}{
		"message": message,
		"error":   false,
		"data":    data,
	}

	if flag.Lookup("test.v") == nil {
		JsonByte, _ := json.Marshal(payload)
		logger.LogEventViaRabbit(fmt.Sprintf("res:%s", string(JsonByte)))
	}

	c.JSON(status, payload)
}

func ErrorJson(c router.Context, status int, message string, data interface{}) {
	payload := map[string]interface{}{
		"message": message,
		"error":   true,
		"data":    data,
	}

	if flag.Lookup("test.v") == nil {
		JsonByte, _ := json.Marshal(payload)
		logger.LogEventViaRabbit(fmt.Sprintf("res:%s", string(JsonByte)))
	}

	c.JSON(status, payload)
}
