package router

import (
	"github.com/zayarhtet/seap-api/src/server/controller"
)

func publicRoutes() {
	public := seapRouter.Group("/")
	public.GET("/", controller.Welcome())

	authed := seapRouter.Group("/api/auth")
	authed.POST("/register", controller.Register())
	authed.POST("/login", controller.Login())
}

func adminRoutes() {
	admin := seapRouter.Group("/api/admin/")
	admin.Use(controller.AdminMiddleware())
	admin.GET("/roles", controller.GetAllRoles())
	admin.GET("/members", controller.GetAllMembers())
	admin.GET("/role/:id", controller.GetRoleById())
}

func individualRoutes() {
	protected := seapRouter.Group("/api/my/")
	protected.Use(controller.IndividualMiddleware())
	protected.GET("/role", controller.Welcome())
	protected.GET("/member", controller.Welcome())
}
