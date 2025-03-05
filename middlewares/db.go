package middlewares

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DBMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("DbKey", db)
		context.Next()
	}
}
