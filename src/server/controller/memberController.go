package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zayarhtet/seap-api/src/server/service"
)

type MemberController interface {
	getAllMembers(*gin.Context)
	getAllMembersWithFamilies(*gin.Context)
	deleteMember(*gin.Context)
	getMemberById(*gin.Context)
}

type memberControllerImpl struct {
	ms service.MemberService
}

var memberControllerObj MemberController

func initMember() {
	if memberControllerObj != nil {
		return
	}
	memberControllerObj = &memberControllerImpl{ms: service.NewMemberService()}
}

func (mc *memberControllerImpl) getAllMembers(context *gin.Context) {
	size, page := paginated(context)
	response, err := mc.ms.GetAllMembersResponse(size, page)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, response)
}

func (mc *memberControllerImpl) getAllMembersWithFamilies(context *gin.Context) {
	size, page := paginated(context)
	response, err := mc.ms.GetAllMembersWithFamiliesResponse(size, page)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, response)
}

func (mc *memberControllerImpl) getMemberById(context *gin.Context) {
	idRaw := context.Param("id")
	resp, err := mc.ms.GetMemberByIdResponse(idRaw)
	if err != nil {
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, err.Error())))
		return
	}
	context.JSON(http.StatusOK, resp)
}

func (mc *memberControllerImpl) deleteMember(context *gin.Context) {
	idRaw := context.Param("id")
	resp, err := mc.ms.DeleteMemberResponse(idRaw)
	if err != nil {
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, err.Error())))
		return
	}
	context.JSON(http.StatusOK, resp)
}

func GetAllMembers() func(*gin.Context) {
	return memberControllerObj.getAllMembers
}

func GetAllMembersWithFamilies() func(*gin.Context) {
	return memberControllerObj.getAllMembersWithFamilies
}

func GetMemberById() func(*gin.Context) {
	return memberControllerObj.getMemberById
}

func DeleteMember() func(*gin.Context) {
	return memberControllerObj.deleteMember
}
