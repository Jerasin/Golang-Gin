package mocks

import (
	"github.com/Jerasin/app/dto"
	"github.com/Jerasin/app/response"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockRoleInfoService struct {
	mock.Mock
}

func (m *MockRoleInfoService) CreateRoleInfo(ctx *gin.Context, body dto.RoleInfoCreateRequest) {
	// args := m.Called(body)
	// return args.Get(0).(dto.RoleInfoResponse), args.Error(1)
}

func (m *MockRoleInfoService) GetPaginationRoleInfo(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.RoleInfo) {
	// args := m.Called(body)
	// return args.Get(0).(dto.RoleInfoResponse), args.Error(1)
}

func (m *MockRoleInfoService) GetRoleInfoById(ctx *gin.Context) {
	// args := m.Called(body)
	// return args.Get(0).(dto.RoleInfoResponse), args.Error(1)
}

func (m *MockRoleInfoService) UpdateRoleInfo(ctx *gin.Context, body dto.UpdateRoleInfoCreateRequest) {
	// args := m.Called(body)
	// return args.Get(0).(dto.RoleInfoResponse), args.Error(1)
}

func (m *MockRoleInfoService) DeleteRoleInfo(ctx *gin.Context) {
	// args := m.Called(body)
	// return args.Get(0).(dto.RoleInfoResponse), args.Error(1)
}
