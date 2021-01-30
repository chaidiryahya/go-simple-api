package handlers

import (
	"net/http"
	"time"

	"github.com/chaidiryahya/go-simple-api/services"
	"github.com/chaidiryahya/go-simple-api/utils"
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

		response.ServerProcessTime = time.Since(start).String()
		response.Status = http.StatusText(http.StatusOK)

		if err != nil {
			response.MessageError = err.Error()
		} else {
			response.Data = data
		}

		utils.WriteJsonResponse(w, response, http.StatusOK)

	}()

	data, err = services.SampleGet()

}
