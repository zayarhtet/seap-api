package dao

import (
	"time"

	"github.com/zayarhtet/seap-api/src/server/util"
)

type Duty struct {
	DutyId             string        `gorm:"primary_key" json:"dutyId"`
	Title              string        `json:"title" binding:"required"`
	Instruction        string        `json:"instruction"`
	PublishingDate     util.WrapTime `json:"publishedAt" binding:"required"`
	DeadlineDate       util.WrapTime `json:"dueDate" binding:"required"`
	ClosingDate        util.WrapTime `json:"closingDate" binding:"required"`
	FamilyId           string        `json:"familyId" binding:"required"`
	Family_            Family        `gorm:"foreignKey:FamilyId;references:FamilyId" json:"family"`
	PointSystem        bool          `json:"isPointSystem"`
	PossiblePoints     int           `json:"totalPoints"`
	MultipleSubmission bool          `json:"multipleSubmission"`
	Files              []GivenFile   `gorm:"foreignKey:DutyId" json:"files"`
}

func (Duty) TableName() string {
	return "duty"
}

type DutyChild struct {
	DutyId             string `gorm:"primary_key" json:"dutyId"`
	Title              string `json:"title"`
	PointSystem        bool   `json:"isPointSystem"`
	PossiblePoints     int    `json:"totalPoints"`
	MultipleSubmission bool   `json:"multipleSubmission"`
}

func (DutyChild) TableName() string {
	return "duty"
}

type DutiesForFamily struct {
	DutyId         string    `gorm:"primary_key" json:"dutyId"`
	Title          string    `json:"title"`
	PublishingDate time.Time `json:"publishedAt"`
	DeadlineDate   time.Time `json:"dueDate"`
	ClosingDate    time.Time `json:"closingDate"`
	FamilyId       string    `json:"-"`
}

func (DutiesForFamily) TableName() string {
	return "duty"
}

type DutiesForMember struct {
	GradingId string    `gorm:"primary_key" json:"-"`
	Username  string    `gorm:"column:username;" json:"-"`
	DutyId    string    `json:"-"`
	Duty_     DutyChild `gorm:"foreignKey:DutyId;references:DutyId" json:"duty"`
	FamilyId  string    `json:"familyId"`
}

func (DutiesForMember) TableName() string {
	return "grading"
}
