package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/zayarhtet/seap-api/src/server/service"
)

type RoleController interface {
	GetAllRoles(*gin.Context)
	GetRoleById(*gin.Context)
}

type roleControllerImpl struct {
	rs service.RoleService
}

var roleControllerObj RoleController

func initRole() {
	if roleControllerObj != nil {
		return
	}
	rs := service.NewRoleService()
	roleControllerObj = NewRoleController(rs)
}

func (rc *roleControllerImpl) SetRoleService(rs service.RoleService) {
	rc.rs = rs
}

func NewRoleController(ms service.RoleService) RoleController {
	return &roleControllerImpl{rs: ms}
}

func (rc *roleControllerImpl) GetAllRoles(context *gin.Context) {
	getPaginatedResponseByCallBack(context, rc.rs.GetAllRolesResponse)
}

func (rc *roleControllerImpl) GetRoleById(context *gin.Context) {
	idRaw, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := uint(idRaw)
	response, err := rc.rs.GetRoleByIdResponse(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, response)
}

func GetAllRoles() gin.HandlerFunc {
	return roleControllerObj.GetAllRoles
}

func GetRoleById() gin.HandlerFunc {
	return roleControllerObj.GetRoleById
}
