package repository

import "gorm.io/gorm"

// MockDataCenter is a mock implementation of the DataCenter interface
type MockDataCenter struct{}

func (m *MockDataCenter) ConnectDatabase() {
	//TODO implement me
	panic("implement me")
}

func (m *MockDataCenter) GetAll(a any) *gorm.DB {
	//TODO implement me
	panic("implement me")
}

func (m *MockDataCenter) GetAllByPagination(a any, i int, i2 int, a2 any, s ...string) *gorm.DB {
	return &gorm.DB{Error: nil}
}

func (m *MockDataCenter) GetRowCount(s string, i *int64) *gorm.DB {
	//TODO implement me
	panic("implement me")
}

func (m *MockDataCenter) GetById(a any, a2 any, s ...string) *gorm.DB {
	//TODO implement me
	return &gorm.DB{Error: nil}
}

func (m *MockDataCenter) InsertOne(a any) *gorm.DB {
	//TODO implement me
	return &gorm.DB{Error: nil}
}

func (m *MockDataCenter) InsertAll(a any) *gorm.DB {
	//TODO implement me
	panic("implement me")
}

func (m *MockDataCenter) DeleteOneById(a any) *gorm.DB {
	//TODO implement me
	panic("implement me")
}

func (m *MockDataCenter) GetAllByPaginationWithCondition(a any, i int, i2 int, a2 any, a3 any, s ...string) *gorm.DB {
	return &gorm.DB{Error: nil}
}

func (m *MockDataCenter) GetByIdWithCondition(a any, s string, a2 any, s2 ...string) *gorm.DB {
	return &gorm.DB{Error: nil}
}

func (m *MockDataCenter) GetOneByStructCondition(a any, a2 any) *gorm.DB {
	//TODO implement me
	panic("implement me")
}

func (m *MockDataCenter) GetAllByStructCondition(a any, a2 any, a3 any) *gorm.DB {
	//TODO implement me
	panic("implement me")
}

func (m *MockDataCenter) UpdateModelByMap(m2 map[string]any, a any) *gorm.DB {
	//TODO implement me
	panic("implement me")
}
