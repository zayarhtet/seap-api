package repository

import (
	"fmt"
	"github.com/zayarhtet/seap-api/src/server/model/dao"
)

type RoleRepository interface {
	GetAllRoles(int, int) *[]dao.Role
	GetRowCount() *int64
	GetRoleById(id uint) (*dao.Role, error)
}

type RoleRepositoryImpl struct{}

func (r RoleRepositoryImpl) GetAllRoles(offset, limit int) *[]dao.Role {
	var roles []dao.Role
	dc.getAllByPagination(&roles, offset, limit)
	return &roles
}

func (r RoleRepositoryImpl) GetRowCount() *int64 {
	var count int64
	dc.getRowCount("role", &count)
	fmt.Println(count)
	return &count
}

func (r RoleRepositoryImpl) GetRoleById(id uint) (*dao.Role, error) {
	role := dao.Role{
		RoleId: id,
	}
	dr := dc.getById(&role)
	return &role, dr.Error
}
