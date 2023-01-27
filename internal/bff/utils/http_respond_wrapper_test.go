package utils_test

import (
	"github.com/aasumitro/goms/internal/bff/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHTTPRespondWrapper(t *testing.T) {
	tests := []struct {
		name     string
		code     int
		data     interface{}
		expected interface{}
	}{
		{
			name:     "success with no pagination",
			code:     http.StatusOK,
			data:     []string{"foo", "bar"},
			expected: utils.SuccessRespond{Code: http.StatusOK, Status: "OK", Data: []string{"foo", "bar"}},
		},
		{
			name:     "error with data",
			code:     http.StatusBadRequest,
			data:     "invalid request",
			expected: utils.ErrorRespond{Code: http.StatusBadRequest, Status: "Bad Request", Data: "invalid request"},
		},
		{
			name:     "error with no data",
			code:     http.StatusBadRequest,
			expected: utils.ErrorRespond{Code: http.StatusBadRequest, Status: "Bad Request", Data: "something went wrong with the request"},
		},
		{
			name:     "error with no data and server error code",
			code:     http.StatusInternalServerError,
			expected: utils.ErrorRespond{Code: http.StatusInternalServerError, Status: "Internal Server Error", Data: "something went wrong with the server"},
		},
		{
			name: "error unprocessable entity",
			code: http.StatusUnprocessableEntity,
			data: []string{"foo", "bar"},
			expected: utils.ValidationError{
				Code:    http.StatusUnprocessableEntity,
				Status:  "Unprocessable Entity",
				Message: []string{"foo", "bar"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			writer := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(writer)
			utils.WrapHTTPMessage(c, test.code, test.data)
		})
	}
}
