package services

func SampleGet() (interface{}, error) {
	var resp SampleGetResponse

	resp.Name = "John Doe"
	resp.Age = "16 Years Old"
	resp.PhoneNumber = "123456"

	return resp, nil
}
