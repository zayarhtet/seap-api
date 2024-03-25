package dao

import (
	"time"
)

type Role struct {
	RoleId     uint      `gorm:"primary_key" json:"roleId"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	//Members  []MemberChild `gorm:"foreignKey:Username" json:"members"`
}

func (Role) TableName() string {
	return "role"
}
