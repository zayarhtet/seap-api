package service

import "github.com/zayarhtet/seap-api/src/server/model/dao"

// MockDutyRepository is a mock implementation of DutyRepository interface
type MockDutyRepository struct {
	Duties   []dao.Duty
	MyDuties []dao.MyDuty
	Grades   []dao.Grading
}

func (m *MockDutyRepository) GetRowCount() *int64 {
	//TODO implement me
	var res int64 = 10
	return &res
}

func (m *MockDutyRepository) GetMemberWithDutiesByUsername(duties *dao.MemberWithDuties) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockDutyRepository) SaveDuty(duty *dao.Duty) error {
	return nil
}

func (m *MockDutyRepository) CreateGrades(gradings []*dao.Grading) error {
	//TODO implement me
	return nil
}

func (m *MockDutyRepository) GetAllGradingByDutyId(grading *dao.Grading, i int, i2 int) *[]dao.Grading {
	return &m.Grades
}

func (m *MockDutyRepository) GetDutyById(duty *dao.Duty) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockDutyRepository) GetDutiesByUsername(duty *dao.MyDuty) *[]dao.MyDuty {
	//TODO implement me
	return &m.MyDuties
}

func (m *MockDutyRepository) UpdateGrading(m2 map[string]any, grading *dao.Grading) error {
	//TODO implement me
	return nil
}

func (m *MockDutyRepository) InsertGivenFilesMetadata(files []*dao.GivenFile) error {
	//TODO implement me
	return nil
}

func (m *MockDutyRepository) GetGivenFileById(file *dao.GivenFile) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockDutyRepository) GetSubmittedFileById(file *dao.SubmittedFile) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockDutyRepository) GetGradingByStructCondition(grading *dao.Grading, grading2 *dao.Grading) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockDutyRepository) InsertSubmittedFilesMetadata(files []*dao.SubmittedFile) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockDutyRepository) GetAllSubmittedFilesMetadata(file *dao.SubmittedFile) *[]dao.SubmittedFile {
	//TODO implement me
	panic("implement me")
}

func (m *MockDutyRepository) DeleteSubmittedFileById(file *dao.SubmittedFile) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockDutyRepository) DeleteDutyById(duty *dao.Duty) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockDutyRepository) GetAllDuties(offset, limit int) *[]dao.Duty {
	return &m.Duties
}
