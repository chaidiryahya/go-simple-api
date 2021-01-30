package handlers

import (
	"net/http"
	"time"

	"github.com/chaidiryahya/go-simple-api/utils"
	"github.com/julienschmidt/httprouter"
)

func Home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		response APIResponse
		start    = time.Now()
	)

	response.ServerProcessTime = time.Since(start).String()
	response.Status = http.StatusText(http.StatusOK)
	response.Data = "Hello :)"

	utils.WriteJsonResponse(w, response, http.StatusOK)
}
