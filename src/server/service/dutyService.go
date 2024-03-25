package service

import (
	"github.com/zayarhtet/seap-api/src/server/model/dao"
	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"github.com/zayarhtet/seap-api/src/server/repository"
	"github.com/zayarhtet/seap-api/src/server/util"
)

type DutyService interface {
	GetAllDutiesResponse(int, int) (dto.Response, error)
	GetAllDutiesByMemberResponse(string) (dto.Response, error)
	AddNewGradeResponse(dto.NewGradeRequest) (dto.Response, error)
	SaveNewDutyResponse(dao.Duty) (dto.Response, error)
	GetGradingByDutyIdResponse(string, int, int) (dto.Response, error)
	GetDutyByIdResponse(string) (dto.Response, error)
}

type dutyServiceImpl struct {
	dr repository.DutyRepository
	fr repository.FamilyRepository
}

func (ds dutyServiceImpl) GetRowCount() *int64 {
	return ds.dr.GetRowCount()
}

func (ds dutyServiceImpl) GetAllDutiesResponse(size, page int) (dto.Response, error) {
	var newResp dto.Response

	total, offset := calculateOffset(ds, size, page)

	var data *[]dao.Duty
	if offset == -1 {
		data = &[]dao.Duty{}
	} else {
		data = ds.dr.GetAllDuties(offset, size)
	}

	newResp = BeforeDataResponse[dao.Duty](data, *total, size, page)

	return newResp, nil
}

func (ds dutyServiceImpl) GetAllDutiesByMemberResponse(username string) (dto.Response, error) {
	var data *dao.MemberWithDuties = &dao.MemberWithDuties{Username: username}

	err := ds.dr.GetMemberWithDutiesByUsername(data)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(404, "Username does not exist.")), err
	}
	newResp := BeforeDataResponse[dao.MemberWithDuties](&[]dao.MemberWithDuties{*data}, 1)
	return newResp, nil
}

func (ds dutyServiceImpl) AddNewGradeResponse(request dto.NewGradeRequest) (dto.Response, error) {
	return nil, nil
}

func (ds dutyServiceImpl) SaveNewDutyResponse(duty dao.Duty) (dto.Response, error) {
	duty.DutyId = util.NewUUID()
	err := ds.dr.SaveDuty(&duty)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(400, err.Error())), err
	}

	var memberList *dao.FamilyWithMembers = &dao.FamilyWithMembers{
		FamilyId: duty.FamilyId,
	}
	err = ds.fr.GetFamilyById(memberList)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(400, err.Error())), err
	}

	var newGrades []*dao.Grading

	for _, member := range memberList.Members {
		if member.MemberRole.Name != "tutee" {
			continue
		}
		newGrades = append(newGrades, &dao.Grading{
			GradingId: util.NewUUID(),
			DutyId:    duty.DutyId,
			Username:  member.Username,
			FamilyId:  duty.FamilyId,
		})
	}

	err = ds.dr.CreateGrades(newGrades)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(400, err.Error())), err
	}
	return duty, err
}

func (ds dutyServiceImpl) GetGradingByDutyIdResponse(dutyId string, size, page int) (dto.Response, error) {
	var newResp dto.Response

	total, offset := calculateOffset(ds, size, page)

	var data *[]dao.Grading
	if offset == -1 {
		data = &[]dao.Grading{}
	} else {
		condition := &dao.Grading{DutyId: dutyId}
		data = ds.dr.GetAllGradingByDutyId(condition, offset, size)
	}

	newResp = BeforeDataResponse[dao.Grading](data, *total, size, page)

	return newResp, nil
}

func (ds dutyServiceImpl) GetDutyByIdResponse(dutyId string) (dto.Response, error) {
	var duty *dao.Duty = &dao.Duty{
		DutyId: dutyId,
	}
	err := ds.dr.GetDutyById(duty)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(404, "Username does not exist.")), err
	}
	newResp := BeforeDataResponse[dao.Duty](&[]dao.Duty{*duty}, 1)
	return newResp, nil
}

func NewDutyService() DutyService {
	return &dutyServiceImpl{dr: repository.DutyRepositoryImpl{}, fr: repository.FamilyRepositoryImpl{}}
}
