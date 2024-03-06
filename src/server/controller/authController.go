package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"github.com/zayarhtet/seap-api/src/server/service"
)

type AuthController interface {
	registerResp(*gin.Context)
	loginResp(*gin.Context)
	adminMiddleware(*gin.Context)
	individualMiddleware(*gin.Context)
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

func (a *authControllerImpl) loginResp(context *gin.Context) {
	var input dto.LoginRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, "Invalid Input")))
		return
	}
	loggedUser, err := a.ms.Login(input)

	if err != nil {
		context.JSON(http.StatusNotFound, loggedUser)
		return
	}

	context.JSON(http.StatusOK, loggedUser)
}

func (a *authControllerImpl) adminMiddleware(context *gin.Context) {
	// validate token
	context.Set("username", "miyuki")
	if false {
		context.JSON(http.StatusUnauthorized, service.BeforeErrorResponse(service.PrepareErrorMap(401, "Unauthorized access")))
		context.Abort()
		return
	}
	context.Next()
}

func (a *authControllerImpl) individualMiddleware(context *gin.Context) {
	// validate token
	context.Set("username", "admin")
	if false {
		context.JSON(http.StatusUnauthorized, service.BeforeErrorResponse(service.PrepareErrorMap(401, "Unauthorized access")))
		context.Abort()
		return
	}
	context.Next()
}

func Register() func(*gin.Context) {
	return authControllerObj.registerResp
}

func Login() func(*gin.Context) {
	return authControllerObj.loginResp
}

func AdminMiddleware() gin.HandlerFunc {
	return authControllerObj.adminMiddleware
}

func IndividualMiddleware() gin.HandlerFunc {
	return authControllerObj.individualMiddleware
}
