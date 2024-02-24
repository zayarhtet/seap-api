package dao

import (
	"golang.org/x/crypto/bcrypt"
)

type Credential struct {
	credentialId string
	password     string
}

func (Credential) TableName() string {
	return "credential"
}

func Encrypt(password *string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    *password = string(passwordHash)
	return nil
}
