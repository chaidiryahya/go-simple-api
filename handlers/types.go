package handlers

type APIResponse struct {
	Status            string      `json:"status"`
	StatusCode        int         `json:"statusCode"`
	ServerProcessTime string      `json:"serverProcessTime"`
	Data              interface{} `json:"data,omitempty"`
	MessageError      string      `json:"messageError,omitempty"`
}
