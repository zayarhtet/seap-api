package dao

import (
	"github.com/zayarhtet/seap-api/src/server/util"
)

type SubmittedFile struct {
	FileId    string `gorm:"primary_key" json:"fileId"`
	GradingId string `json:"gradingId"`
	//Grading_    Grading   `gorm:"foreignKey:GradingId;references:GradingId" json:"grade"`
	FilePath    string        `json:"filePath"`
	SubmittedAt util.WrapTime `json:"submittedAt"`
}

func (SubmittedFile) TableName() string {
	return "submitted_file"
}
