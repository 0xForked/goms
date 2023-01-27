package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SuccessRespond struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type ErrorRespond struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   string `json:"data"`
}

type ValidationError struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message any    `json:"message"`
}

func WrapHTTPMessage(context *gin.Context, code int, data interface{}) {
	switch code {
	case http.StatusOK, http.StatusCreated:
		context.JSON(code, SuccessRespond{
			Code:   code,
			Status: http.StatusText(code),
			Data:   data,
		})
		return
	case http.StatusUnprocessableEntity:
		context.JSON(code, ValidationError{
			Code:    code,
			Status:  http.StatusText(code),
			Message: data,
		})
		return
	default:
		msg := func() string {
			switch {
			case data != nil:
				return data.(string)
			case code == http.StatusBadRequest:
				return "something went wrong with the request"
			default:
				return "something went wrong with the server"
			}
		}()

		context.JSON(code, ErrorRespond{
			Code:   code,
			Status: http.StatusText(code),
			Data:   msg,
		})
	}
}
