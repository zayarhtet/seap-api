package service

import (
	"github.com/zayarhtet/seap-api/src/server/repository"
	"github.com/zayarhtet/seap-api/src/server/model/dao"
	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"github.com/gin-gonic/gin"
)

type RoleService interface {
	GetAllRolesResponse(*gin.Context) (*dto.Response,error)
}

type roleServiceImpl struct {
	rp repository.RoleRepository
}

func (rs roleServiceImpl) GetAllRolesResponse(context *gin.Context) (*dto.Response, error) {
	newResp := dto.NewResponse()

	var data *[]dao.Role
	data, newResp.Total = rs.rp.GetAllRoles()
	newResp.Data = data
	newResp.Size = uint(len(*data))

	newResp.Total = 3
	newResp.StartAt = 0
	newResp.Username = "HELLO"

	return newResp, nil
}

func NewRoleService() RoleService {
	return &roleServiceImpl{rp:  repository.RoleRepositoryImpl{}}
}