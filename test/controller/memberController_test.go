package controller_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/zayarhtet/seap-api/src/server/controller"
	"github.com/zayarhtet/seap-api/src/server/model/dao"
	service "github.com/zayarhtet/seap-api/test/controller"
)

func TestGetAllMembers_Success(t *testing.T) {
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	context.Request = httptest.NewRequest(http.MethodGet, "/members", nil)
	context.Set("size", "10")
	context.Set("page", "1")

	dummyMembers := []dao.Member{
		{Username: "johndoe", FirstName: "John", LastName: "Doe"},
		{Username: "janedoe", FirstName: "Jane", LastName: "Doe"},
	}

	mockService := &service.MockMemberService{
		Data: dummyMembers,
		Err:  nil,
	}
	controller := controller.NewMemberController(mockService)

	controller.GetAllMembers(context)

	if context.Writer.Status() != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, context.Writer.Status())
	}

	var response []dao.Member
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, dummyMembers, response)
}

func TestGetAllMembers_Error(t *testing.T) {
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	context.Request = httptest.NewRequest(http.MethodGet, "/members", nil)

	mockService := &service.MockMemberService{
		Data: nil,
		Err:  errors.New("mock error"),
	}
	controller := controller.NewMemberController(mockService)

	controller.GetAllMembers(context)

	if context.Writer.Status() != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, context.Writer.Status())
	}
}

func TestGetMemberById_Success(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "johndoe"}}

	mockService := &service.MockMemberService{
		Data: dao.Member{Username: "johndoe", FirstName: "John", LastName: "Doe"},
		Err:  nil,
	}

	mc := controller.NewMemberController(mockService)

	mc.GetMemberById(c)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := dao.Member{Username: "johndoe", FirstName: "John", LastName: "Doe"}
	var response dao.Member
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
}

func TestGetMemberById_Error(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	mockService := &service.MockMemberService{
		Data: nil,
		Err:  errors.New("mock error"),
	}

	mc := controller.NewMemberController(mockService)

	mc.GetMemberById(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "mock error")
}

func TestGrantTutorRole_Success(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "username", Value: "john"}}

	mockService := &service.MockMemberService{
		Data: "Role granted successfully",
		Err:  nil,
	}

	mc := controller.NewMemberController(mockService)

	mc.GrantTutorRole(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Role granted successfully", strings.Trim(w.Body.String(), `"`))
}

func TestGrantTutorRole_Error(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "username", Value: "john"}}

	mockService := &service.MockMemberService{
		Data: "",
		Err:  errors.New("mock error"),
	}

	mc := controller.NewMemberController(mockService)

	mc.GrantTutorRole(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "mock error")
}
