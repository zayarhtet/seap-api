package service

import (
	"path/filepath"

	"github.com/zayarhtet/seap-api/src/engine"
	"github.com/zayarhtet/seap-api/src/server/model/dao"
	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"github.com/zayarhtet/seap-api/src/server/repository"
	"github.com/zayarhtet/seap-api/src/util"
)

type EngineService interface {
	ExecuteSubmittedFile(string) (dto.Response, error)
	GetPluginListResponse() (dto.Response, error)
}

type engineServiceImpl struct {
	dr repository.DutyRepository
	fr repository.FamilyRepository
}

func (es engineServiceImpl) ExecuteSubmittedFile(dutyId string) (dto.Response, error) {
	var duty *dao.Duty = &dao.Duty{
		DutyId: dutyId,
	}
	err := es.dr.GetDutyById(duty)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(404, "record not found.")), err
	}
	go engine.ExecuteDuty("fpclean", filepath.Join(util.ABSOLUTE_SUBMITTED_STORAGE_PATH(), dutyId), filepath.Join(util.ABSOLUTE_INPUT_FILE_PATH(), dutyId))
	return "EXECUTING", nil
}

func (es engineServiceImpl) GetPluginListResponse() (dto.Response, error) {
	plugins := engine.GetPluginList()
	return BeforeDataResponse[string](plugins, 1), nil
}

func NewEngineService() EngineService {
	return &engineServiceImpl{dr: repository.DutyRepositoryImpl{}, fr: repository.FamilyRepositoryImpl{}}
}
