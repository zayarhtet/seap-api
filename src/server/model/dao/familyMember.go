package dao

import (
	"time"

	"github.com/zayarhtet/seap-api/src/server/model/dto"
)

type MemberForFamily struct {
	Username   string              `gorm:"column:username;primary_key" json:"-"`
	User       dto.MemberDtoFamily `gorm:"foreignKey:Username;references:Username" json:"member"`
	FamilyId   string              `gorm:"column:family_id;primary_key" json:"-"`
	RoleId     int                 `gorm:"column:role_id;" json:"-"`
	MemberRole dto.RoleDto         `gorm:"foreignKey:RoleId;references:RoleId" json:"familyRole"`
	CreatedAt  time.Time           `json:"addedAt"`
}

func (MemberForFamily) TableName() string {
	return "family_member"
}

type FamilyForMember struct {
	Username   string              `gorm:"column:username;primary_key" json:"-"`
	FamilyId   string              `gorm:"column:family_id;primary_key" json:"-"`
	Family     dto.FamilyDtoMember `gorm:"foreignKey:FamilyId;references:FamilyId" json:"families"`
	RoleId     int                 `gorm:"column:role_id;" json:"-"`
	MemberRole dto.RoleDto         `gorm:"foreignKey:RoleId;references:RoleId" json:"familyRole"`
	CreatedAt  time.Time           `json:"addedAt"`
}

func (FamilyForMember) TableName() string {
	return "family_member"
}
