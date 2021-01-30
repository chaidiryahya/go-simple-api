package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chaidiryahya/go-simple-api/services"
	"github.com/chaidiryahya/go-simple-api/utils"
	"github.com/julienschmidt/httprouter"
)

func SampleGetWithRouteParam(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		err      error
		data     interface{}
		response APIResponse
		start    = time.Now()
		userID   int64
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

	userIDparam := ps.ByName("user_id")
	if userIDparam == "" {
		err = errors.New("UserID cannot empty")
		return
	}

	if userID, err = strconv.ParseInt(userIDparam, 10, 64); err != nil {
		log.Println("UserID must be a number")
		return
	}

	data, err = services.SampleGetWithRouteParam(userID)

}
