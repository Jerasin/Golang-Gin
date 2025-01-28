package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jerasin/app/mocks"
	"github.com/Jerasin/app/model"
	"github.com/Jerasin/app/repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func TestRoleInfoServiceModel_CreateRoleInfo(t *testing.T) {
	type fields struct {
		BaseRepository repository.BaseRepositoryInterface
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Success case with valid JSON",
			fields: fields{
				BaseRepository: func() repository.BaseRepositoryInterface {
					mockRepo := new(mocks.MockBaseRepository)
					mockRepo.On("ClientDb").Return(nil)                           // Mock ClientDb
					mockRepo.On("Save", mock.Anything, mock.Anything).Return(nil) // Mock Save success
					return mockRepo
				}(),
			},
			args: args{
				c: func() *gin.Context {
					// สร้าง Gin context สำหรับการทดสอบ request
					w := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(w)

					roleInfo := model.RoleInfo{
						Name: "Admin",
					}

					// Marshal body
					body, _ := json.Marshal(roleInfo)
					req, _ := http.NewRequest(http.MethodPost, "/createRoleInfo", bytes.NewReader(body))
					req.Header.Set("Content-Type", "application/json")
					c.Request = req

					return c
				}(),
			},
			wantErr: false, // หวังว่าจะไม่เกิด error
		},
		{
			name: "Error case with invalid JSON",
			fields: fields{
				BaseRepository: func() repository.BaseRepositoryInterface {
					// Mock repository with no operations
					mockRepo := new(mocks.MockBaseRepository)
					mockRepo.On("ClientDb").Return(nil)
					return mockRepo
				}(),
			},
			args: args{
				c: func() *gin.Context {
					// สร้าง Gin context สำหรับ request ที่มี JSON ผิด
					w := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(w)

					// ส่งข้อมูล JSON ที่ไม่ถูกต้อง
					invalidBody := []byte(`{"roleName":""}`) // ตัวอย่างที่ไม่ถูกต้อง
					req, _ := http.NewRequest(http.MethodPost, "/createRoleInfo", bytes.NewReader(invalidBody))
					req.Header.Set("Content-Type", "application/json")
					c.Request = req

					return c
				}(),
			},
			wantErr: true, // หวังว่าจะเกิด error (BadRequest)
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := RoleInfoServiceModel{
				BaseRepository: tt.fields.BaseRepository,
			}
			p.CreateRoleInfo(tt.args.c)
		})
	}
}
