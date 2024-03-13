package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/zayarhtet/seap-api/src/server/service"
)

type IndividualController interface {
	getMyMember(*gin.Context)
	getMyRole(*gin.Context)
	getMyFamilies(*gin.Context)
}

type individualControllerImpl struct {
	ms service.MemberService
	rs service.RoleService
	fs service.FamilyService
}

var individualControllerObj IndividualController

func initIndividual() {
	if individualControllerObj != nil {
		return
	}
	individualControllerObj = &individualControllerImpl{
		ms: service.NewMemberService(),
		rs: service.NewRoleService(),
		fs: service.NewFamilyService(),
	}
}

func (ic *individualControllerImpl) getMyMember(context *gin.Context) {
	idRaw := context.MustGet("username").(string)
	getOneResponseByCallBack(context, idRaw, ic.ms.GetMemberByIdResponse)
}

func (ic *individualControllerImpl) getMyRole(context *gin.Context) {
	idRaw := context.MustGet("username").(string)
	getOneResponseByCallBack(context, idRaw, ic.rs.GetRoleByMemberResponse)
}

func (ic *individualControllerImpl) getMyFamilies(context *gin.Context) {
	idRaw := context.MustGet("username").(string)
	getOneResponseByCallBack(context, idRaw, ic.fs.GetFamiliesOnlyByUsername)
}

func GetMyMember() func(*gin.Context) {
	return individualControllerObj.getMyMember
}

func GetMyFamilies() func(*gin.Context) {
	return individualControllerObj.getMyFamilies
}

func GetMyRole() func(*gin.Context) {
	return individualControllerObj.getMyRole
}
