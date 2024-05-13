package repository_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zayarhtet/seap-api/src/server/model/dao"
	"github.com/zayarhtet/seap-api/src/server/repository"
	database "github.com/zayarhtet/seap-api/test/repository"
)

func TestGetAllMembers(t *testing.T) {
	mockDataCenter := &database.MockDataCenter{}
	repository.InitializeDataCenter(mockDataCenter)

	memberRepo := repository.MemberRepositoryImpl{}

	offset := 0
	limit := 10

	result := memberRepo.GetAllMembers(offset, limit)

	assert.NotNil(t, result)
	assert.Equal(t, reflect.TypeOf(result).String(), "*[]dao.Member")
}

func TestGetDutiesByUsername(t *testing.T) {
	mockDataCenter := &database.MockDataCenter{}

	repository.InitializeDataCenter(mockDataCenter)

	dutyRepo := repository.DutyRepositoryImpl{}

	condition := &dao.MyDuty{Username: "testuser"}

	result := dutyRepo.GetDutiesByUsername(condition)

	assert.NotNil(t, result)
	assert.Equal(t, reflect.TypeOf(result).String(), "*[]dao.MyDuty")
}

func TestGetMemberWithDutiesByUsername(t *testing.T) {
	mockDataCenter := &database.MockDataCenter{}

	repository.InitializeDataCenter(mockDataCenter)

	dutyRepo := repository.DutyRepositoryImpl{}

	mockMember := &dao.MemberWithDuties{}

	err := dutyRepo.GetMemberWithDutiesByUsername(mockMember)

	assert.NoError(t, err)
}
