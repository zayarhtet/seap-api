package controller_test

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/zayarhtet/seap-api/src/server/controller"
	service "github.com/zayarhtet/seap-api/test/controller"
)

func TestGetProfileImage_Success(t *testing.T) {
	file, err := os.CreateTemp("", "mock-image-*.jpg")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(file.Name())

	mockService := &service.MockFamilyService{
		FilePath: file.Name(),
		Err:      nil,
	}

	cc := controller.NewCDNController(mockService, nil)

	router := gin.Default()
	w := httptest.NewRecorder()

	router.GET("/your-endpoint/:famId", cc.GetProfileImage)

	req, err := http.NewRequest("GET", "/your-endpoint/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "image/jpeg", w.Header().Get("Content-Type"))
}

func TestGetProfileImage_Error(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = append(c.Params, gin.Param{Key: "famId", Value: "1"})

	mockService := &service.MockFamilyService{
		FilePath: "",
		Err:      errors.New("mock error"),
	}

	cc := controller.NewCDNController(mockService, nil)

	cc.GetProfileImage(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "mock error")
}

func TestSaveGivenFiles_Success(t *testing.T) {
	mockService := &service.MockDutyService{
		InputFileErr: nil,
		GivenFileErr: nil,
	}

	cc := controller.NewCDNController(nil, mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "dutyId", Value: "1"})

	tempFile, err := os.CreateTemp("", "example.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	file, err := os.Open(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("files", "example.txt")
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatalf("Failed to copy file content: %v", err)
	}
	writer.Close()

	req, err := http.NewRequest("POST", "/", body)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	c.Request = req

	cc.SaveGivenFiles(c)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.JSONEq(t, `{"message": "uploaded."}`, w.Body.String())
}

func TestSaveGivenFiles_CreateGivenFilesError(t *testing.T) {
	mockService := &service.MockDutyService{
		GivenFileErr: errors.New("failed to create given files"),
	}

	cc := controller.NewCDNController(nil, mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "dutyId", Value: "1"})

	tempFile, err := os.CreateTemp("", "example.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	file, err := os.Open(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("files", "example.txt")
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatalf("Failed to copy file content: %v", err)
	}
	writer.Close()

	req, err := http.NewRequest("POST", "/", body)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	c.Request = req

	cc.SaveGivenFiles(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	assert.JSONEq(t, `{"error": "Failed to create assignment and upload files"}`, w.Body.String())
}

func TestSaveGivenFiles_SaveInputFilesError(t *testing.T) {
	mockService := &service.MockDutyService{
		InputFileErr: errors.New("failed to save input files"),
	}

	cc := controller.NewCDNController(nil, mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "dutyId", Value: "1"})

	tempFile, err := os.CreateTemp("", "example.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	file, err := os.Open(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("files", "example.txt")
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatalf("Failed to copy file content: %v", err)
	}
	writer.Close()

	req, err := http.NewRequest("POST", "/", body)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	c.Request = req

	cc.SaveGivenFiles(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	assert.JSONEq(t, `{"error": "Failed to create assignment and upload files"}`, w.Body.String())
}
