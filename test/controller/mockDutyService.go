package controller

import (
	"mime/multipart"

	"github.com/zayarhtet/seap-api/src/server/model/dao"
	"github.com/zayarhtet/seap-api/src/server/model/dto"
)

type MockDutyService struct {
	GivenFileErr error
	InputFileErr error
	Data         any   // Predefined data to return
	Err          error // Predefined error to return
}

func (m *MockDutyService) GetAllDutiesResponse(i int, i2 int) (dto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockDutyService) GetAllDutiesByMemberResponse(s string) (dto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockDutyService) SaveNewDutyResponse(duty dao.Duty) (dto.Response, error) {
	return m.Data, m.Err
}

func (m *MockDutyService) GetGradingByDutyIdResponse(s string, i int, i2 int) (dto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockDutyService) GetDutyByIdResponse(s string) (dto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockDutyService) CreateGivenFiles(headers []*multipart.FileHeader, s string) error {
	return m.GivenFileErr
}

func (m *MockDutyService) GetGivenFilePath(s string, s2 string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockDutyService) UploadSubmittedFiles(headers []*multipart.FileHeader, s string, s2 string, s3 string) (dto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockDutyService) GetSubmittedFilePath(s string, s2 string, s3 string, s4 string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockDutyService) SubmitDutyResponse(s string, s2 string) error {
	return m.Err
}

func (m *MockDutyService) DeleteSubmittedFileResponse(s string, s2 string, s3 string) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockDutyService) DeleteDutyResponse(s string) (dto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockDutyService) GetMyGradingResponse(s string, s2 string) (dto.Response, error) {
	return m.Data, m.Err
}

func (m *MockDutyService) GetReportContent(s string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockDutyService) SaveInputFiles(headers []*multipart.FileHeader, s string) error {
	return m.InputFileErr
}

func (m *MockDutyService) AddNewGradeResponse(input dto.NewGradeRequest) (dto.Response, error) {
	return m.Data, m.Err
}
