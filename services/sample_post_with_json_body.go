package services

import "fmt"

var SamplePostWithJsonBody = samplePostWithJsonBody

func samplePostWithJsonBody(param SamplePostWithJsonBodyParam) (interface{}, error) {

	resp := fmt.Sprintf("Hi %s, userID for your parameters : %d", param.YourName, param.UserID)

	return resp, nil
}
