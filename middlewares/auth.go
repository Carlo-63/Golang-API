package middlewares

import (
	"main/models"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		headerToken := context.GetHeader("Authorization")

		if headerToken == "" {
			context.JSON(http.StatusUnauthorized, "Token mancante")
			context.Abort()
			return
		}

		secretKey := context.MustGet("ConfigKey").(models.Config).SecretKey

		tokenMap, err := utils.ValidateToken(headerToken, []byte(secretKey))

		if err != nil {
			context.JSON(http.StatusUnauthorized, "Token mancante")
			context.Abort()
			return
		}

		context.Set("Username", (*tokenMap)["sub"])

		context.Next()
	}
}
