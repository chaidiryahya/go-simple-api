package services

type SampleGetResponse struct {
	UserID      int64  `json:"userID"`
	Name        string `json:"name"`
	Age         string `json:"age"`
	PhoneNumber string `json:"phoneNumber"`
}

type SamplePostWithJsonBodyParam struct {
	YourName string `json:"yourName"`
	UserID   int64  `json:"userID"`
}
