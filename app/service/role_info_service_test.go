package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jerasin/app/dto"
	"github.com/Jerasin/app/mocks"
	"github.com/Jerasin/app/model"
	"github.com/Jerasin/app/repository"
	"github.com/Jerasin/app/response"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRoleInfoServiceModel_CreateRoleInfo(t *testing.T) {
	type fields struct {
		BaseRepository repository.BaseRepositoryInterface
	}
	type args struct {
		body dto.RoleInfoCreateRequest
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantErr          bool
		statusCode       int
		expectedResponse interface{}
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
				body: dto.RoleInfoCreateRequest{
					Name: "Admin",
				},
			},
			wantErr:    false,
			statusCode: 200,
			expectedResponse: map[string]interface{}{"response_key": "SUCCESS", "data": map[string]interface{}{
				"message": "create success",
			}, "response_message": "Success"},
		},
		{
			name: "Error Insert RoleInfo",
			fields: fields{
				BaseRepository: func() repository.BaseRepositoryInterface {
					mockRepo := new(mocks.MockBaseRepository)
					mockDb, sqlMock := mocks.NewMockDB()
					mockRepo.On("ClientDb").Return(mockDb)

					// Mock Start Transaction
					sqlMock.ExpectBegin()

					// Mock Find
					mockRepo.On("Find", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
						// Set the mock data to return when Find is called
						mockModel := args.Get(1).(*model.RoleInfo) // รับค่าตัวแปร model ที่ส่งไป
						// Fill model with mock data
						*mockModel = model.RoleInfo{
							BaseModel: model.BaseModel{
								ID: 1,
							},
							Name: "Admin",
						}
					}).Return(nil).Once()
					// Mock Save
					mockRepo.On("Save", mock.Anything, mock.Anything).Return(nil).Once()

					// Mock Commit
					sqlMock.ExpectCommit()

					return mockRepo
				}(),
			},
			args: args{
				body: dto.RoleInfoCreateRequest{
					Name: "Admin",
				},
			},
			wantErr:          true,
			statusCode:       400,
			expectedResponse: map[string]interface{}{"response_key": "DATA_IS_EXIT", "data": nil, "response_message": "DataIsExit"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.ReleaseMode)
			p := RoleInfoServiceModel{
				BaseRepository: tt.fields.BaseRepository,
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			p.CreateRoleInfo(c, tt.args.body)
			var responseBody map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &responseBody)

			if err != nil {
				panic(err)
			}

			fmt.Printf("Response Body = %+v", responseBody)

			if tt.wantErr {
				fmt.Printf("Writer Status = %+v", c.Writer.Status())
				if tt.statusCode == c.Writer.Status() {
					assert.NotNil(t, c)
					assert.Equal(t, tt.statusCode, c.Writer.Status())
				} else {
					panic("error")
				}

			} else {

				assert.Equal(t, http.StatusOK, c.Writer.Status())
			}

			assert.Equal(t, tt.expectedResponse, responseBody)
		})
	}
}

func TestRoleInfoServiceModel_GetPaginationRoleInfo(t *testing.T) {
	type fields struct {
		BaseRepository repository.BaseRepositoryInterface
	}
	type args struct {
		page      int
		pageSize  int
		search    string
		sortField string
		sortValue string
		field     response.RoleInfo
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantErr          bool
		statusCode       int
		expectedResponse interface{}
	}{
		// TODO: Add test cases.
		{
			name: "Success GetPaginationRoleInfo",
			fields: fields{
				BaseRepository: func() repository.BaseRepositoryInterface {
					mockRepo := new(mocks.MockBaseRepository)
					mockDest := []model.RoleInfo{
						// {BaseModel: model.BaseModel{ID: 1}, Name: "Admin"},
						// {BaseModel: model.BaseModel{ID: 2}, Name: "User"},
					}
					mockData := []model.RoleInfo{
						{BaseModel: model.BaseModel{ID: 1}, Name: "Admin"},
						{BaseModel: model.BaseModel{ID: 2}, Name: "User"},
					}
					// mockData := []model.RoleInfo(nil)
					var e []interface{}

					paginationModel := repository.PaginationModel{
						Limit:     10,
						Offset:    0,
						Search:    "Admin",
						SortField: "name",
						SortValue: "asc",
						Field: map[string]interface{}{
							"description": "",
							"name":        "",
							"id":          "",
						},
						Dest: mockDest,
					}

					mockRepo.On("Pagination", paginationModel, nil, e).Return(mockData, nil)

					mockRepo.On("TotalPage", &mockDest, 10).Return(int64(12), nil)

					return mockRepo
				}(),
			},
			args: args{
				page:      1,
				pageSize:  10,
				search:    "Admin",
				sortField: "name",
				sortValue: "asc",
				field:     response.RoleInfo{},
			},
			wantErr:    false,
			statusCode: 200,
			expectedResponse: map[string]interface{}{"response_key": "SUCCESS", "data": []interface{}{
				map[string]interface{}{"id": float64(1), "name": "Admin", "description": ""},
				map[string]interface{}{"id": float64(2), "name": "User", "description": ""},
			}, "response_message": "Success", "page": float64(1), "pageSize": float64(10), "totalPage": float64(12)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := RoleInfoServiceModel{
				BaseRepository: tt.fields.BaseRepository,
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			// c.Request = tt.args.c.Request

			p.GetPaginationRoleInfo(c, tt.args.page, tt.args.pageSize, tt.args.search, tt.args.sortField, tt.args.sortValue, tt.args.field)

			var responseBody map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &responseBody)

			if err != nil {
				panic(err)
			}

			fmt.Printf("Response Body = %+v", responseBody)

			if tt.wantErr {
				fmt.Printf("Writer Status = %+v", c.Writer.Status())
				if tt.statusCode == c.Writer.Status() {
					assert.NotNil(t, c)
					assert.Equal(t, tt.statusCode, c.Writer.Status())
				} else {
					panic("error")
				}

			} else {

				assert.Equal(t, http.StatusOK, c.Writer.Status())
			}

			assert.Equal(t, tt.expectedResponse, responseBody)
		})
	}
}
