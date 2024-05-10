package controller

import (
	"mime/multipart"

	"github.com/zayarhtet/seap-api/src/server/model/dto"
)

type MockFamilyService struct {
	Data any   // Predefined data to return
	Err  error // Predefined error to return
}

func (m *MockFamilyService) GetAllFamiliesResponse(i int, i2 int) (dto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockFamilyService) GetAllFamiliesWithMembersResponse(i int, i2 int) (dto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockFamilyService) GetMemberByIdWithFamiliesResponse(s string) (dto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockFamilyService) GetFamiliesOnlyByUsername(s string) (dto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockFamilyService) GetFamilyByIdWithMembers(s string) (dto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockFamilyService) GetFamilyByIdWithDuties(s string, s2 string) (dto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockFamilyService) SaveNewFamily(s string, request dto.NewFamilyRequest, header *multipart.FileHeader) (dto.Response, error) {
	return m.Data, m.Err
}

func (m *MockFamilyService) AddMemberToFamily(request dto.MemberToFamilyRequest) (dto.Response, error) {
	return m.Data, m.Err
}

func (m *MockFamilyService) RoleInFamily(s string, s2 string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockFamilyService) IsTutorInFamily(s string, s2 string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockFamilyService) GetFamilyProfileImagePath(s string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockFamilyService) DeleteFamilyResponse(s string) (dto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockFamilyService) GetMyRoleInFamily(userID string, famID string) (dto.Response, error) {
	return m.Data, m.Err
}
