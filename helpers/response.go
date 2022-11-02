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

type ResponseUser struct {
	Success bool                   `json:"status"`
	Message string                 `json:"message"`
	Errors  interface{}            `json:"errors"`
	Data    interface{}            `json:"data"`
	Image   map[string]interface{} `json:"image"`
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

func BuildSuccessUpdate(success bool, message string, data interface{}, image map[string]interface{}) ResponseUser {
	res := ResponseUser{
		Success: success,
		Message: message,
		Errors:  nil,
		Data:    data,
		Image:   image,
	}
	return res
}
