package repository

import (
	"github.com/zayarhtet/seap-api/src/server/model/dao"
)

type RoleRepository interface {
	GetAllRoles() (*[]dao.Role, uint)
	// GetRoleById(id int) (dao.Role, error)
}

type RoleRepositoryImpl struct{}

func (r RoleRepositoryImpl) GetAllRoles() (*[]dao.Role, uint) {
	roles := []dao.Role{}
	dc.getAll(&roles)
	return &roles, 1
}
