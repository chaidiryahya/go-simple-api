package handlers

type APIResponse struct {
	Status            string      `json:"status"`
	ServerProcessTime string      `json:"serverProcessTime"`
	Data              interface{} `json:"data,omitempty"`
	MessageError      string      `json:"messageError,omitempty"`
}
