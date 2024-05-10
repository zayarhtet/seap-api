package controller_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/zayarhtet/seap-api/src/server/controller"
	"github.com/zayarhtet/seap-api/src/server/model/dao"
	service "github.com/zayarhtet/seap-api/test/controller"
)

func TestGetAllRoles_Success(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockService := &service.MockRoleService{
		Data: []dao.Role{{RoleId: 1, Name: "tutor"}, {RoleId: 2, Name: "tutee"}},
		Err:  nil,
	}

	rc := controller.NewRoleController(mockService)

	rc.GetAllRoles(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []dao.Role
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	expectedRoles := []dao.Role{{RoleId: 1, Name: "tutor"}, {RoleId: 2, Name: "tutee"}}
	assert.Equal(t, expectedRoles, response)
}

func TestGetAllRoles_Error(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockService := &service.MockRoleService{
		Data: nil,
		Err:  errors.New("mock error"),
	}

	rc := controller.NewRoleController(mockService)

	rc.GetAllRoles(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "mock error")
}

func TestGetRoleById_Success(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	mockService := &service.MockRoleService{
		Data: dao.Role{RoleId: 1, Name: "Admin"},
		Err:  nil,
	}

	rc := controller.NewRoleController(mockService)

	rc.GetRoleById(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response dao.Role
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	expectedRole := dao.Role{RoleId: 1, Name: "Admin"}
	assert.Equal(t, expectedRole, response)
}

func TestGetRoleById_InvalidID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "invalid"}}

	mockService := &service.MockRoleService{}

	rc := controller.NewRoleController(mockService)
	rc.GetRoleById(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid syntax")
}

func TestGetRoleById_Error(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	mockService := &service.MockRoleService{
		Data: dao.Role{},
		Err:  errors.New("mock error"),
	}

	rc := controller.NewRoleController(mockService)

	rc.GetRoleById(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "mock error")
}
