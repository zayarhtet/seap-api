package service

import (
	"github.com/zayarhtet/seap-api/src/server/model/dao"
	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"github.com/zayarhtet/seap-api/src/server/repository"
)

type RoleService interface {
	GetAllRolesResponse(int, int) (dto.Response, error)
	GetRoleByIdResponse(uint) (dto.Response, error)
	GetRoleByMemberResponse(string) (dto.Response, error)
}

type roleServiceImpl struct {
	rp repository.RoleRepository
	mr repository.MemberRepository
}

func (rs roleServiceImpl) GetRowCount() *int64 {
	return rs.rp.GetRowCount()
}

func (rs roleServiceImpl) GetAllRolesResponse(size int, page int) (dto.Response, error) {
	var newResp dto.Response

	total, offset := calculateOffset(rs, size, page)
	var data *[]dao.Role
	if offset == -1 {
		data = &[]dao.Role{}
	} else {
		data = rs.rp.GetAllRoles(offset, size)
	}

	newResp = BeforeDataResponse[dao.Role](data, *total, size, page)

	if false {
		newResp = BeforeErrorResponse(PrepareErrorMap(404, "Not Found"))
	}

	return newResp, nil
}

func (rs roleServiceImpl) GetRoleByIdResponse(id uint) (dto.Response, error) {
	var newResp dto.Response

	var role, err = rs.rp.GetRoleById(id)
	if err != nil {
		newResp = BeforeErrorResponse(PrepareErrorMap(404, "Not Found"))
		return newResp, nil
	}
	var total int64 = 1
	newResp = BeforeDataResponse[dao.Role](&[]dao.Role{*role}, total)

	return newResp, nil
}

func (rs roleServiceImpl) GetRoleByMemberResponse(username string) (dto.Response, error) {
	var member *dao.Member = &dao.Member{
		Username: username,
	}
	err := rs.mr.GetMemberByUsername(member)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(404, "Username does not exist.")), err
	}
	newResp := BeforeDataResponse[dto.RoleDto](&[]dto.RoleDto{member.Role}, 1)
	return newResp, nil
}

func NewRoleService() RoleService {
	return &roleServiceImpl{rp: repository.RoleRepositoryImpl{}, mr: repository.MemberRepositoryImpl{}}
}
