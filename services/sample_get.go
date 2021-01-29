package services

func SampleGet() (interface{}, error) {
	var (
		resp []SampleGetResponse
	)

	resp = []SampleGetResponse{
		SampleGetResponse{
			UserID:      100,
			Name:        "John Doe",
			Age:         "16 Years Old",
			PhoneNumber: "12345",
		},
		SampleGetResponse{
			UserID:      101,
			Name:        "Maria",
			Age:         "17 Years Old",
			PhoneNumber: "54321",
		},
	}

	return resp, nil
}
