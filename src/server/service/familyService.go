package service

import (
	"github.com/zayarhtet/seap-api/src/server/model/dao"
	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"github.com/zayarhtet/seap-api/src/server/repository"
	"github.com/zayarhtet/seap-api/src/server/util"
)

type FamilyService interface {
	GetAllFamiliesResponse(int, int) (dto.Response, error)
	GetAllFamiliesWithMembersResponse(int, int) (dto.Response, error)
	GetMemberByIdWithFamiliesResponse(string) (dto.Response, error)
	GetFamiliesOnlyByUsername(string) (dto.Response, error)
	GetFamilyById(familyId string) (dto.Response, error)
	SaveNewFamily(string, dto.NewFamilyRequest) (dto.Response, error)
	AddMemberToFamily(dto.MemberToFamilyRequest) (dto.Response, error)
	IsTutorInFamily(string, string) (bool, error)
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

func (fs familyServiceImpl) GetMemberByIdWithFamiliesResponse(username string) (dto.Response, error) {
	var member *dao.MemberWithFamilies = &dao.MemberWithFamilies{
		Username: username,
	}

	err := fs.fr.GetMemberByIdWithFamilies(member)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(404, "Username does not exist.")), err
	}

	newResp := BeforeDataResponse[dao.MemberWithFamilies](&[]dao.MemberWithFamilies{*member}, 1)
	return newResp, nil
}

func (fs familyServiceImpl) GetFamiliesOnlyByUsername(username string) (dto.Response, error) {
	var member *dao.MemberWithFamilies = &dao.MemberWithFamilies{
		Username: username,
	}

	err := fs.fr.GetMemberByIdWithFamilies(member)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(404, "Username does not exist.")), err
	}

	newResp := BeforeDataResponse[dao.FamilyForMember](&member.Families, 1)
	return newResp, nil
}

func (fs familyServiceImpl) GetFamilyById(familyId string) (dto.Response, error) {
	var member *dao.FamilyWithMembers = &dao.FamilyWithMembers{
		FamilyId: familyId,
	}
	err := fs.fr.GetFamilyById(member)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(404, "Username does not exist.")), err
	}

	newResp := BeforeDataResponse[dao.FamilyWithMembers](&[]dao.FamilyWithMembers{*member}, 1)
	return newResp, nil
}

func (fs familyServiceImpl) SaveNewFamily(username string, input dto.NewFamilyRequest) (dto.Response, error) {
	var id string = util.NewUUID()
	var family *dao.Family = &dao.Family{
		FamilyId:   id,
		FamilyName: input.FamilyName,
		FamilyIcon: input.FamilyIcon,
		FamilyInfo: input.FamilyInfo,
	}
	err := fs.fr.SaveNewFamily(family)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(404, err.Error())), err
	}
	if username == "admin" {
		return fs.GetFamilyById(id)
	}
	return fs.AddMemberToFamily(dto.MemberToFamilyRequest{FamilyId: id, Username: username, RoleId: 1})
}

func (fs familyServiceImpl) AddMemberToFamily(input dto.MemberToFamilyRequest) (dto.Response, error) {
	err := fs.fr.SaveNewMember(&input)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(404, err.Error())), err
	}
	return fs.GetFamilyById(input.FamilyId)
}

func (fs familyServiceImpl) IsTutorInFamily(username, familyId string) (bool, error) {
	var rq *dto.MemberToFamilyRequest = &dto.MemberToFamilyRequest{
		FamilyId: familyId,
		Username: username,
	}
	err := fs.fr.GetMemberRoleInFamily(rq)
	if err != nil {
		return false, err
	}
	return rq.RoleId <= 1, nil
}

func NewFamilyService() FamilyService {
	return &familyServiceImpl{fr: repository.FamilyRepositoryImpl{}}
}
