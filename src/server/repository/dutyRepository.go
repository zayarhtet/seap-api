package repository

import "github.com/zayarhtet/seap-api/src/server/model/dao"

type DutyRepository interface {
	GetAllDuties(int, int) *[]dao.Duty
	GetRowCount() *int64
	GetMemberWithDutiesByUsername(*dao.MemberWithDuties) error
	SaveDuty(*dao.Duty) error
	CreateGrades([]*dao.Grading) error
	GetAllGradingByDutyId(*dao.Grading, int, int) *[]dao.Grading
	GetDutyById(*dao.Duty) error
	GetDutiesByUsername(*dao.MyDuty) *[]dao.MyDuty
	UpdateGrading(map[string]any, *dao.Grading) error

	InsertGivenFilesMetadata([]*dao.GivenFile) error
	GetGivenFileById(*dao.GivenFile) error
	GetSubmittedFileById(*dao.SubmittedFile) error
	GetGradingByStructCondition(*dao.Grading, *dao.Grading) error
	InsertSubmittedFilesMetadata([]*dao.SubmittedFile) error
	GetAllSubmittedFilesMetadata(*dao.SubmittedFile) *[]dao.SubmittedFile
	DeleteSubmittedFileById(*dao.SubmittedFile) error

	DeleteDutyById(*dao.Duty) error
}

type DutyRepositoryImpl struct{}

func (d DutyRepositoryImpl) GetAllDuties(offset, limit int) *[]dao.Duty {
	var duties []dao.Duty
	dc.GetAllByPagination(&duties, offset, limit, &dao.Duty{}, "Family_", "Files")
	return &duties
}
func (d DutyRepositoryImpl) GetRowCount() *int64 {
	var count int64
	dc.GetRowCount("duty", &count)
	return &count
}

func (d DutyRepositoryImpl) GetMemberWithDutiesByUsername(member *dao.MemberWithDuties) error {
	return dc.GetById(member, &dao.MemberWithDuties{}, "Duties.Duty_", "Duties").Error
}

func (d DutyRepositoryImpl) GetDutiesByUsername(condition *dao.MyDuty) *[]dao.MyDuty {
	var duties []dao.MyDuty
	dc.GetAllByPaginationWithCondition(&duties, 0, 100, condition, &dao.MyDuty{}, "Duty_", "Family_")
	return &duties
}

func (d DutyRepositoryImpl) SaveDuty(duty *dao.Duty) error {
	return dc.InsertOne(duty).Error
}

func (d DutyRepositoryImpl) CreateGrades(grades []*dao.Grading) error {
	if len(grades) == 0 {
		return nil
	}
	return dc.InsertAll(grades).Error
}

func (d DutyRepositoryImpl) GetAllGradingByDutyId(condition *dao.Grading, offset, limit int) *[]dao.Grading {
	var gradings []dao.Grading
	dc.GetAllByPaginationWithCondition(&gradings, offset, limit, condition, &dao.Grading{}, "Duty_", "Member_", "Files")
	return &gradings
}

func (d DutyRepositoryImpl) GetDutyById(duty *dao.Duty) error {
	return dc.GetById(duty, &dao.Duty{}, "Family_", "Files").Error
}

func (d DutyRepositoryImpl) InsertGivenFilesMetadata(metadata []*dao.GivenFile) error {
	if len(metadata) == 0 {
		return nil
	}
	return dc.InsertAll(metadata).Error
}

func (d DutyRepositoryImpl) GetGivenFileById(gFile *dao.GivenFile) error {
	return dc.GetById(gFile, &dao.GivenFile{}).Error
}

func (d DutyRepositoryImpl) GetSubmittedFileById(sFile *dao.SubmittedFile) error {
	return dc.GetById(sFile, &dao.SubmittedFile{}).Error
}

func (d DutyRepositoryImpl) GetGradingByStructCondition(dest *dao.Grading, condition *dao.Grading) error {
	return dc.GetOneByStructCondition(dest, condition).Error
}

func (d DutyRepositoryImpl) InsertSubmittedFilesMetadata(metadata []*dao.SubmittedFile) error {
	if len(metadata) == 0 {
		return nil
	}
	return dc.InsertAll(metadata).Error
}

func (d DutyRepositoryImpl) GetAllSubmittedFilesMetadata(condition *dao.SubmittedFile) *[]dao.SubmittedFile {
	var files []dao.SubmittedFile
	dc.GetAllByStructCondition(&files, condition, &dao.SubmittedFile{})
	return &files
}
func (d DutyRepositoryImpl) DeleteSubmittedFileById(sFile *dao.SubmittedFile) error {
	return dc.DeleteOneById(sFile).Error
}

func (d DutyRepositoryImpl) DeleteDutyById(duty *dao.Duty) error {
	return dc.DeleteOneById(duty).Error
}

func (d DutyRepositoryImpl) UpdateGrading(fields map[string]any, model *dao.Grading) error {
	return dc.UpdateModelByMap(fields, model).Error
}
