package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/model"
	"github.com/Jerasin/app/pkg"
	"github.com/Jerasin/app/repository"
	"github.com/Jerasin/app/request"
	"github.com/Jerasin/app/response"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PermissionInfoServiceInterface interface {
	CreatePermissionInfo(c *gin.Context)
	GetPaginationPermissionInfo(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.PermissionInfo)
	GetPermissionInfoById(c *gin.Context)
	UpdatePermissionInfo(c *gin.Context)
	DeletePermissionInfo(c *gin.Context)
}

type PermissionInfoServiceModel struct {
	BaseRepository repository.BaseRepositoryInterface
}

func PermissionInfoServiceInit(baseRepo repository.BaseRepositoryInterface) *PermissionInfoServiceModel {
	return &PermissionInfoServiceModel{
		BaseRepository: baseRepo,
	}
}

func (p PermissionInfoServiceModel) CreatePermissionInfo(c *gin.Context) {
	fmt.Println("CreatePermissionInfo")
	defer pkg.PanicHandler(c)
	var err error
	var body model.PermissionInfo

	err = c.ShouldBindJSON(&body)
	if err != nil {
		pkg.PanicException(constant.BadRequest)
	}

	p.BaseRepository.ClientDb().Transaction(func(tx *gorm.DB) error {
		err := p.BaseRepository.Save(tx, &body)

		if err != nil {
			pkg.PanicException(constant.BadRequest)
		}

		c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.CreateResponse()))
		return nil
	})
}

func (p PermissionInfoServiceModel) GetPaginationPermissionInfo(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.PermissionInfo) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get list permissionInfo")

	log.Info("start to execute get all data user")
	offset := (page - 1) * pageSize
	limit := pageSize
	fields := DbHandleSelectField(field)
	fmt.Println("query", search)
	fmt.Println("fields", fields)
	var permissionInfos []model.PermissionInfo

	paginationModel := repository.PaginationModel{
		Limit:     limit,
		Offset:    offset,
		Search:    search,
		SortField: sortField,
		SortValue: sortValue,
		Field:     fields,
		Dest:      permissionInfos,
	}
	data, err := p.BaseRepository.Pagination(paginationModel, nil)

	if err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	totalPage, err := p.BaseRepository.TotalPage(&permissionInfos, pageSize)
	if err != nil {
		log.Error("Count Data Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	fmt.Println("data", data)

	var res []response.PermissionInfo
	pkg.ModelDump(&res, data)

	c.JSON(http.StatusOK, pkg.BuildPaginationResponse(constant.Success, res, totalPage, page, pageSize))
}

func (p PermissionInfoServiceModel) GetPermissionInfoById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	permissionInfoID, _ := strconv.Atoi(c.Param("permissionInfoID"))

	var permissionInfo model.PermissionInfo
	err := p.BaseRepository.FindOne(nil, &permissionInfo, "id = ?", permissionInfoID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, permissionInfo))
}

func (p PermissionInfoServiceModel) UpdatePermissionInfo(c *gin.Context) {
	defer pkg.PanicHandler(c)

	p.BaseRepository.ClientDb().Transaction(func(tx *gorm.DB) error {
		var err error
		permissionInfoID, _ := strconv.Atoi(c.Param("permissionInfoID"))
		var request request.UpdatePermissionInfo

		err = c.ShouldBindJSON(&request)
		if err != nil {
			pkg.PanicException(constant.BadRequest)
		}

		var permissionInfo model.PermissionInfo
		err = p.BaseRepository.FindOne(tx, &permissionInfo, "id = ?", permissionInfoID)
		if err != nil {
			log.Error("Happened error when get data from database. Error", err)
			pkg.PanicException(constant.DataNotFound)
		}

		err = p.BaseRepository.Update(tx, permissionInfoID, &permissionInfo, &request)
		if err != nil {
			pkg.PanicException(constant.BadRequest)
		}

		c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.UpdateResponse()))
		return nil
	})

}

func (p PermissionInfoServiceModel) DeletePermissionInfo(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data user by id")
	permissionInfoID, _ := strconv.Atoi(c.Param("permissionInfoID"))

	var permissionInfo model.PermissionInfo
	err := p.BaseRepository.Delete(&permissionInfo, permissionInfoID)
	if err != nil {
		log.Error("Happened Error when try delete data user from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.DeleteResponse()))
}
