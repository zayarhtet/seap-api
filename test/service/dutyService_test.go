package service_test

import (
	"mime/multipart"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/zayarhtet/seap-api/src/server/model/dao"
	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"github.com/zayarhtet/seap-api/src/server/service"
	repository "github.com/zayarhtet/seap-api/test/service"
)

func TestGetAllDutiesResponse(t *testing.T) {
	mockDuties := []dao.Duty{
		{DutyId: "hello1", Title: "Duty 1"},
		{DutyId: "hello2", Title: "Duty 2"},
	}

	mockRepo := &repository.MockDutyRepository{Duties: mockDuties}

	ds := service.NewDutyServiceForTest(mockRepo, nil)

	size := 10
	page := 1

	response, err := ds.GetAllDutiesResponse(size, page)

	assert.NotNil(t, response)
	assert.NoError(t, err)

	resp := *(response.(**dto.DataResponse))
	assert.Equal(t, len(mockDuties), len(*(resp.Data.(*[]dao.Duty))))

	for i, duty := range *(resp.Data.(*[]dao.Duty)) {
		assert.Equal(t, mockDuties[i].DutyId, duty.DutyId)
		assert.Equal(t, mockDuties[i].Title, duty.Title)
	}
}

func TestGetAllDutiesByMemberResponse(t *testing.T) {
	mockDuties := []dao.MyDuty{
		{DutyId: "1", Duty_: dao.DutiesForFamily{PublishingDate: time.Now().Add(-24 * time.Hour)}},
		{DutyId: "2", Duty_: dao.DutiesForFamily{PublishingDate: time.Now().Add(24 * time.Hour)}},
	}

	mockRepo := &repository.MockDutyRepository{MyDuties: mockDuties}

	ds := service.NewDutyServiceForTest(mockRepo, nil)

	username := "testuser"

	response, err := ds.GetAllDutiesByMemberResponse(username)

	assert.NotNil(t, response)
	assert.NoError(t, err)

	resp := *(response.(**dto.DataResponse))
	assert.Equal(t, len(mockDuties)-1, len(*(resp.Data.(*[]dao.MyDuty)))) // One duty should be filtered out
}

func TestAddNewGradeResponse(t *testing.T) {
	mockRepo := &repository.MockDutyRepository{}

	ds := service.NewDutyServiceForTest(mockRepo, nil)

	request := dto.NewGradeRequest{
		Username:     "testuser",
		GradingId:    "123",
		Points:       10,
		GradeComment: "Good job!",
	}

	response, err := ds.AddNewGradeResponse(request)

	assert.NotNil(t, response)
	assert.NoError(t, err)
	assert.Equal(t, "HELLO SUCCESS", response)
}

func TestSaveNewDutyResponse(t *testing.T) {
	mockDutyRepo := &repository.MockDutyRepository{}

	mockFamilyRepo := &repository.MockFamilyRepository{}

	ds := service.NewDutyServiceForTest(mockDutyRepo, mockFamilyRepo)

	newDuty := dao.Duty{}

	response, err := ds.SaveNewDutyResponse(newDuty)

	assert.NotNil(t, response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response.(dao.Duty).DutyId)
}

func TestCreateGivenFiles(t *testing.T) {
	mockDutyRepo := &repository.MockDutyRepository{}

	ds := service.NewDutyServiceForTest(mockDutyRepo, nil)

	mockFileHeaders := []*multipart.FileHeader{}

	dutyID := "mockDutyID"

	err := ds.CreateGivenFiles(mockFileHeaders, dutyID)

	assert.NoError(t, err)
}

func TestGetGradingByDutyIdResponse(t *testing.T) {
	var mockTotal int64 = 10
	var mockOffset uint = 0

	mockGradingData := []dao.Grading{}

	mockDutyRepo := &repository.MockDutyRepository{}

	ds := service.NewDutyServiceForTest(mockDutyRepo, nil)

	response, err := ds.GetGradingByDutyIdResponse("mockDutyID", 10, 1)

	assert.NoError(t, err)

	assert.NotNil(t, response)

	resp := *(response.(**dto.DataResponse))
	assert.Equal(t, mockTotal, resp.Total)
	assert.Equal(t, mockOffset, resp.Size)
	assert.Equal(t, len(mockGradingData), len(*resp.Data.(*[]dao.Grading)))
}
