package repository

import (
	"github.com/zayarhtet/seap-api/src/server/model/dao"
)

type MemberRepository interface {
	GetAllMembers(int, int) *[]dao.Member
	GetRowCount() *int64
	SaveMember(*dao.Member) (*dao.Member, error)
	GetMemberByUsername(*dao.Member) error
	DeleteMember(*dao.Member) (string, error)
}

type MemberRepositoryImpl struct{}

func (m MemberRepositoryImpl) GetAllMembers(offset, limit int) *[]dao.Member {
	var members []dao.Member
	dc.getAllByPagination(&members, offset, limit, &dao.Member{}, "Role")
	return &members
}

func (m MemberRepositoryImpl) GetRowCount() *int64 {
	var count int64
	dc.getRowCount("member", &count)
	return &count
}

func (m MemberRepositoryImpl) SaveMember(member *dao.Member) (*dao.Member, error) {
	err := dc.insertOne(member).Error

	if err != nil {
		return nil, err
	}
	member = &dao.Member{
		Username: member.Username,
	}
	err = m.GetMemberByUsername(member)
	return member, err
}

func (m MemberRepositoryImpl) GetMemberByUsername(member *dao.Member) error {
	return dc.getById(member, &dao.Member{}, "Role").Error
}

func (m MemberRepositoryImpl) DeleteMember(member *dao.Member) (string, error) {
	err := dc.getById(member, &dao.Member{}).Error
	if err != nil {
		return "", err
	}
	credentialId := member.CredentialId
	err = dc.deleteOneById(member).Error
	if err != nil {
		return "", err
	}
	return credentialId, nil
}
