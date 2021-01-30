package services

import "errors"

func SampleGetWithRouteAndQueryParam(UserID int64) (interface{}, error) {

	var (
		err  error
		resp SampleGetResponse
	)

	switch UserID {
	case 100:

		resp.UserID = 100
		resp.Name = "John Doe"
		resp.Age = "16 Years Old"
		resp.PhoneNumber = "12345"

	case 101:

		resp.UserID = 101
		resp.Name = "Maria"
		resp.Age = "17 Years Old"
		resp.PhoneNumber = "54321"

	default:
		err = errors.New("UserID not found")
	}

	return resp, err
}
