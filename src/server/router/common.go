package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var seapRouter *gin.Engine

func Init() {
	if seapRouter == nil {
		seapRouter = gin.Default()
	}
	seapRouter.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	publicRoutes()
	adminRoutes()
	individualRoutes()
	tutorRoutes()
	familyTutorRoutes()
	familyTuteeRoutes()
	familyMemberRoutes()
	seapRouter.Run(":8000")
}
