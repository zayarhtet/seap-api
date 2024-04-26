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
	familyTutorMiddleware(*gin.Context)
	familyTuteeMiddleware(*gin.Context)
	familyMemberMiddleware(*gin.Context)
	corsMiddleware(*gin.Context)
}

type authControllerImpl struct {
	ms service.MemberService
	fs service.FamilyService
}

var authControllerObj AuthController

func initAuth() {
	if authControllerObj != nil {
		return
	}

	authControllerObj = &authControllerImpl{ms: service.NewMemberService(), fs: service.NewFamilyService()}
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

func authorizeByFamilyRole(context *gin.Context, roles []string, a *authControllerImpl) {
	username, role, err := validateTokenAndClaims(context)
	if err != nil {
		context.JSON(parseErrorAndResponse(err))
		context.Abort()
		return
	}
	famIdRaw := context.Param("famId")

	familyRole, err := a.fs.RoleInFamily(username, famIdRaw)

	if err != nil {
		context.JSON(http.StatusUnauthorized, service.BeforeErrorResponse(service.PrepareErrorMap(401, "You are not a tutor of this family.")))
		context.Abort()
		return
	}
	for _, s := range roles {
		if s == familyRole {
			context.Set("username", username)
			context.Set("familyRole", familyRole)
			context.Set("role", role)
			context.Next()
			return
		}
	}
	context.JSON(http.StatusUnauthorized, service.BeforeErrorResponse(service.PrepareErrorMap(401, "You are not a tutor of this family.")))
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

func (a *authControllerImpl) familyTutorMiddleware(context *gin.Context) {
	authorizeByFamilyRole(context, []string{"tutor"}, a)
}

func (a *authControllerImpl) familyTuteeMiddleware(context *gin.Context) {
	authorizeByFamilyRole(context, []string{"tutee"}, a)
}

func (a *authControllerImpl) familyMemberMiddleware(context *gin.Context) {
	authorizeByFamilyRole(context, []string{"tutee", "tutor"}, a)
}

func (a *authControllerImpl) corsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

func Register() gin.HandlerFunc {
	return authControllerObj.registerResp
}

func Login() gin.HandlerFunc {
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

func FamilyTutorMiddleware() gin.HandlerFunc {
	return authControllerObj.familyTutorMiddleware
}

func FamilyTuteeMiddleware() gin.HandlerFunc {
	return authControllerObj.familyTuteeMiddleware
}

func FamilyMemberMiddleware() gin.HandlerFunc {
	return authControllerObj.familyMemberMiddleware
}

func CorsMiddleware() gin.HandlerFunc {
	return authControllerObj.corsMiddleware
}
