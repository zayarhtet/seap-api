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

func TestGetMyMember_Success(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("username", "john")

	mockService := &service.MockMemberService{
		Data: dao.Member{Username: "johndoe", FirstName: "John", LastName: "Doe"},
		Err:  nil,
	}

	ic := controller.NewIndividualController(mockService, nil, nil, nil)

	ic.GetMyMember(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response dao.Member
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	expectedMember := dao.Member{Username: "johndoe", FirstName: "John", LastName: "Doe"}
	assert.Equal(t, expectedMember, response)
}

func TestGetMyMember_Error(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("username", "john")

	mockService := &service.MockMemberService{
		Data: dao.Member{},
		Err:  errors.New("mock error"),
	}

	ic := controller.NewIndividualController(mockService, nil, nil, nil)

	ic.GetMyMember(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "mock error")
}

func TestGetMyRoleInFamily_Success(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("username", "john")
	c.Params = []gin.Param{{Key: "famId", Value: "1"}}

	mockService := &service.MockFamilyService{
		Data: dao.Role{RoleId: 1, Name: "tutor"}, // Dummy role data
		Err:  nil,                                // No error
	}

	ic := controller.NewIndividualController(nil, nil, mockService, nil)

	ic.GetMyRoleInFamily(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response dao.Role
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	expectedRole := dao.Role{RoleId: 1, Name: "tutor"}
	assert.Equal(t, expectedRole, response)
}

func TestGetMyRoleInFamily_Error(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("username", "john")
	c.Params = []gin.Param{{Key: "famId", Value: "1"}}

	mockService := &service.MockFamilyService{
		Data: dao.Role{},
		Err:  errors.New("mock error"),
	}

	ic := controller.NewIndividualController(nil, nil, mockService, nil)

	ic.GetMyRoleInFamily(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "mock error")
}
