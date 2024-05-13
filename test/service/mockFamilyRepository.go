package service

import (
	"github.com/zayarhtet/seap-api/src/server/model/dao"
	"github.com/zayarhtet/seap-api/src/server/model/dto"
)

// MockFamilyRepository is a mock implementation of FamilyRepository
type MockFamilyRepository struct {
	RoleDB           string
	FamilyWithDuties *dao.FamilyWithDuties
	Role             string
	Err              error
}

func (m *MockFamilyRepository) GetAllFamilies(i int, i2 int) *[]dao.Family {
	//TODO implement me
	panic("implement me")
}

func (m *MockFamilyRepository) GetAllFamiliesWithMembers(i int, i2 int) *[]dao.FamilyWithMembers {
	//TODO implement me
	panic("implement me")
}

func (m *MockFamilyRepository) GetMemberByIdWithFamilies(families *dao.MemberWithFamilies) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockFamilyRepository) GetFamilyById(members *dao.FamilyWithMembers) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockFamilyRepository) GetFamilyOnlyById(family *dao.Family) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockFamilyRepository) SaveNewFamily(family *dao.Family) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockFamilyRepository) SaveNewMember(request *dto.MemberToFamilyRequest) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockFamilyRepository) GetMemberRoleInFamily(family *dao.MemberForFamily) error {
	if m.Err != nil {
		return m.Err
	}
	family.MemberRole.Name = m.RoleDB
	return nil
}

func (m *MockFamilyRepository) GetMyRoleInFamily(member *dao.FamilyForMember) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockFamilyRepository) GetRowCount() *int64 {
	//TODO implement me
	panic("implement me")
}

func (m *MockFamilyRepository) DeleteFamilyById(family *dao.Family) error {
	//TODO implement me
	panic("implement me")
}

// GetFamilyByIdWithDutiesForTutee mocks the GetFamilyByIdWithDutiesForTutee method of FamilyRepository
func (m *MockFamilyRepository) GetFamilyByIdWithDutiesForTutee(family *dao.FamilyWithDuties, username string) error {
	if m.Err != nil {
		return m.Err
	}
	*family = *m.FamilyWithDuties
	return nil
}

// GetFamilyByIdWithDutiesForTutor mocks the GetFamilyByIdWithDutiesForTutor method of FamilyRepository
func (m *MockFamilyRepository) GetFamilyByIdWithDutiesForTutor(family *dao.FamilyWithDuties) error {
	if m.Err != nil {
		return m.Err
	}
	*family = *m.FamilyWithDuties
	return nil
}

// MockFamilyService is a mock implementation of FamilyService
type MockFamilyService struct {
	Role string
	Err  error
}

// RoleInFamily mocks the RoleInFamily method of FamilyService
func (m *MockFamilyService) RoleInFamily(username, familyId string) (string, error) {
	if m.Err != nil {
		return "", m.Err
	}
	return m.Role, nil
}
