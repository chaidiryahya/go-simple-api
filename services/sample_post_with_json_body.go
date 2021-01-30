package services

import "fmt"

func SamplePostWithJsonBody(param SamplePostWithJsonBodyParam) (interface{}, error) {

	resp := fmt.Sprintf("Hi %s, userID for your parameters : %d", param.YourName, param.UserID)

	return resp, nil
}
