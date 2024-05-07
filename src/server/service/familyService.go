package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/zayarhtet/seap-api/src/server/model/dao"
	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"github.com/zayarhtet/seap-api/src/server/repository"
	"github.com/zayarhtet/seap-api/src/util"
)

type FamilyService interface {
	GetAllFamiliesResponse(int, int) (dto.Response, error)
	GetAllFamiliesWithMembersResponse(int, int) (dto.Response, error)
	GetMemberByIdWithFamiliesResponse(string) (dto.Response, error)
	GetFamiliesOnlyByUsername(string) (dto.Response, error)
	GetFamilyByIdWithMembers(string) (dto.Response, error)
	GetFamilyByIdWithDuties(string, string) (dto.Response, error)
	SaveNewFamily(string, dto.NewFamilyRequest, *multipart.FileHeader) (dto.Response, error)
	AddMemberToFamily(dto.MemberToFamilyRequest) (dto.Response, error)
	RoleInFamily(string, string) (string, error)
	GetMyRoleInFamily(string, string) (dto.Response, error)
	IsTutorInFamily(string, string) (bool, error)
	GetFamilyProfileImagePath(string) (string, error)
	DeleteFamilyResponse(string) (dto.Response, error)
}

type familyServiceImpl struct {
	fr repository.FamilyRepository
	dr repository.DutyRepository
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

func (fs familyServiceImpl) GetFamilyByIdWithMembers(familyId string) (dto.Response, error) {
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

func (fs familyServiceImpl) GetFamilyByIdWithDuties(familyId string, username string) (dto.Response, error) {
	var member *dao.FamilyWithDuties = &dao.FamilyWithDuties{
		FamilyId: familyId,
	}
	role, err1 := fs.RoleInFamily(username, familyId)
	if err1 != nil {
		return BeforeErrorResponse(PrepareErrorMap(404, "Username does not exist.")), err1
	}
	var err error
	if role == "tutee" {
		err = fs.fr.GetFamilyByIdWithDutiesForTutee(member, username)
		util.RemoveElementsInPlace[dao.GradingForFamily](&(member.DutiesWithSubmission), func(grading dao.GradingForFamily) bool {
			return time.Now().After(grading.Duty_.PublishingDate)
		})
	} else {
		err = fs.fr.GetFamilyByIdWithDutiesForTutor(member)
	}
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(404, "Username does not exist.")), err
	}

	newResp := BeforeDataResponse[dao.FamilyWithDuties](&[]dao.FamilyWithDuties{*member}, 1)
	return newResp, nil
}

func (fs familyServiceImpl) GetMyRoleInFamily(username, familyId string) (dto.Response, error) {
	var rq *dao.MemberForFamily = &dao.MemberForFamily{
		FamilyId: familyId,
		Username: username,
	}
	err := fs.fr.GetMemberRoleInFamily(rq)
	if err != nil {
		return "", err
	}
	newResp := BeforeDataResponse[dto.RoleDto](&[]dto.RoleDto{rq.MemberRole}, 1)
	return newResp, nil
}

func (fs familyServiceImpl) SaveNewFamily(username string, input dto.NewFamilyRequest, file *multipart.FileHeader) (dto.Response, error) {
	var id string = util.NewUUID()
	var iconPath string = ""
	if file != nil {
		fmt.Println("profile picture submitted")
		if err := util.VerifyImageFile(file); err != nil {
			return BeforeErrorResponse(PrepareErrorMap(400, err.Error())), err
		}

		iconPath = id + filepath.Ext(file.Filename)
		err := util.SaveIcons(file, iconPath)
		if err != nil {
			return BeforeErrorResponse(PrepareErrorMap(500, err.Error())), err
		}
	}

	var family *dao.Family = &dao.Family{
		FamilyId:   id,
		FamilyName: input.FamilyName,
		FamilyIcon: iconPath,
		FamilyInfo: input.FamilyInfo,
	}
	err := fs.fr.SaveNewFamily(family)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(404, err.Error())), err
	}
	if username == "admin" {
		return fs.GetFamilyByIdWithMembers(id)
	}
	return fs.AddMemberToFamily(dto.MemberToFamilyRequest{FamilyId: id, Username: username, RoleId: 1})
}

func (fs familyServiceImpl) AddMemberToFamily(input dto.MemberToFamilyRequest) (dto.Response, error) {
	err := fs.fr.SaveNewMember(&input)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(404, err.Error())), err
	}
	err = fs.AddNewMemberToGrading(input)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(404, err.Error())), err
	}
	return fs.GetFamilyByIdWithMembers(input.FamilyId)
}

func (fs familyServiceImpl) AddNewMemberToGrading(input dto.MemberToFamilyRequest) error {
	if input.RoleId == 1 {
		return nil
	}
	var fam *dao.FamilyWithDuties = &dao.FamilyWithDuties{
		FamilyId: input.FamilyId,
	}

	err := fs.fr.GetFamilyByIdWithDutiesForTutor(fam)

	if err != nil {
		return err
	}
	var newGrades []*dao.Grading

	for _, v := range fam.Duties {
		if time.Now().After(v.ClosingDate) {
			continue
		}
		newGrades = append(newGrades, &dao.Grading{
			GradingId: util.NewUUID(),
			DutyId:    v.DutyId,
			Username:  input.Username,
			FamilyId:  input.FamilyId,
		})
	}

	err = fs.dr.CreateGrades(newGrades)
	if err != nil {
		return err
	}
	return nil
}

func (fs familyServiceImpl) RoleInFamily(username, familyId string) (string, error) {
	var rq *dao.MemberForFamily = &dao.MemberForFamily{
		FamilyId: familyId,
		Username: username,
	}
	err := fs.fr.GetMemberRoleInFamily(rq)
	if err != nil {
		return "", err
	}
	return rq.MemberRole.Name, nil
}

func (fs familyServiceImpl) IsTutorInFamily(username, familyId string) (bool, error) {
	role, err := fs.RoleInFamily(username, familyId)
	if err != nil {
		return false, err
	}
	return role == "tutor" || role == "admin", nil
}

func (fs familyServiceImpl) GetFamilyProfileImagePath(famId string) (string, error) {
	var member *dao.FamilyWithMembers = &dao.FamilyWithMembers{
		FamilyId: famId,
	}
	err := fs.fr.GetFamilyById(member)
	if err != nil {
		return "", err
	}
	absolutePath := util.GetFamilyIconAbsolutePath(member.FamilyIcon)
	if len(absolutePath) == 0 {
		return "", errors.New("404 not found")
	}
	return absolutePath, nil
}

func (fs familyServiceImpl) DeleteFamilyResponse(famId string) (dto.Response, error) {
	var family *dao.Family = &dao.Family{FamilyId: famId}
	err := fs.fr.DeleteFamilyById(family)
	if err != nil {
		return "", err
	}
	return "success", nil
}

func NewFamilyService() FamilyService {
	return &familyServiceImpl{fr: repository.FamilyRepositoryImpl{}, dr: repository.DutyRepositoryImpl{}}
}
