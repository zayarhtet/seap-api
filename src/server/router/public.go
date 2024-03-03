package router

import (
	"github.com/zayarhtet/seap-api/src/server/auth"
	"github.com/zayarhtet/seap-api/src/server/controller"
)

func publicRoutes() {
	public := seapRouter.Group("/")
	public.GET("/", controller.Welcome())
	auth := seapRouter.Group("/api/auth")
	auth.POST("/register", controller.Register())
}

func protectedRoutes() {
	protected := seapRouter.Group("/api")
	protected.Use(auth.JwtAuthMiddleware())
	protected.GET("/roles", controller.GetAllRoles())
	protected.GET("/role/:id", controller.GetRoleById())
	protected.GET("members", controller.GetAllMembers())
}
