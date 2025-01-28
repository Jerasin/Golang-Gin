package mocks

import (
	"github.com/Jerasin/app/repository"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockBaseRepository is a mock for BaseRepository
type MockBaseRepository struct {
	mock.Mock
}

func (m *MockBaseRepository) ClientDb() *gorm.DB {
	args := m.Called()
	return args.Get(0).(*gorm.DB)
}

func (m *MockBaseRepository) Pagination(p repository.PaginationModel, query interface{}, args ...interface{}) (interface{}, error) {
	argsPagination := m.Called(p, query, args)
	return argsPagination.Get(0), argsPagination.Error(0)
}

func (m *MockBaseRepository) Save(tx *gorm.DB, model interface{}) error {
	args := m.Called(tx, model)
	return args.Error(0)
}

func (m *MockBaseRepository) IsExits(tx *gorm.DB, model interface{}, query interface{}, args ...interface{}) error {
	argsIsExits := m.Called(tx, model, query, args)
	return argsIsExits.Error(0)
}

func (m *MockBaseRepository) FindOne(tx *gorm.DB, model interface{}, query interface{}, args ...interface{}) error {
	argsIsExits := m.Called(tx, model, query, args)
	return argsIsExits.Error(0)
}

func (m *MockBaseRepository) Find(tx *gorm.DB, model interface{}, query interface{}, args ...interface{}) error {
	argsIsExits := m.Called(tx, model, query, args)
	return argsIsExits.Error(0)
}

func (m *MockBaseRepository) Update(tx *gorm.DB, id int, model interface{}, update interface{}) error {
	argsIsExits := m.Called(tx, id, model, update)
	return argsIsExits.Error(0)
}

func (m *MockBaseRepository) TotalPage(model interface{}, pageSize int) (int64, error) {
	argsIsExits := m.Called(model, pageSize)
	return argsIsExits.Get(0).(int64), argsIsExits.Error(1)
}

func (m *MockBaseRepository) Delete(model interface{}, id int) error {
	argsIsExits := m.Called(model, id)
	return argsIsExits.Error(0)
}

func (m *MockBaseRepository) FindOneV2(tx *gorm.DB, model interface{}, options repository.Options) error {
	argsIsExits := m.Called(tx, model, options)
	return argsIsExits.Error(0)
}
