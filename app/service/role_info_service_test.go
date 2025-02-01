package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jerasin/app/mocks"
	"github.com/Jerasin/app/model"
	"github.com/Jerasin/app/repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
		name       string
		fields     fields
		args       args
		wantErr    bool
		statusCode int
	}{
		// TODO: Add test cases.
		{
			name: "Success case with valid JSON",
			fields: fields{
				BaseRepository: func() repository.BaseRepositoryInterface {
					mockRepo := new(mocks.MockBaseRepository)
					mockDb, _ := mocks.NewMockDB()
					mockRepo.On("ClientDb").Return(mockDb)
					mockRepo.On("Transaction").Return(nil)
					mockRepo.On("Save", mock.Anything, mock.Anything).Return(nil)
					return mockRepo
				}(),
			},
			args: args{
				c: func() *gin.Context {
					gin.SetMode(gin.ReleaseMode)
					w := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(w)

					roleInfo := model.RoleInfo{
						Name: "Admin",
					}

					body, _ := json.Marshal(roleInfo)
					req, _ := http.NewRequest(http.MethodPost, "/createRoleInfo", bytes.NewReader(body))
					req.Header.Set("Content-Type", "application/json")
					c.Request = req

					return c
				}(),
			},
			wantErr:    false,
			statusCode: 200,
		},
		{
			name: "Error case with invalid JSON format",
			fields: fields{
				BaseRepository: func() repository.BaseRepositoryInterface {
					mockRepo := new(mocks.MockBaseRepository)
					mockDb, _ := mocks.NewMockDB()
					mockRepo.On("ClientDb").Return(mockDb)
					return mockRepo
				}(),
			},
			args: args{
				c: func() *gin.Context {
					gin.SetMode(gin.ReleaseMode)
					w := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(w)
					invalidBody := []byte(`{"Name":"4"`)
					req, _ := http.NewRequest(http.MethodPost, "/createRoleInfo", bytes.NewReader(invalidBody))
					req.Header.Set("Content-Type", "application/json")
					c.Request = req

					return c
				}(),
			},
			wantErr:    true,
			statusCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := RoleInfoServiceModel{
				BaseRepository: tt.fields.BaseRepository,
			}
			p.CreateRoleInfo(tt.args.c)

			if tt.wantErr {
				if tt.statusCode == tt.args.c.Writer.Status() {
					fmt.Println("dd", tt.args.c.Writer.Status())
					assert.NotNil(t, tt.args.c)
					assert.Equal(t, tt.statusCode, tt.args.c.Writer.Status())
				} else {
					panic("error")
				}

			} else {
				assert.Equal(t, http.StatusOK, tt.args.c.Writer.Status())
			}
		})
	}
}
