package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/chaidiryahya/go-simple-api/services"
	"github.com/chaidiryahya/go-simple-api/utils"
	"github.com/julienschmidt/httprouter"
)

func SamplePostwithJsonBody(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		err      error
		data     interface{}
		response APIResponse
		start    = time.Now()
		param    services.SamplePostWithJsonBodyParam
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

	r.ParseForm()

	if err = json.NewDecoder(r.Body).Decode(&param); err != nil {
		err = errors.New("Bad request body")
		return
	}

	if param.YourName == "" {
		err = errors.New("Your Name cannot be empty")
		return
	}

	if param.UserID == 0 {
		err = errors.New("UserID cannot be empty")
		return
	}

	data, err = services.SamplePostWithJsonBody(param)
}
