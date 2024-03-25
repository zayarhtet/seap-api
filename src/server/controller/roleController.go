package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/zayarhtet/seap-api/src/server/service"
)

type RoleController interface {
	getAllRoles(*gin.Context)
	getRoleById(*gin.Context)
}

type roleControllerImpl struct {
	rs service.RoleService
}

var roleControllerObj RoleController

func initRole() {
	if roleControllerObj != nil {
		return
	}

	roleControllerObj = &roleControllerImpl{rs: service.NewRoleService()}
}

func (rc *roleControllerImpl) getAllRoles(context *gin.Context) {
	getPaginatedResponseByCallBack(context, rc.rs.GetAllRolesResponse)
}

func (rc *roleControllerImpl) getRoleById(context *gin.Context) {
	idRaw, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := uint(idRaw)
	response, err := rc.rs.GetRoleByIdResponse(id)

	context.JSON(http.StatusOK, response)
}

func GetAllRoles() gin.HandlerFunc {
	return roleControllerObj.getAllRoles
}

func GetRoleById() gin.HandlerFunc {
	return roleControllerObj.getRoleById
}
