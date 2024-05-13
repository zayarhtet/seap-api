package service

import "github.com/zayarhtet/seap-api/src/server/model/dao"

// MockCredentialRepository is a mock implementation of CredentialRepository
type MockCredentialRepository struct {
	CredentialId string
	Err          error
}

func (m *MockCredentialRepository) GetCredentialById(credential *dao.Credential) error {
	//TODO implement me
	panic("implement me")
}

// SaveCredential mocks the SaveCredential method of CredentialRepository
func (m *MockCredentialRepository) SaveCredential(credential *dao.Credential) (string, error) {
	if m.Err != nil {
		return "", m.Err
	}
	return m.CredentialId, nil
}

// DeleteCredential mocks the DeleteCredential method of CredentialRepository
func (m *MockCredentialRepository) DeleteCredential(credentialId string) error {
	// Implement mock behavior as needed
	return nil
}
