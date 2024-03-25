package dao

import "time"

type Family struct {
	FamilyId   string    `gorm:"column:family_id;primary_key" json:"familyId"`
	FamilyName string    `json:"name"`
	FamilyInfo string    `json:"info"`
	FamilyIcon string    `json:"icon"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

func (Family) TableName() string {
	return "family"
}

type FamilyWithMembers struct {
	FamilyId   string            `gorm:"column:family_id;primary_key" json:"familyId"`
	FamilyName string            `json:"name"`
	FamilyInfo string            `json:"info"`
	FamilyIcon string            `json:"icon"`
	Members    []MemberForFamily `gorm:"foreignKey:FamilyId" json:"members"`
}

type FamilyWithDuties struct {
	FamilyId   string            `gorm:"column:family_id;primary_key" json:"familyId"`
	FamilyName string            `json:"name"`
	FamilyInfo string            `json:"info"`
	FamilyIcon string            `json:"icon"`
	Duties     []DutiesForFamily `gorm:"foreignKey:FamilyId" json:"families"`
}

func (FamilyWithDuties) TableName() string {
	return "family"
}

func (FamilyWithMembers) TableName() string {
	return "family"
}
