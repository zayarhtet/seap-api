package controller

import "github.com/zayarhtet/seap-api/src/server/model/dto"

type MockEngineService struct {
	Data any   // Predefined data to return
	Err  error // Predefined error to return
}

func (m *MockEngineService) GetPluginListResponse() (dto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineService) ExecuteSubmittedFile(dutyId string) (dto.Response, error) {
	return m.Data, m.Err
}
