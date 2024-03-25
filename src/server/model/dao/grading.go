package dao

type Grading struct {
	GradingId        string      `gorm:"primary_key" json:"gradingId"`
	DutyId           string      `json:"dutyId"`
	Duty_            DutyChild   `gorm:"foreignKey:DutyId;references:DutyId" json:"duty"`
	Username         string      `json:"username"`
	Member_          MemberChild `gorm:"foreignKey:Username;references:Username" json:"member"`
	FamilyId         string      `json:"familyId"`
	Submitted        bool        `json:"isSubmitted"`
	Points           int         `json:"points"`
	IsLate           bool        `json:"isLate"`
	IsPassed         bool        `json:"isPassed"`
	GradeComment     string      `json:"gradeComment"`
	ExecutionComment string      `json:"executionComment"`
}

func (Grading) TableName() string {
	return "grading"
}
