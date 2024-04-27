package dto

type NewGradeRequest struct {
	Username string `json:"username" binding:"required"`
	//FamilyId     string `json:"familyId" binding:"required"`
	GradingId string `json:"gradingId" binding:"required"`
	//DutyId       string `json:"dutyId" binding:"required"`
	Points       int    `json:"points" binding:"required"`
	GradeComment string `json:"gradeComment"`
}
