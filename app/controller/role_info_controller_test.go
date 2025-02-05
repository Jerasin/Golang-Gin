package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jerasin/app/dto"
	"github.com/Jerasin/app/mocks"
	"github.com/Jerasin/app/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRoleInfoController_CreateRoleInfo(t *testing.T) {
	type fields struct {
		svc service.RoleInfoServiceInterface
	}
	type args struct {
		// c *gin.Context
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
			name:       "Success CreateRoleInfo",
			wantErr:    false,
			statusCode: 200,
			args: args{
				body: dto.RoleInfoCreateRequest{
					Name:        "Admin",
					Description: "Admin",
				},
			},
			fields: fields{
				svc: func() service.RoleInfoServiceInterface {
					mockSvc := new(mocks.MockRoleInfoService)

					mockSvc.On("CreateRoleInfo", mock.Anything, mock.Anything).Return(nil)

					return mockSvc
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.ReleaseMode)

			p := RoleInfoController{
				svc: tt.fields.svc,
			}

			jsonData, err := json.Marshal(tt.args.body)
			if err != nil {
				t.Fatalf("could not marshal mock body: %v", err)
			}

			req := httptest.NewRequest(http.MethodPost, "/create-role", bytes.NewReader(jsonData))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req
			p.CreateRoleInfo(c)

			logrus.Printf("Writer Status = %+v", c.Writer.Status())
			logrus.Printf("Writer Body = %+v", w.Body.Bytes())

			var responseBody map[string]interface{}
			err2 := json.Unmarshal(w.Body.Bytes(), &responseBody)

			if err2 != nil {
				panic(err)
			}

			t.Logf("Response Body = %+v", responseBody)

			if tt.wantErr {
				if tt.statusCode == c.Writer.Status() {
					assert.NotNil(t, c)
					assert.Equal(t, tt.statusCode, c.Writer.Status())
				} else {
					panic("error")
				}

			} else {

				assert.Equal(t, http.StatusOK, c.Writer.Status())
			}
		})
	}
}
