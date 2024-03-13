package controller

import (
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
	getPaginatedResponseByCallBack(context, mc.ms.GetAllMembersResponse)
}

func (mc *memberControllerImpl) getAllMembersWithFamilies(context *gin.Context) {
	getPaginatedResponseByCallBack(context, mc.ms.GetAllMembersWithFamiliesResponse)
}

func (mc *memberControllerImpl) getMemberById(context *gin.Context) {
	idRaw := context.Param("id")
	getOneResponseByCallBack(context, idRaw, mc.ms.GetMemberByIdResponse)
}

func (mc *memberControllerImpl) deleteMember(context *gin.Context) {
	idRaw := context.Param("id")
	getOneResponseByCallBack(context, idRaw, mc.ms.DeleteMemberResponse)

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
