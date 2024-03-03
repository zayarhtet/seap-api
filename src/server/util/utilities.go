package util

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func CurrentTimeString() string {
	return time.Now().Format(time.RFC3339)
}

func Encrypt(password *string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	*password = string(passwordHash)
	return nil
}

func NewUUID() string {
	return uuid.New().String()
}
