package handlers

import (
	"net/http"
	"time"

	"github.com/chaidiryahya/go-simple-api/utils"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	var (
		response APIResponse
		start    = time.Now()
	)

	response.ServerProcessTime = time.Since(start).String()
	response.Status = http.StatusText(http.StatusNotFound)
	response.MessageError = "Endpoint not found :)"

	utils.WriteJsonResponse(w, response, http.StatusNotFound)
}
