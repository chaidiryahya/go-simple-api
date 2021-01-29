package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/chaidiryahya/go-simple-api/services"
	"github.com/julienschmidt/httprouter"
)

func SampleGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		err      error
		data     interface{}
		response APIResponse
		start    = time.Now()
	)

	defer func() {

		if err != nil {
			response.MessageError = err.Error()
		} else {
			response.Data = data
		}

		response.ServerProcessTime = time.Since(start).String()
		response.StatusCode = http.StatusOK
		response.Status = http.StatusText(http.StatusOK)

		json.NewEncoder(w).Encode(response)
	}()

	data, err = services.SampleGet()

}
