package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zayarhtet/seap-api/src/server/service"
)

type MemberController interface {
	GetAllMembers(*gin.Context)
	GetAllMembersWithFamilies(*gin.Context)
	DeleteMember(*gin.Context)
	GetMemberById(*gin.Context)
	GrantTutorRole(*gin.Context)
	RevokeTutorRole(*gin.Context)
}

type memberControllerImpl struct {
	ms service.MemberService
}

var memberControllerObj MemberController

func initMember() {
	if memberControllerObj != nil {
		return
	}
	ms := service.NewMemberService()
	memberControllerObj = NewMemberController(ms)
}

func (mc *memberControllerImpl) SetMemberService(ms service.MemberService) {
	mc.ms = ms
}

func NewMemberController(ms service.MemberService) MemberController {
	return &memberControllerImpl{ms: ms}
}

func (mc *memberControllerImpl) GetAllMembers(context *gin.Context) {
	getPaginatedResponseByCallBack(context, mc.ms.GetAllMembersResponse)
}

func (mc *memberControllerImpl) GetAllMembersWithFamilies(context *gin.Context) {
	getPaginatedResponseByCallBack(context, mc.ms.GetAllMembersWithFamiliesResponse)
}

func (mc *memberControllerImpl) GetMemberById(context *gin.Context) {
	idRaw := context.Param("id")
	getOneResponseByCallBack(context, idRaw, mc.ms.GetMemberByIdResponse)
}

func (mc *memberControllerImpl) DeleteMember(context *gin.Context) {
	idRaw := context.Param("id")
	getOneResponseByCallBack(context, idRaw, mc.ms.DeleteMemberResponse)
}

func (mc *memberControllerImpl) GrantTutorRole(context *gin.Context) {
	username := context.Param("username")

	resp, err := mc.ms.GrantRoleResponse(username, 1)
	if err != nil {
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, err.Error())))
		return
	}
	context.JSON(http.StatusOK, resp)
}
func (mc *memberControllerImpl) RevokeTutorRole(context *gin.Context) {
	username := context.Param("username")
	resp, err := mc.ms.GrantRoleResponse(username, 2)
	if err != nil {
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, err.Error())))
		return
	}
	context.JSON(http.StatusOK, resp)
}

func GetAllMembers() gin.HandlerFunc {
	return memberControllerObj.GetAllMembers
}

func GetAllMembersWithFamilies() gin.HandlerFunc {
	return memberControllerObj.GetAllMembersWithFamilies
}

func GetMemberById() gin.HandlerFunc {
	return memberControllerObj.GetMemberById
}

func DeleteMember() gin.HandlerFunc {
	return memberControllerObj.DeleteMember
}

func PromoteRole() gin.HandlerFunc {
	return memberControllerObj.GrantTutorRole
}

func DemoteRole() gin.HandlerFunc {
	return memberControllerObj.RevokeTutorRole
}
