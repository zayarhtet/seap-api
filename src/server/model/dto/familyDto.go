package dto

type MemberToFamilyRequest struct {
	FamilyId string `gorm:"primary_key" json:"familyId" binding:"required"`
	Username string `gorm:"primary_key" json:"username" binding:"required"`
	RoleId   int    `json:"roleId" binding:"required"`
}

func (MemberToFamilyRequest) TableName() string {
	return "family_member"
}

type NewFamilyRequest struct {
	FamilyName string `json:"familyName" binding:"required"`
	FamilyInfo string `json:"familyInfo" binding:"required"`
	FamilyIcon string `json:"icon"`
}
