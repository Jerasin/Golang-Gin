package controller

import (
	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/dto"
	"github.com/Jerasin/app/pkg"
	"github.com/Jerasin/app/response"
	"github.com/Jerasin/app/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/validator.v2"
)

type RoleInfoController struct {
	svc service.RoleInfoServiceInterface
}

type RoleInfoControllerInterface interface {
	CreateRoleInfo(c *gin.Context)
	GetListRoleInfo(c *gin.Context)
	GetRoleInfoById(c *gin.Context)
	UpdateRoleInfoData(c *gin.Context)
	DeleteRoleInfo(c *gin.Context)
}

func RoleInfoControllerInit(roleInfoSvc service.RoleInfoServiceInterface) *RoleInfoController {
	return &RoleInfoController{
		svc: roleInfoSvc,
	}
}

// @Summary Create RoleInfo
// @Schemes
// @Description Create RoleInfo
// @Tags RoleInfo
//
// @Param request body request.RoleInfoRequest true "query params"
//
//	@Success		200	{object}	response.CreateDataResponse
//
// @Security Bearer
//
// @Router /role_infos [post]
func (p RoleInfoController) CreateRoleInfo(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var err error
	body := dto.RoleInfoCreateRequest{}
	err = c.ShouldBindJSON(&body)

	log.Infof("body = %+v \n", body)

	if err != nil {
		pkg.PanicException(constant.BadRequest)
	}

	if err = validator.Validate(&body); err != nil {
		pkg.PanicException(constant.BadRequest)
	}

	log.Infoln("Before CreateRoleInfo")
	p.svc.CreateRoleInfo(c, body)
}

// @Summary Get RoleInfo List
// @Schemes
// @Description Get RoleInfo List
// @Tags RoleInfo
//
// @Param   page         query     int        false  "int valid"
// @Param   pageSize         query     int        false  "int valid"
// @Param   sortField         query     string        false  "string valid"
// @Param   sortValue         query     string        false  "string valid"
//
//	@Success		200	{object}	response.RoleInfoPagination
//
// @Security Bearer
//
// @Router /role_infos [get]
func (p RoleInfoController) GetListRoleInfo(c *gin.Context) {
	query := CreatePagination(c)
	permissionInfo := response.RoleInfo{}
	p.svc.GetPaginationRoleInfo(c, query.page, query.pageSize, query.search, query.sortField, query.sortValue, permissionInfo)
}

// @Summary Get RoleInfo By Id
// @Schemes
// @Description Get RoleInfo By Id
// @Tags RoleInfo
// @Param roleInfoID  path int true "RoleInfo ID"
//
//	@Success		200	{object}	response.RoleInfo
//
// @Security Bearer
//
// @Router /role_infos/{roleInfoID} [get]
func (p RoleInfoController) GetRoleInfoById(c *gin.Context) {
	p.svc.GetRoleInfoById(c)
}

// @Summary Update RoleInfo By Id
// @Schemes
// @Description Update RoleInfo By Id
// @Tags RoleInfo
// @Param roleInfoID  path int true "RoleInfo ID"
// @Param request body request.UpdateRoleInfo true "query params"
//
//	@Success		200	{object}	response.UpdateDataResponse
//
// @Security Bearer
//
// @Router /role_infos/{roleInfoID} [put]
func (p RoleInfoController) UpdateRoleInfoData(c *gin.Context) {
	defer pkg.PanicHandler(c)
	var err error
	body := dto.UpdateRoleInfoCreateRequest{}
	err = c.ShouldBindJSON(&body)

	if err != nil {
		log.Info("Error when binding data")
		pkg.PanicException(constant.BadRequest)
	}

	if err = validator.Validate(&body); err != nil {
		log.Info("Error when validating data")
		pkg.PanicException(constant.BadRequest)
	}

	p.svc.UpdateRoleInfo(c, body)
}

// @Summary Delete RoleInfo By Id
// @Schemes
// @Description Delete RoleInfo By Id
// @Tags RoleInfo
// @Param roleInfoID  path int true "RoleInfo ID"
//
//	@Success		200	{object}	response.DeleteDataResponse
//
// @Security Bearer
//
// @Router /role_infos/{roleInfoID} [delete]
func (p RoleInfoController) DeleteRoleInfo(c *gin.Context) {
	p.svc.DeleteRoleInfo(c)
}
