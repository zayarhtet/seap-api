package controller

import (
	"github.com/zayarhtet/seap-api/src/server/model/dto"
)

type MockRoleService struct {
	Data any   // Predefined data to return
	Err  error // Predefined error to return
}

func (m *MockRoleService) GetAllRolesResponse(i int, i2 int) (dto.Response, error) {
	//TODO implement me
	return m.Data, m.Err
}

func (m *MockRoleService) GetRoleByIdResponse(u uint) (dto.Response, error) {
	//TODO implement me
	return m.Data, m.Err
}

func (m *MockRoleService) GetRoleByMemberResponse(s string) (dto.Response, error) {
	//TODO implement me
	panic("implement me")
}
