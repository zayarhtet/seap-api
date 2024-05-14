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

	fs := service.NewFamilyServiceForTest(mockFamilyRepo, nil)

	familyId := "mockFamilyId"
	username := "mockUsername"
	response, err := fs.GetFamilyByIdWithDuties(familyId, username)

	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestRoleInFamily_Success(t *testing.T) {
	mockFamilyRepo := &repository.MockFamilyRepository{
		RoleDB: "tutor",
		Err:    nil,
	}

	fs := service.NewFamilyServiceForTest(mockFamilyRepo, nil)

	username := "mockUsername"
	familyId := "mockFamilyId"
	role, err := fs.RoleInFamily(username, familyId)

	assert.Nil(t, err)
	assert.NotNil(t, role)
}

func TestRoleInFamily_Error(t *testing.T) {
	mockFamilyRepo := &repository.MockFamilyRepository{
		Err: errors.New("error message"),
	}

	fs := service.NewFamilyServiceForTest(mockFamilyRepo, nil)

	username := "mockUsername"
	familyId := "mockFamilyId"
	role, err := fs.RoleInFamily(username, familyId)

	assert.NotNil(t, err)
	assert.Empty(t, role)
}
