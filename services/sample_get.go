package services

func SampleGet() (interface{}, error) {
	var (
		resp []SampleGetResponse
	)

	resp = []SampleGetResponse{
		SampleGetResponse{
			Name:        "John Doe",
			Age:         "16 Years Old",
			PhoneNumber: "12345",
		},
		SampleGetResponse{
			Name:        "Maria",
			Age:         "17 Years Old",
			PhoneNumber: "54321",
		},
	}

	return resp, nil
}
