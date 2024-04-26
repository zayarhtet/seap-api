package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zayarhtet/seap-api/src/server/service"
)

type IndividualController interface {
	getMyMember(*gin.Context)
	getMyRole(*gin.Context)
	getMyFamilies(*gin.Context)
	getMyDuties(*gin.Context)
	getMyRoleInFamily(*gin.Context)
}

type individualControllerImpl struct {
	ms service.MemberService
	rs service.RoleService
	fs service.FamilyService
	ds service.DutyService
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
		ds: service.NewDutyService(),
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

func (ic *individualControllerImpl) getMyDuties(context *gin.Context) {
	idRaw := context.MustGet("username").(string)
	getOneResponseByCallBack(context, idRaw, ic.ds.GetAllDutiesByMemberResponse)
}

func (ic *individualControllerImpl) getMyRoleInFamily(context *gin.Context) {
	idRaw := context.MustGet("username").(string)
	famIdRaw := context.Param("famId")

	resp, err := ic.fs.GetMyRoleInFamily(idRaw, famIdRaw)
	if err != nil {
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, err.Error())))
		return
	}
	context.JSON(http.StatusOK, resp)
}

func GetMyMember() gin.HandlerFunc {
	return individualControllerObj.getMyMember
}

func GetMyFamilies() gin.HandlerFunc {
	return individualControllerObj.getMyFamilies
}

func GetMyRole() gin.HandlerFunc {
	return individualControllerObj.getMyRole
}

func GetMyRoleInFamily() gin.HandlerFunc {
	return individualControllerObj.getMyRoleInFamily
}

func GetMyDuties() gin.HandlerFunc {
	return individualControllerObj.getMyDuties
}
