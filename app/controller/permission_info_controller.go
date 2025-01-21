package controller

import (
	"github.com/Jerasin/app/response"
	"github.com/Jerasin/app/service"
	"github.com/gin-gonic/gin"
)

type PermissionInfoController struct {
	svc service.PermissionInfoServiceInterface
}

type PermissionInfoControllerInterface interface {
	CreatePermissionInfo(c *gin.Context)
	GetListPermissionInfo(c *gin.Context)
	GetPermissionInfoById(c *gin.Context)
	UpdatePermissionInfoData(c *gin.Context)
	DeletePermissionInfo(c *gin.Context)
}

func PermissionInfoControllerInit(permissionInfoSvc service.PermissionInfoServiceInterface) *PermissionInfoController {
	return &PermissionInfoController{
		svc: permissionInfoSvc,
	}
}

// @Summary Create PermissionInfo
// @Schemes
// @Description Create PermissionInfo
// @Tags PermissionInfo
//
// @Param request body request.PermissionInfoRequest true "query params"
//
//	@Success		200	{object}	response.CreateDataResponse
//
// @Security Bearer
//
// @Router /permission_infos [post]
func (p PermissionInfoController) CreatePermissionInfo(c *gin.Context) {
	p.svc.CreatePermissionInfo(c)
}

// @Summary Get PermissionInfo List
// @Schemes
// @Description Get PermissionInfo List
// @Tags PermissionInfo
//
// @Param   page         query     int        false  "int valid"
// @Param   pageSize         query     int        false  "int valid"
// @Param   sortField         query     string        false  "string valid"
// @Param   sortValue         query     string        false  "string valid"
//
//	@Success		200	{object}	response.PermissionInfoPagination
//
// @Security Bearer
//
// @Router /permission_infos [get]
func (p PermissionInfoController) GetListPermissionInfo(c *gin.Context) {
	query := CreatePagination(c)
	permissionInfo := response.PermissionInfo{}
	p.svc.GetPaginationPermissionInfo(c, query.page, query.pageSize, query.search, query.sortField, query.sortValue, permissionInfo)
}

// @Summary Get PermissionInfo By Id
// @Schemes
// @Description Get PermissionInfo By Id
// @Tags PermissionInfo
// @Param permissionInfoID  path int true "PermissionInfo ID"
//
//	@Success		200	{object}	response.PermissionInfo
//
// @Security Bearer
//
// @Router /permission_infos/{permissionInfoID} [get]
func (p PermissionInfoController) GetPermissionInfoById(c *gin.Context) {
	p.svc.GetPermissionInfoById(c)
}

// @Summary Update PermissionInfo By Id
// @Schemes
// @Description Update PermissionInfo By Id
// @Tags PermissionInfo
// @Param permissionInfoID  path int true "PermissionInfo ID"
// @Param request body request.UpdatePermissionInfo true "query params"
//
//	@Success		200	{object}	response.UpdateDataResponse
//
// @Security Bearer
//
// @Router /permission_infos/{permissionInfoID} [put]
func (p PermissionInfoController) UpdatePermissionInfoData(c *gin.Context) {
	p.svc.UpdatePermissionInfo(c)
}

// @Summary Delete PermissionInfo By Id
// @Schemes
// @Description Delete PermissionInfo By Id
// @Tags PermissionInfo
// @Param permissionInfoID  path int true "PermissionInfo ID"
//
//	@Success		200	{object}	response.DeleteDataResponse
//
// @Security Bearer
//
// @Router /permission_infos/{permissionInfoID} [delete]
func (p PermissionInfoController) DeletePermissionInfo(c *gin.Context) {
	p.svc.DeletePermissionInfo(c)
}
