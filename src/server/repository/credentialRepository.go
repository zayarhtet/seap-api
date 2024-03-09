package repository

import (
	"github.com/zayarhtet/seap-api/src/server/model/dao"
)

type CredentialRepository interface {
	SaveCredential(*dao.Credential) (string, error)
	GetCredentialById(*dao.Credential) error
	DeleteCredential(string) error
}

type CredentialRepositoryImpl struct{}

func (cr CredentialRepositoryImpl) SaveCredential(credential *dao.Credential) (string, error) {
	dr := dc.insertOne(credential)
	return credential.CredentialId, dr.Error
}

func (cr CredentialRepositoryImpl) DeleteCredential(id string) error {
	return dc.deleteOneById(&dao.Credential{CredentialId: id}).Error
}

func (cr CredentialRepositoryImpl) GetCredentialById(credential *dao.Credential) error {
	return dc.getById(credential, &dao.Credential{}).Error
}
