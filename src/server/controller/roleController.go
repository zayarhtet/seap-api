package controller

import (
	"fmt"
	"net/http"

	"github.com/zayarhtet/seap-api/src/server/service"
	"github.com/gin-gonic/gin"
)

type RoleController interface {
	getAllRoles(*gin.Context)
}

type roleControllerImpl struct {
	rs service.RoleService
}

var roleControllerObj RoleController

func initRole() {
	if roleControllerObj != nil { return }

	roleControllerObj = &roleControllerImpl{rs: service.NewRoleService()}
}

func (rc *roleControllerImpl) getAllRoles(context *gin.Context) {
    roles, err := rc.rs.GetAllRolesResponse(context)

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	fmt.Println(roles)

    context.JSON(http.StatusOK, gin.H{"data": *roles})
}

func GetAllRoles() func(*gin.Context) {
	return roleControllerObj.getAllRoles
}

