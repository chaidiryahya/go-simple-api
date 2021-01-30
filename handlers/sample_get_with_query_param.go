package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chaidiryahya/go-simple-api/services"
	"github.com/chaidiryahya/go-simple-api/utils"
	"github.com/julienschmidt/httprouter"
)

func SampleGetWithQueryParam(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		err      error
		data     interface{}
		response APIResponse
		start    = time.Now()
		UserID   int64
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

	queryParam := r.URL.Query()

	userID := queryParam.Get("user_id")

	if UserID, err = strconv.ParseInt(userID, 10, 64); err != nil {
		log.Println("UserID must be a number")
		return
	}

	data, err = services.SampleGetWithRouteAndQueryParam(UserID)

}
