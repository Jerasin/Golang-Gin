package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jerasin/app/dto"
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
		c    *gin.Context
		body dto.RoleInfoCreateRequest
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
			name: "Success Insert RoleInfo",
			fields: fields{
				BaseRepository: func() repository.BaseRepositoryInterface {
					mockRepo := new(mocks.MockBaseRepository)
					mockDb, sqlMock := mocks.NewMockDB()
					mockRepo.On("ClientDb").Return(mockDb)

					// Mock Start Transaction
					sqlMock.ExpectBegin()

					// Mock Find
					mockRepo.On("Find", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

					// Mock Save
					mockRepo.On("Save", mock.Anything, mock.Anything).Return(nil).Once()

					// Mock Commit
					sqlMock.ExpectCommit()

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
				body: dto.RoleInfoCreateRequest{
					Name: "Admin",
				},
			},
			wantErr:    false,
			statusCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := RoleInfoServiceModel{
				BaseRepository: tt.fields.BaseRepository,
			}
			p.CreateRoleInfo(tt.args.c, tt.args.body)

			if tt.wantErr {
				fmt.Printf("Writer Status = %+v", tt.args.c.Writer.Status())
				if tt.statusCode == tt.args.c.Writer.Status() {
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
