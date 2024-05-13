package service

import "github.com/zayarhtet/seap-api/src/server/model/dao"

type MockMemberRepository struct {
	Member *dao.Member
	Err    error
}

func (m *MockMemberRepository) GetAllMembers(i int, i2 int) *[]dao.Member {
	//TODO implement me
	panic("implement me")
}

func (m *MockMemberRepository) GetAllMembersWithFamilies(i int, i2 int) *[]dao.MemberWithFamilies {
	//TODO implement me
	panic("implement me")
}

func (m *MockMemberRepository) GetRowCount() *int64 {
	//TODO implement me
	panic("implement me")
}

func (m *MockMemberRepository) SaveMember(member *dao.Member) (*dao.Member, error) {
	//TODO implement me
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Member, nil
}

func (m *MockMemberRepository) DeleteMember(member *dao.Member) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockMemberRepository) UpdateMember(m2 map[string]any, member *dao.Member) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockMemberRepository) GetMemberByUsername(member *dao.Member) error {
	if m.Err != nil {
		return m.Err
	}
	member.Username = m.Member.Username
	return nil
}
