package repository

import "github.com/zayarhtet/seap-api/src/server/model/dao"

type FamilyRepository interface {
	GetAllFamilies(int, int) *[]dao.Family
	GetAllFamiliesWithMembers(int, int) *[]dao.FamilyWithMembers
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

func (fr FamilyRepositoryImpl) GetRowCount() *int64 {
	var count int64
	dc.getRowCount("family", &count)
	return &count
}
