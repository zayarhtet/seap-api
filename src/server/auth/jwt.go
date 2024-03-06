package auth

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"github.com/zayarhtet/seap-api/src/server/model/dto"
)

func GenerateToken(username, role string) dto.Response {
	// generate token
	return dto.NewDataResponse(username)
}

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWT(id string) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}

func ValidateJWT(tokenString string) error {
	token, err := getToken(tokenString)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}
func getToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

//func CurrentUser(context *gin.Context) (model.User, error) {
//	err := ValidateJWT(context)
//	if err != nil {
//		return model.User{}, err
//	}
//	token, _ := getToken(context)
//	claims, _ := token.Claims.(jwt.MapClaims)
//	userId := uint(claims["id"].(float64))
//
//	user, err := model.FindUserById(userId)
//	if err != nil {
//		return model.User{}, err
//	}
//	return user, nil
//}

func getTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
