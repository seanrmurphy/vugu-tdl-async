package util

import (
	"github.com/seanrmurphy/vugu-tdl-async/backend/lambda/types"
)

type GenericReturnMessage struct {
	Message string
}

func CreateResponse(t string, status string, msg string, data string) types.Response {
	return types.Response{
		Status:  status,
		Type:    t,
		Message: msg,
		Data:    data,
	}
}
