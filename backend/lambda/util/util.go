package util

import (
	"github.com/seanrmurphy/ws-echo/backend/lambda/types"
)

type GenericReturnMessage struct {
	Message string
}

func CreateResponse(status string, msg string, data string) types.Response {
	return types.Response{
		Status:  status,
		Message: msg,
		Data:    data,
	}
}
