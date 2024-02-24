package dao

import (
	"time"
)

type Member struct {
	firstName     	string `gorm:"column:first_name"`
	lastName      	string
	username      	string
	email         	string
	credentialId  	string
	roleId 			uint
	CreatedAt  		time.Time
	ModifiedAt 		time.Time
}

func (Member) TableName() string {
	return "member"
}