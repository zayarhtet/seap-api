package dto

type MemberDtoFamily struct {
	FirstName string `gorm:"column:first_name" json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `gorm:"primary_key" json:"username"`
	Email     string `json:"email"`
}

func (MemberDtoFamily) TableName() string {
	return "member"
}

type FamilyDtoMember struct {
	FamilyId   string `gorm:"column:family_id;primary_key" json:"familyId"`
	FamilyName string `json:"name"`
	FamilyInfo string `json:"info"`
	FamilyIcon string `json:"icon"`
}

func (FamilyDtoMember) TableName() string {
	return "family"
}
