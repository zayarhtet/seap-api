package router

import (
	"github.com/zayarhtet/seap-api/src/server/controller"
)

func publicRoutes() {
	public := seapRouter.Group("/api")
	public.GET("/roles", controller.GetAllRoles())
	public.GET("/roles/:id", controller.GetRoleById())

}
