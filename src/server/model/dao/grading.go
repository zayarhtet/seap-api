package dao

type Grading struct {
	GradingId        string          `gorm:"primary_key" json:"gradingId"`
	DutyId           string          `json:"dutyId"`
	Duty_            DutyChild       `gorm:"foreignKey:DutyId;references:DutyId" json:"duty"`
	Username         string          `json:"username"`
	Member_          MemberChild     `gorm:"foreignKey:Username;references:Username" json:"member"`
	FamilyId         string          `json:"familyId"`
	Submitted        bool            `json:"isSubmitted"`
	Points           int             `json:"points"`
	IsLate           bool            `json:"isLate"`
	IsPassed         bool            `json:"isPassed"`
	HasGraded        bool            `json:"hasGraded"`
	GradeComment     string          `json:"gradeComment"`
	ExecutionComment string          `json:"executionComment"`
	ReportPath       string          `json:"-"`
	Files            []SubmittedFile `gorm:"foreignKey:GradingId" json:"files"`
}

type GradingForFamily struct {
	GradingId string          `gorm:"primary_key" json:"-"`
	DutyId    string          `json:"-"`
	Duty_     DutiesForFamily `gorm:"foreignKey:DutyId;references:DutyId" json:"duty"`
	Username  string          `json:"-"`
	FamilyId  string          `json:"-"`
	Submitted bool            `json:"isSubmitted"`
}

type MyDuty struct {
	GradingId string          `gorm:"primary_key" json:"-"`
	DutyId    string          `json:"-"`
	Duty_     DutiesForFamily `gorm:"foreignKey:DutyId;references:DutyId" json:"duty"`
	Username  string          `json:"-"`
	FamilyId  string          `json:"familyId"`
	Family_   FamilyForGrade  `gorm:"foreignKey:FamilyId;references:FamilyId" json:"family"`
	Submitted bool            `json:"isSubmitted"`
}

func (Grading) TableName() string {
	return "grading"
}
func (GradingForFamily) TableName() string {
	return "grading"
}
func (MyDuty) TableName() string {
	return "grading"
}
