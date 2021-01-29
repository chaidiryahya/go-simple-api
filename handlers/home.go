package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		response APIResponse
		start    = time.Now()
	)

	response.ServerProcessTime = time.Since(start).String()
	response.StatusCode = http.StatusOK
	response.Status = http.StatusText(http.StatusOK)
	response.Data = "Hello :)"

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
