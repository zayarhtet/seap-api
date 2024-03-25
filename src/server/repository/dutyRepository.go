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
}

type DutyRepositoryImpl struct{}

func (d DutyRepositoryImpl) GetAllDuties(offset, limit int) *[]dao.Duty {
	var duties []dao.Duty
	dc.getAllByPagination(&duties, offset, limit, &dao.Duty{}, "Family_", "Files")
	return &duties
}
func (d DutyRepositoryImpl) GetRowCount() *int64 {
	var count int64
	dc.getRowCount("duty", &count)
	return &count
}

func (d DutyRepositoryImpl) GetMemberWithDutiesByUsername(member *dao.MemberWithDuties) error {
	return dc.getById(member, &dao.MemberWithDuties{}, "Duties.Duty_", "Duties").Error
}

func (d DutyRepositoryImpl) SaveDuty(duty *dao.Duty) error {
	return dc.insertOne(duty).Error
}

func (d DutyRepositoryImpl) CreateGrades(grades []*dao.Grading) error {
	return dc.insertAll(grades).Error
}

func (d DutyRepositoryImpl) GetAllGradingByDutyId(condition *dao.Grading, offset, limit int) *[]dao.Grading {
	var gradings []dao.Grading
	dc.getAllByPaginationWithCondition(&gradings, offset, limit, condition, &dao.Grading{}, "Duty_", "Member_")
	return &gradings
}

func (d DutyRepositoryImpl) GetDutyById(duty *dao.Duty) error {
	return dc.getById(duty, &dao.Duty{}, "Family_", "Files").Error
}
