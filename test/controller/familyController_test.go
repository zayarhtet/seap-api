package controller_test

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/zayarhtet/seap-api/src/server/controller"
	service "github.com/zayarhtet/seap-api/test/controller"
)

func TestSaveNewFamily_Success(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("username", "john")

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	_ = writer.WriteField("familyName", "Doe")
	_ = writer.WriteField("familyInfo", "Family of John Doe")
	writer.Close()

	req, _ := http.NewRequest("POST", "/", &b)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	c.Request = req

	mockService := &service.MockFamilyService{
		Data: "New family saved successfully", // Dummy response data
		Err:  nil,                             // No error
	}

	fc := controller.NewFamilyController(mockService)

	fc.SaveNewFamily(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "New family saved successfully", strings.Trim(w.Body.String(), "\""))
}

func TestSaveNewFamily_InvalidInput(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("username", "john")

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	writer.Close()

	req, _ := http.NewRequest("POST", "/", &b)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	c.Request = req

	mockService := &service.MockFamilyService{}

	fc := controller.NewFamilyController(mockService)

	fc.SaveNewFamily(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid Input")
}

func TestSaveNewFamily_Error(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("username", "john")

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	_ = writer.WriteField("familyName", "Doe")
	_ = writer.WriteField("familyInfo", "Family of John Doe")
	writer.Close()

	req, _ := http.NewRequest("POST", "/", &b)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	c.Request = req

	mockService := &service.MockFamilyService{
		Data: "",                       // No data
		Err:  errors.New("mock error"), // Predefined error
	}

	fc := controller.NewFamilyController(mockService)

	fc.SaveNewFamily(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "mock error")
}

func TestAddNewMemberToFamily_Success(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "famId", Value: "1"}}

	payload := `{"username": "john", "roleId": 1}`

	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	mockService := &service.MockFamilyService{
		Data: "Member added to family successfully", // Dummy response data
		Err:  nil,                                   // No error
	}

	fc := controller.NewFamilyController(mockService)

	fc.AddNewMemberToFamily(c)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, "Member added to family successfully", strings.Trim(w.Body.String(), "\""))
}

func TestAddNewMemberToFamily_InvalidInput(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "famId", Value: "1"}}

	payload := `{}`

	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	mockService := &service.MockFamilyService{}

	fc := controller.NewFamilyController(mockService)

	fc.AddNewMemberToFamily(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid Input")
}

func TestAddNewMemberToFamily_Error(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "famId", Value: "1"}}

	payload := `{"username": "john", "roleId": 1}`

	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	mockService := &service.MockFamilyService{
		Data: "",                       // No data
		Err:  errors.New("mock error"), // Predefined error
	}

	fc := controller.NewFamilyController(mockService)

	fc.AddNewMemberToFamily(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "mock error")
}
