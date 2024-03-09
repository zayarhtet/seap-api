package dao

import (
	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"time"
)

type Member struct {
	FirstName    string      `gorm:"column:first_name" json:"firstName"`
	LastName     string      `json:"lastName"`
	Username     string      `gorm:"primary_key" json:"username"`
	Email        string      `json:"email"`
	CredentialId string      `json:"credentialId"`
	RoleId       uint        `json:"roleId"`
	Role         dto.RoleDto `gorm:"references:RoleId" json:"role"`
	CreatedAt    time.Time   `json:"createdAt"`
	ModifiedAt   time.Time   `json:"modifiedAt"`
	//Role         Role `gorm:"ForeignKey:RoleId;references:RoleId"`
}

func (Member) TableName() string {
	return "member"
}
