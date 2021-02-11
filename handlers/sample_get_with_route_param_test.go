package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chaidiryahya/go-simple-api/services"
	"github.com/julienschmidt/httprouter"
)

func TestSampleGetWithRouteParam(t *testing.T) {
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		ps httprouter.Params
	}

	/* const for scenario name */
	const (
		Success          = "Success"
		ErrorParseUserID = "ErrorParseUserID"
		ErrorEmptyUserID = "ErrorEmptyUserID"
	)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/get_with_route_param/:user_id", nil)

	tests := []struct {
		name string
		args args
	}{
		{
			Success, args{
				w:  rr,
				r:  req,
				ps: []httprouter.Param{{Key: "user_id", Value: "1"}},
			},
		},
		{
			ErrorParseUserID, args{
				w:  rr,
				r:  req,
				ps: []httprouter.Param{{Key: "user_id", Value: "X"}},
			},
		},
		{
			ErrorEmptyUserID, args{
				w:  rr,
				r:  req,
				ps: []httprouter.Param{},
			},
		},
	}
	for _, tt := range tests {

		var SampleGetWithRouteAndQueryParamOriginal = services.SampleGetWithRouteAndQueryParam

		defer func() {
			services.SampleGetWithRouteAndQueryParam = SampleGetWithRouteAndQueryParamOriginal
		}()

		services.SampleGetWithRouteAndQueryParam = func(userID int64) (interface{}, error) {
			return nil, nil
		}

		t.Run(tt.name, func(t *testing.T) {
			SampleGetWithRouteParam(tt.args.w, tt.args.r, tt.args.ps)
		})
	}
}
