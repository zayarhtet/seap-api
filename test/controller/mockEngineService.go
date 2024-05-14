package controller

import "github.com/zayarhtet/seap-api/src/server/model/dto"

type MockEngineService struct {
	Data any
	Err  error
}

func (m *MockEngineService) GetPluginListResponse() (dto.Response, error) {
	panic("implement me")
}

func (m *MockEngineService) ExecuteSubmittedFile(dutyId string) (dto.Response, error) {
	return m.Data, m.Err
}
