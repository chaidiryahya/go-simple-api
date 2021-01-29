package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	var (
		response APIResponse
		start    = time.Now()
	)

	response.ServerProcessTime = time.Since(start).String()
	response.StatusCode = http.StatusNotFound
	response.Status = http.StatusText(http.StatusNotFound)
	response.MessageError = "Endpoint not found :)"

	json.NewEncoder(w).Encode(response)
}
