package repository

import (
	"github.com/zayarhtet/seap-api/src/server/model/dao"
)

type RoleRepository interface {
	GetAllRoles(int, int) *[]dao.Role
	GetRowCount() *int64
	GetRoleById(uint) (*dao.Role, error)
}

type RoleRepositoryImpl struct{}

func (r RoleRepositoryImpl) GetAllRoles(offset, limit int) *[]dao.Role {
	var roles []dao.Role
	dc.GetAllByPagination(&roles, offset, limit, &dao.Role{})
	return &roles
}

func (r RoleRepositoryImpl) GetRowCount() *int64 {
	var count int64
	dc.GetRowCount("role", &count)
	return &count
}

func (r RoleRepositoryImpl) GetRoleById(id uint) (*dao.Role, error) {
	role := dao.Role{
		RoleId: id,
	}
	dr := dc.GetById(&role, &dao.Role{})
	return &role, dr.Error
}
