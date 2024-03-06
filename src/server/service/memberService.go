package service

import (
	"github.com/zayarhtet/seap-api/src/server/auth"
	"github.com/zayarhtet/seap-api/src/server/model/dao"
	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"github.com/zayarhtet/seap-api/src/server/repository"
	"github.com/zayarhtet/seap-api/src/server/util"
)

type MemberService interface {
	SignUp(dto.SignUpRequest) (dto.Response, error)
	Login(dto.LoginRequest) (dto.Response, error)
	GetAllMembersResponse(int, int) (dto.Response, error)
}

type memberServiceImpl struct {
	mr repository.MemberRepository
	cr repository.CredentialRepository
}

func (ms memberServiceImpl) GetRowCount() *int64 {
	return ms.mr.GetRowCount()
}

func (ms memberServiceImpl) SignUp(request dto.SignUpRequest) (dto.Response, error) {
	var credential *dao.Credential
	var err error
	credential, err = dao.NewCredential(request.Password)

	if err != nil {
		return dto.NewErrorResponse("hello"), err
	}

	var credentialId string
	credentialId, err = ms.cr.SaveCredential(credential)
	if err != nil {
		return dto.NewErrorResponse("hello"), err
	}

	member := &dao.Member{
		Username:     request.UserName,
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		Email:        request.Email,
		CredentialId: credentialId,
		RoleId:       2,
	}
	member, err = ms.mr.SaveMember(member)
	if err != nil {
		ms.cr.DeleteCredential(credentialId)
		return BeforeErrorResponse(PrepareErrorMap(409, "Username or Email already exists.")), err
	}
	newResp := BeforeDataResponse[dao.Member](&[]dao.Member{*member}, 1)

	return newResp, nil
}

func (ms memberServiceImpl) Login(request dto.LoginRequest) (dto.Response, error) {
	var newResp dto.Response
	var err error
	var user *dao.Member = &dao.Member{
		Username: request.Username,
	}
	err = ms.mr.GetMemberByUsername(user)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(404, "Username does not exist.")), err
	}
	var credential *dao.Credential = &dao.Credential{
		CredentialId: user.CredentialId,
	}
	err = ms.cr.GetCredentialById(credential)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(400, "Password is incorrect.")), err
	}

	err = util.ValidatePassword(request.Password, credential.Password)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(400, "Password is incorrect.")), err
	}

	newResp = auth.GenerateToken(user.Username, user.Role.Name)

	return newResp, nil
}

func (ms memberServiceImpl) GetAllMembersResponse(size, page int) (dto.Response, error) {
	var newResp dto.Response

	total, offset := calculateOffset(ms, size, page)

	var data *[]dao.Member
	if offset == -1 {
		data = &[]dao.Member{}
	} else {
		data = ms.mr.GetAllMembers(offset, size)
	}

	newResp = BeforeDataResponse[dao.Member](data, *total, size, page)

	return newResp, nil
}

func NewMemberService() MemberService {
	return &memberServiceImpl{mr: repository.MemberRepositoryImpl{}, cr: repository.CredentialRepositoryImpl{}}
}
