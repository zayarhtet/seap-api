package controller

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"github.com/zayarhtet/seap-api/src/server/auth"
	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"github.com/zayarhtet/seap-api/src/server/service"
)

type AuthController interface {
	registerResp(*gin.Context)
	loginResp(*gin.Context)
	adminMiddleware(*gin.Context)
	individualMiddleware(*gin.Context)
	tutorMiddleware(*gin.Context)
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

func getTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}

func validateTokenAndClaims(context *gin.Context) (string, string, error) {
	tokenString := getTokenFromRequest(context)
	if len(tokenString) == 0 {
		return "", "", errors.New("no valid authentication token")
	}
	claims, err := auth.ValidateToken(tokenString)
	if err != nil {
		return "", "", err
	}
	return claims["id"].(string), claims["role"].(string), nil
}

func parseErrorAndResponse(err error) (int, dto.Response) {
	if errors.Is(err, jwt.ErrTokenMalformed) {
		return http.StatusUnauthorized, service.PrepareErrorMap(http.StatusUnauthorized, "Invalid token")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return http.StatusUnauthorized, service.PrepareErrorMap(419, "Expired Token")
	} else {
		return http.StatusUnauthorized, service.PrepareErrorMap(http.StatusUnauthorized, "Invalid token")
	}
}

func authorizeByRole(context *gin.Context, roles []string) {
	username, role, err := validateTokenAndClaims(context)
	if err != nil {
		context.JSON(parseErrorAndResponse(err))
		context.Abort()
		return
	}
	for _, s := range roles {
		if s == role {
			context.Set("username", username)
			context.Set("role", role)
			context.Next()
			return
		}
	}
	context.JSON(http.StatusUnauthorized, service.BeforeErrorResponse(service.PrepareErrorMap(401, "Unauthorized access")))
	context.Abort()
}

func (a *authControllerImpl) adminMiddleware(context *gin.Context) {
	authorizeByRole(context, []string{"admin"})
}

func (a *authControllerImpl) individualMiddleware(context *gin.Context) {
	authorizeByRole(context, []string{"tutor", "tutee"})
}

func (a *authControllerImpl) tutorMiddleware(context *gin.Context) {
	authorizeByRole(context, []string{"admin", "tutor"})
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

func TutorMiddleware() gin.HandlerFunc {
	return authControllerObj.tutorMiddleware
}
