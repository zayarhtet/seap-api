package controller

import "github.com/zayarhtet/seap-api/src/server/model/dto"

type MockMemberService struct {
	Data interface{} // Predefined data to return
	Err  error       // Predefined error to return
}

func (m *MockMemberService) SignUp(request dto.SignUpRequest) (dto.Response, error) {
	panic("implement me")
}

func (m *MockMemberService) Login(request dto.LoginRequest) (dto.Response, error) {
	panic("implement me")
}

func (m *MockMemberService) GetAllMembersResponse(i int, i2 int) (dto.Response, error) {
	return m.Data, m.Err
}

func (m *MockMemberService) GetAllMembersWithFamiliesResponse(i int, i2 int) (dto.Response, error) {
	panic("implement me")
}

func (m *MockMemberService) GetMemberByIdResponse(s string) (dto.Response, error) {
	return m.Data, m.Err
}

func (m *MockMemberService) DeleteMemberResponse(s string) (dto.Response, error) {
	panic("implement me")
}

func (m *MockMemberService) GrantRoleResponse(s string, i int) (dto.Response, error) {
	return m.Data, m.Err
}
