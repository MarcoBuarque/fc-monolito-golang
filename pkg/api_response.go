package pkg

import "github.com/MarcoBuarque/monolito/constant"

type ApiResponse[T any] struct {
	ResponseKey     string `json:"response_key"`
	ResponseMessage string `json:"response_message"`
	Data            T      `json:"data"`
}

type Null struct{}

func BuildCustomResponse(status, message string) ApiResponse[any] {
	return ApiResponse[any]{
		ResponseKey:     status,
		ResponseMessage: message,
		Data:            Null{},
	}
}

func BuildResponse[T any](responseStatus constant.ResponseStatus, data T) ApiResponse[T] {
	return BuildResponse_(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data)
}

func BuildResponse_[T any](status string, message string, data T) ApiResponse[T] {
	return ApiResponse[T]{
		ResponseKey:     status,
		ResponseMessage: message,
		Data:            data,
	}
}
