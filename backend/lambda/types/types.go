package types

type Message struct {
	Type string `json:"messagetype"`
	Data string `json:"data"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Type    string `json:"messagetype"`
	Data    string `json:"data"`
}
