package controller_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/zayarhtet/seap-api/src/server/controller"
	service "github.com/zayarhtet/seap-api/test/controller"
)

func TestAddNewGrade_Success(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	payload := `{"username": "john", "gradingId": "1", "points": 100, "gradeComment": "Excellent work"}`

	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	mockService := &service.MockDutyService{
		Data: "Grade added successfully",
		Err:  nil,
	}

	dc := controller.NewDutyController(mockService, nil)

	dc.AddNewGrade(c)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, "Grade added successfully", strings.Trim(w.Body.String(), "\""))
}

func TestAddNewGrade_InvalidInput(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	payload := `{}`

	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	mockService := &service.MockDutyService{}

	dc := controller.NewDutyController(mockService, nil)

	dc.AddNewGrade(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid Input")
}

func TestAddNewGrade_Error(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	payload := `{"username": "john", "gradingId": "1", "points": 100, "gradeComment": "Excellent work"}`

	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	mockService := &service.MockDutyService{
		Data: "",
		Err:  errors.New("mock error"),
	}

	dc := controller.NewDutyController(mockService, nil)

	dc.AddNewGrade(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "mock error")
}

func TestSaveNewDuty_Success(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	payload := `{
        "dutyId": "1",
        "title": "Sample Duty",
        "instruction": "Complete the task",
        "publishedAt": "2024-05-25T00:00",
        "dueDate": "2024-06-25T00:00",
        "closingDate": "2024-06-20T00:00",
        "familyId": "2",
        "isPointSystem": true,
        "totalPoints": 100,
        "multipleSubmission": false,
        "pluginName": "plugin",
        "files": []
    }`

	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	mockService := &service.MockDutyService{
		Data: "Duty saved successfully",
		Err:  nil,
	}

	dc := controller.NewDutyController(mockService, nil)

	dc.SaveNewDuty(c)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, `"Duty saved successfully"`, w.Body.String())
}

func TestSaveNewDuty_InvalidInput(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	payload := `{}`

	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	mockService := &service.MockDutyService{}

	dc := controller.NewDutyController(mockService, nil)

	dc.SaveNewDuty(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid Input")
}

func TestSaveNewDuty_Error(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	payload := `{
        "dutyId": "1",
        "title": "Sample Duty",
        "instruction": "Complete the task",
        "publishedAt": "2024-05-25T00:00",
        "dueDate": "2024-06-25T00:00",
        "closingDate": "2024-06-20T00:00",
        "familyId": "2",
        "isPointSystem": true,
        "totalPoints": 100,
        "multipleSubmission": false,
        "pluginName": "plugin",
        "files": []
    }`

	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	mockService := &service.MockDutyService{
		Data: nil,
		Err:  errors.New("mock error"),
	}

	dc := controller.NewDutyController(mockService, nil)

	dc.SaveNewDuty(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "mock error")
}

func TestSubmitDuty_Success(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = append(c.Params, gin.Param{Key: "gradingId", Value: "1"}, gin.Param{Key: "dutyId", Value: "2"})

	mockService := &service.MockDutyService{
		Err: nil,
	}

	dc := controller.NewDutyController(mockService, nil)

	dc.SubmitDuty(c)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, `{"message":"success"}`, w.Body.String())
}

func TestSubmitDuty_Error(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = append(c.Params, gin.Param{Key: "gradingId", Value: "1"}, gin.Param{Key: "dutyId", Value: "2"})

	mockService := &service.MockDutyService{
		Err: errors.New("mock error"),
	}

	dc := controller.NewDutyController(mockService, nil)

	dc.SubmitDuty(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "mock error")
}

func TestGetMyGrading_Success(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = append(c.Params, gin.Param{Key: "dutyId", Value: "1"})
	c.Set("username", "testuser")

	mockService := &service.MockDutyService{
		Data: "My grading data",
		Err:  nil,
	}

	dc := controller.NewDutyController(mockService, nil)

	dc.GetMyGrading(c)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, `"My grading data"`, w.Body.String())
}

func TestGetMyGrading_Error(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = append(c.Params, gin.Param{Key: "dutyId", Value: "1"})
	c.Set("username", "testuser")

	mockService := &service.MockDutyService{
		Data: "",
		Err:  errors.New("mock error"),
	}

	dc := controller.NewDutyController(mockService, nil)

	dc.GetMyGrading(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "mock error")
}

func TestTriggerPluginExecution_Success(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = append(c.Params, gin.Param{Key: "dutyId", Value: "1"})

	mockService := &service.MockEngineService{
		Data: "Plugin execution result",
		Err:  nil,
	}

	dc := controller.NewDutyController(nil, mockService)

	dc.TriggerPluginExecution(c)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, `"Plugin execution result"`, w.Body.String())
}

func TestTriggerPluginExecution_Error(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = append(c.Params, gin.Param{Key: "dutyId", Value: "1"})

	mockService := &service.MockEngineService{
		Data: "",
		Err:  errors.New("mock error"),
	}

	dc := controller.NewDutyController(nil, mockService)

	dc.TriggerPluginExecution(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "mock error")
}
