package controller

import (
	"github.com/Jerasin/app/response"
	"github.com/Jerasin/app/service"
	"github.com/gin-gonic/gin"
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

	p.svc.CreateRoleInfo(c)
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
	p.svc.UpdateRoleInfo(c)
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
