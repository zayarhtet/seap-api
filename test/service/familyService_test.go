package service_test

import (
	"errors"
	"testing"

	"github.com/zayarhtet/seap-api/src/server/model/dao"
	"github.com/zayarhtet/seap-api/src/server/service"
	repository "github.com/zayarhtet/seap-api/test/service"

	// Import necessary packages for testing
	"github.com/stretchr/testify/assert"
)

func TestGetFamilyByIdWithDuties_Tutee_Success(t *testing.T) {
	mockFamilyRepo := &repository.MockFamilyRepository{
		FamilyWithDuties: &dao.FamilyWithDuties{},
		Role:             "tutee",
		Err:              nil,
	}

	fs := service.NewFamilyServiceForTest(mockFamilyRepo, nil)

	familyId := "mockFamilyId"
	username := "mockUsername"
	response, err := fs.GetFamilyByIdWithDuties(familyId, username)

	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestGetFamilyByIdWithDuties_Tutor_Success(t *testing.T) {
	mockFamilyRepo := &repository.MockFamilyRepository{
		FamilyWithDuties: &dao.FamilyWithDuties{},
		Role:             "tutor",
		Err:              nil,
	}

	// Initialize family service with mock repository and service
	fs := service.NewFamilyServiceForTest(mockFamilyRepo, nil)

	// Call the function being tested
	familyId := "mockFamilyId"
	username := "mockUsername"
	response, err := fs.GetFamilyByIdWithDuties(familyId, username)

	// Assertions
	assert.Nil(t, err)         // No error expected
	assert.NotNil(t, response) // Response should not be nil
	// Add more assertions to verify the response contents
}

func TestRoleInFamily_Success(t *testing.T) {
	// Initialize mock family repository
	mockFamilyRepo := &repository.MockFamilyRepository{
		RoleDB: "tutor",
		Err:    nil,
	}

	// Initialize family service with mock repository
	fs := service.NewFamilyServiceForTest(mockFamilyRepo, nil)

	// Call the function being tested
	username := "mockUsername"
	familyId := "mockFamilyId"
	role, err := fs.RoleInFamily(username, familyId)

	// Assertions
	assert.Nil(t, err)     // No error expected
	assert.NotNil(t, role) // Role should not be nil
	// Add more assertions to verify the role value
}

func TestRoleInFamily_Error(t *testing.T) {
	// Initialize mock family repository with error
	mockFamilyRepo := &repository.MockFamilyRepository{
		Err: errors.New("error message"),
	}

	// Initialize family service with mock repository
	fs := service.NewFamilyServiceForTest(mockFamilyRepo, nil)

	// Call the function being tested
	username := "mockUsername"
	familyId := "mockFamilyId"
	role, err := fs.RoleInFamily(username, familyId)

	// Assertions
	assert.NotNil(t, err) // Error expected
	assert.Empty(t, role) // Role should be empty
}
