package service

import (
	"path/filepath"

	"github.com/zayarhtet/seap-api/src/engine"
	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"github.com/zayarhtet/seap-api/src/server/repository"
	"github.com/zayarhtet/seap-api/src/server/util"
)

type EngineService interface {
	ExecuteSubmittedFile(string) (dto.Response, error)
}

type engineServiceImpl struct {
	dr repository.DutyRepository
	fr repository.FamilyRepository
}

func (es engineServiceImpl) ExecuteSubmittedFile(dutyId string) (dto.Response, error) {
	go engine.ExecuteDuty("fpclean", filepath.Join(util.ABSOLUTE_SUBMITTED_STORAGE_PATH, "78910bfe-8d3a-47a2-bdb3-5b2b54bd489c"))
	return "EXECUTING", nil
}

func NewEngineService() EngineService {
	return &engineServiceImpl{dr: repository.DutyRepositoryImpl{}, fr: repository.FamilyRepositoryImpl{}}
}
