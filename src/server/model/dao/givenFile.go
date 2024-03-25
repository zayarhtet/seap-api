package dao

type GivenFile struct {
	FileId string `gorm:"primary_key" json:"fileId"`
	DutyId string `json:"-"`
	//Grading_    Grading `gorm:"foreignKey:GradingId;references:GradingId" json:"grade"`
	FilePath string `json:"filePath"`
}

func (GivenFile) TableName() string {
	return "given_file"
}
