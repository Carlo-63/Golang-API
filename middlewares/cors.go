package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetCorsPolicy(allowedOrigins string) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
		context.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		context.Next()
	}
}

func EarlyExitOnPreflightRequests() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Request.Method == http.MethodOptions {
			context.Status(http.StatusOK)
			return
		}
		context.Next()
	}
}
