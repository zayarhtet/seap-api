package auth

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateToken(username, role string) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   username,
		"role": role,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	fmt.Println(time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix())
	return token.SignedString(privateKey)
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := getToken(tokenString)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token provided")
	}
}

func getToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	fmt.Println(err)
	return token, err
}
