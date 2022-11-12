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
	Count   int64       `json:"count"`
}

type ResponseUser struct {
	Success bool                   `json:"status"`
	Message string                 `json:"message"`
	Errors  interface{}            `json:"errors"`
	Data    interface{}            `json:"data"`
	Image   map[string]interface{} `json:"image"`
}

type ImageResponse struct {
	StatusCode int                    `json:"statusCode"`
	Message    string                 `json:"message"`
	Data       map[string]interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildImageResponse(status int, message string, data map[string]interface{}) ImageResponse {
	res := ImageResponse{
		StatusCode: status,
		Message:    message,
		Data:       data,
	}
	return res
}

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

func BuildSuccessAddCart(success bool, message string, data interface{}, total int, count int64) ResponseCart {
	res := ResponseCart{
		Success: success,
		Message: message,
		Errors:  nil,
		Data:    data,
		Total:   total,
		Count:   count,
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
