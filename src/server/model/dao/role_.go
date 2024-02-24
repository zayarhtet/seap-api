package dao

import (
	"time"
)

type Role struct {
	RoleId    uint
	Name       string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func (Role) TableName() string {
	return "role"
}