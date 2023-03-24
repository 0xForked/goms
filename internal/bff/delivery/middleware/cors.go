package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var allowHeaders = []string{
	"Content-Type",
	"Content-Length",
	"Accept-Encoding",
	"Authorization",
	"accept",
	"origin",
	"Cache-Control",
	"X-Requested-With",
}

var allowMethods = []string{
	http.MethodGet,
	http.MethodPut,
	http.MethodPatch,
	http.MethodPost,
	http.MethodDelete,
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", strings.Join(allowHeaders, ","))
		c.Writer.Header().Set("Access-Control-Allow-Methods", strings.Join(allowMethods, ","))

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
