package model

type response[T any] struct {
	Message string `json:"message"`
	Status  uint   `json:"status"`
	Data    T      `json:"data"`
}

func ResponseBuilder(message string, status uint, data any) response[any] {

	res := response[any]{Message: message, Status: status, Data: data}

	return res

}
