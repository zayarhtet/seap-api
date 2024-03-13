package repository

import (
	"github.com/zayarhtet/seap-api/src/server/model/dao"
	"github.com/zayarhtet/seap-api/src/server/model/dto"
)

type FamilyRepository interface {
	GetAllFamilies(int, int) *[]dao.Family
	GetAllFamiliesWithMembers(int, int) *[]dao.FamilyWithMembers
	GetMemberByIdWithFamilies(*dao.MemberWithFamilies) error
	GetFamilyById(*dao.FamilyWithMembers) error
	SaveNewFamily(*dao.Family) error
	SaveNewMember(*dto.MemberToFamilyRequest) error
	GetMemberRoleInFamily(*dto.MemberToFamilyRequest) error
	GetRowCount() *int64
}

type FamilyRepositoryImpl struct{}

func (fr FamilyRepositoryImpl) GetAllFamilies(offset, limit int) *[]dao.Family {
	var families []dao.Family
	dc.getAllByPagination(&families, offset, limit, &dao.Family{})
	return &families
}

func (fr FamilyRepositoryImpl) GetAllFamiliesWithMembers(offset, limit int) *[]dao.FamilyWithMembers {
	var families []dao.FamilyWithMembers
	dc.getAllByPagination(&families, offset, limit, &dao.Family{}, "Members.User", "Members", "Members.MemberRole")
	return &families
}

func (fr FamilyRepositoryImpl) GetMemberByIdWithFamilies(member *dao.MemberWithFamilies) error {
	return dc.getById(member, &dao.MemberWithFamilies{}, "Families.Family", "Families.MemberRole", "Families").Error
}

func (fr FamilyRepositoryImpl) SaveNewFamily(family *dao.Family) error {
	return dc.insertOne(family).Error
}

func (fr FamilyRepositoryImpl) SaveNewMember(family *dto.MemberToFamilyRequest) error {
	return dc.insertOne(family).Error
}

func (fr FamilyRepositoryImpl) GetFamilyById(family *dao.FamilyWithMembers) error {
	return dc.getById(family, &dao.FamilyWithMembers{}, "Members.User", "Members.MemberRole", "Members").Error
}

func (fr FamilyRepositoryImpl) GetMemberRoleInFamily(rq *dto.MemberToFamilyRequest) error {
	return dc.getById(rq, &dto.MemberToFamilyRequest{}).Error
}

func (fr FamilyRepositoryImpl) GetRowCount() *int64 {
	var count int64
	dc.getRowCount("family", &count)
	return &count
}
