package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zayarhtet/seap-api/src/server/service"
)

type IndividualController interface {
	getMyMember(*gin.Context)
	getMyRole(*gin.Context)
}

type individualControllerImpl struct {
	ms service.MemberService
	rs service.RoleService
}

var individualControllerObj IndividualController

func initIndividual() {
	if individualControllerObj != nil {
		return
	}
	individualControllerObj = &individualControllerImpl{
		ms: service.NewMemberService(),
		rs: service.NewRoleService(),
	}
}

func (ic *individualControllerImpl) getMyMember(context *gin.Context) {
	idRaw := context.MustGet("username").(string)
	resp, err := ic.ms.GetMemberByIdResponse(idRaw)
	if err != nil {
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, err.Error())))
		return
	}
	context.JSON(http.StatusOK, resp)
}

func (ic *individualControllerImpl) getMyRole(context *gin.Context) {
	idRaw := context.MustGet("username").(string)
	resp, err := ic.rs.GetRoleByMemberResponse(idRaw)
	if err != nil {
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, err.Error())))
		return
	}
	context.JSON(http.StatusOK, resp)
}

func GetMyMember() func(*gin.Context) {
	return individualControllerObj.getMyMember
}

func GetMyRole() func(*gin.Context) {
	return individualControllerObj.getMyRole
}
