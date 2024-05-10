package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zayarhtet/seap-api/src/server/service"
)

type IndividualController interface {
	GetMyMember(*gin.Context)
	GetMyRole(*gin.Context)
	GetMyFamilies(*gin.Context)
	GetMyDuties(*gin.Context)
	GetMyRoleInFamily(*gin.Context)
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
	ms := service.NewMemberService()
	rs := service.NewRoleService()
	ds := service.NewDutyService()
	fs := service.NewFamilyService()
	individualControllerObj = NewIndividualController(ms, rs, fs, ds)
}

func (ic *individualControllerImpl) SetIndividualServices(ms service.MemberService, rs service.RoleService, fs service.FamilyService, ds service.DutyService) {
	ic.ms = ms
	ic.fs = fs
	ic.rs = rs
	ic.ds = ds
}

func NewIndividualController(ms service.MemberService, rs service.RoleService, fs service.FamilyService, ds service.DutyService) IndividualController {
	return &individualControllerImpl{ms, rs, fs, ds}
}

func (ic *individualControllerImpl) GetMyMember(context *gin.Context) {
	idRaw := context.MustGet("username").(string)
	getOneResponseByCallBack(context, idRaw, ic.ms.GetMemberByIdResponse)
}

func (ic *individualControllerImpl) GetMyRole(context *gin.Context) {
	idRaw := context.MustGet("username").(string)
	getOneResponseByCallBack(context, idRaw, ic.rs.GetRoleByMemberResponse)
}

func (ic *individualControllerImpl) GetMyFamilies(context *gin.Context) {
	idRaw := context.MustGet("username").(string)
	getOneResponseByCallBack(context, idRaw, ic.fs.GetFamiliesOnlyByUsername)
}

func (ic *individualControllerImpl) GetMyDuties(context *gin.Context) {
	idRaw := context.MustGet("username").(string)
	getOneResponseByCallBack(context, idRaw, ic.ds.GetAllDutiesByMemberResponse)
}

func (ic *individualControllerImpl) GetMyRoleInFamily(context *gin.Context) {
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
	return individualControllerObj.GetMyMember
}

func GetMyFamilies() gin.HandlerFunc {
	return individualControllerObj.GetMyFamilies
}

func GetMyRole() gin.HandlerFunc {
	return individualControllerObj.GetMyRole
}

func GetMyRoleInFamily() gin.HandlerFunc {
	return individualControllerObj.GetMyRoleInFamily
}

func GetMyDuties() gin.HandlerFunc {
	return individualControllerObj.GetMyDuties
}
