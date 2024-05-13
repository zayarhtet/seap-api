package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zayarhtet/seap-api/src/server/model/dao"
	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"github.com/zayarhtet/seap-api/src/server/service"

	repository "github.com/zayarhtet/seap-api/test/service"
)

func TestGetMemberByIdResponse_Success(t *testing.T) {
	mockRepo := &repository.MockMemberRepository{
		Member: &dao.Member{
			Username: "mockUsername",
		},
		Err: nil,
	}

	ms := service.NewMemberServiceForTest(mockRepo, nil)

	response, err := ms.GetMemberByIdResponse("mockId")

	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestGetMemberByIdResponse_Error(t *testing.T) {
	mockRepo := &repository.MockMemberRepository{
		Member: nil,
		Err:    errors.New("mock error"),
	}

	ms := service.NewMemberServiceForTest(mockRepo, nil)

	_, err := ms.GetMemberByIdResponse("mockId")

	assert.NotNil(t, err)
	assert.Equal(t, "mock error", err.Error())
}

func TestSignUp_Success(t *testing.T) {
	mockCredentialRepo := &repository.MockCredentialRepository{
		CredentialId: "mockCredentialId",
		Err:          nil,
	}
	mockMemberRepo := &repository.MockMemberRepository{
		Member: &dao.Member{
			Username: "mockUsername",
		},
		Err: nil,
	}

	ms := service.NewMemberServiceForTest(mockMemberRepo, mockCredentialRepo)

	request := dto.SignUpRequest{}
	response, err := ms.SignUp(request)

	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestSignUp_Error(t *testing.T) {
	mockCredentialRepo := &repository.MockCredentialRepository{
		CredentialId: "",
		Err:          errors.New("mock credential error"),
	}
	mockMemberRepo := &repository.MockMemberRepository{
		Member: nil,
		Err:    errors.New("mock member error"),
	}

	ms := service.NewMemberServiceForTest(mockMemberRepo, mockCredentialRepo)

	request := dto.SignUpRequest{}
	_, err := ms.SignUp(request)

	assert.NotNil(t, err)
	assert.Equal(t, "mock credential error", err.Error())
}
