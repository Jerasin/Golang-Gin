package mocks

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Jerasin/app/repository"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening gorm database", err)
	}

	return gormDB, mock
}

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

func (m *MockBaseRepository) Find(tx *gorm.DB, model interface{}, query interface{}, p repository.PaginationModel, args ...interface{}) error {
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

// ใน mock repository (MockBaseRepository)
func (m *MockBaseRepository) Transaction(fn func(tx *gorm.DB) error) error {
	// ใช้ `Called` เพื่อรับการเรียกของฟังก์ชัน
	args := m.Called(fn)

	// รัน callback (ฟังก์ชันที่รับ tx) ด้วย mock db หรือ nil
	if fn != nil {
		tx := new(gorm.DB) // หรือ mock DB ที่ต้องการ
		return fn(tx)      // รัน callback
	}

	return args.Error(0)
}
