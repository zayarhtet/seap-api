package dao

import "time"

type SubmittedFile struct {
	FileId      string    `gorm:"primary_key" json:"fileId"`
	GradingId   string    `json:"-"`
	Grading_    Grading   `gorm:"foreignKey:GradingId;references:GradingId" json:"grade"`
	FilePath    string    `json:"filePath"`
	SubmittedAt time.Time `json:"submittedAt"`
}

func (SubmittedFile) TableName() string {
	return "submitted_file"
}
