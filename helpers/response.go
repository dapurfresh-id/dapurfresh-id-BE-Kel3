package helpers

import "strings"

type Response struct {
	Success bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type ResponseCart struct {
	Success bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
	Total   int         `json:"total"`
}

type EmptyObj struct{}

func BuildResponse(success bool, message string, data interface{}) Response {
	res := Response{
		Success: success,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Success: false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}

func BuildSuccessAddCart(success bool, message string, data interface{}, total int) ResponseCart {
	res := ResponseCart{
		Success: success,
		Message: message,
		Errors:  nil,
		Data:    data,
		Total:   total,
	}
	return res
}
