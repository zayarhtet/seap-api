package router

import (
	"github.com/gin-gonic/gin"
)

var seapRouter *gin.Engine

func Init() {
	if seapRouter == nil {
		seapRouter = gin.Default()
	}
	publicRoutes()
	seapRouter.Run(":8000")
}
