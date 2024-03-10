package service

import (
	"github.com/zayarhtet/seap-api/src/server/model/dao"
	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"github.com/zayarhtet/seap-api/src/server/repository"
)

type FamilyService interface {
	GetAllFamiliesResponse(int, int) (dto.Response, error)
	GetAllFamiliesWithMembersResponse(int, int) (dto.Response, error)
}

type familyServiceImpl struct {
	fr repository.FamilyRepository
}

func (fs familyServiceImpl) GetRowCount() *int64 {
	return fs.fr.GetRowCount()
}

func (fs familyServiceImpl) GetAllFamiliesResponse(size, page int) (dto.Response, error) {
	var newResp dto.Response
	total, offset := calculateOffset(fs, size, page)

	var data *[]dao.Family
	if offset == -1 {
		data = &[]dao.Family{}
	} else {
		data = fs.fr.GetAllFamilies(offset, size)
	}

	newResp = BeforeDataResponse[dao.Family](data, *total, size, page)

	return newResp, nil
}

func (fs familyServiceImpl) GetAllFamiliesWithMembersResponse(size, page int) (dto.Response, error) {
	var newResp dto.Response
	total, offset := calculateOffset(fs, size, page)

	var data *[]dao.FamilyWithMembers
	if offset == -1 {
		data = &[]dao.FamilyWithMembers{}
	} else {
		data = fs.fr.GetAllFamiliesWithMembers(offset, size)
	}

	newResp = BeforeDataResponse[dao.FamilyWithMembers](data, *total, size, page)

	return newResp, nil
}

func NewFamilyService() FamilyService {
	return &familyServiceImpl{fr: repository.FamilyRepositoryImpl{}}
}
