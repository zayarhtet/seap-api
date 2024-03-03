package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/zayarhtet/seap-api/src/server/service"
	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if false {
			context.JSON(http.StatusUnauthorized, service.BeforeErrorResponse(service.PrepareErrorMap(401, "Unauthorized access")))
			context.Abort()
			return
		}
		context.Next()
	}
}
