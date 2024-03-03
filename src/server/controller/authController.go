package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"github.com/zayarhtet/seap-api/src/server/service"
	"net/http"
)

type AuthController interface {
	registerResp(*gin.Context)
}

type authControllerImpl struct {
	ms service.MemberService
}

var authControllerObj AuthController

func initAuth() {
	if authControllerObj != nil {
		return
	}

	authControllerObj = &authControllerImpl{ms: service.NewMemberService()}
}

func (a *authControllerImpl) registerResp(context *gin.Context) {
	var input dto.SignUpRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, "Invalid Input")))
		return
	}
	savedUser, err := a.ms.SignUp(input)

	if err != nil {
		context.JSON(http.StatusConflict, savedUser)
		return
	}

	context.JSON(http.StatusCreated, savedUser)
}

func Register() func(*gin.Context) {
	return authControllerObj.registerResp
}
