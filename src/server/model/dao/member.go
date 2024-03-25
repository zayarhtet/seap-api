package dao

import (
	"time"

	"github.com/zayarhtet/seap-api/src/server/model/dto"
)

type Member struct {
	FirstName    string      `gorm:"column:first_name" json:"firstName"`
	LastName     string      `json:"lastName"`
	Username     string      `gorm:"primary_key" json:"username"`
	Email        string      `json:"email"`
	CredentialId string      `json:"-"`
	RoleId       uint        `json:"roleId"`
	Role         dto.RoleDto `gorm:"references:RoleId" json:"role"`
	CreatedAt    time.Time   `json:"createdAt"`
	ModifiedAt   time.Time   `json:"modifiedAt"`
	//Role         Role `gorm:"ForeignKey:RoleId;references:RoleId"`
}

func (Member) TableName() string {
	return "member"
}

type MemberWithFamilies struct {
	FirstName string            `gorm:"column:first_name" json:"firstName"`
	LastName  string            `json:"lastName"`
	Username  string            `gorm:"primary_key" json:"username"`
	Email     string            `json:"email"`
	Families  []FamilyForMember `gorm:"foreignKey:Username" json:"families"`
}

func (MemberWithFamilies) TableName() string {
	return "member"
}

type MemberChild struct {
	FirstName string `gorm:"column:first_name" json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `gorm:"primary_key" json:"username"`
	Email     string `json:"email"`
}

func (MemberChild) TableName() string {
	return "member"
}

type MemberWithDuties struct {
	FirstName string            `gorm:"column:first_name" json:"firstName"`
	LastName  string            `json:"lastName"`
	Username  string            `gorm:"primary_key" json:"username"`
	Email     string            `json:"email"`
	Duties    []DutiesForMember `gorm:"foreignKey:Username" json:"duties"`
}

func (MemberWithDuties) TableName() string {
	return "member"
}
