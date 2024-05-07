package dao

import (
	"github.com/zayarhtet/seap-api/src/util"
)

type Credential struct {
	CredentialId string `gorm:"primary_key" json:"credentialId"`
	Password     string `json:"password"`
}

func NewCredential(pword string) (*Credential, error) {
	err := util.Encrypt(&pword)
	if err != nil {
		return nil, err
	}
	return &Credential{
		CredentialId: util.NewUUID(),
		Password:     pword,
	}, nil
}

func (Credential) TableName() string {
	return "credential"
}
