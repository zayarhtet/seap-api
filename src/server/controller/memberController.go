package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zayarhtet/seap-api/src/server/service"
	"net/http"
)

type MemberController interface {
	getAllMembers(*gin.Context)
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

func GetAllMembers() func(ctx *gin.Context) {
	return memberControllerObj.getAllMembers
}
