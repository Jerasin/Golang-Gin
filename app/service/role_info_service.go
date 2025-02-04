package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/dto"
	"github.com/Jerasin/app/model"
	"github.com/Jerasin/app/pkg"
	"github.com/Jerasin/app/repository"
	"github.com/Jerasin/app/request"
	"github.com/Jerasin/app/response"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RoleInfoServiceInterface interface {
	CreateRoleInfo(c *gin.Context, body dto.RoleInfoCreateRequest)
	GetPaginationRoleInfo(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.RoleInfo)
	GetRoleInfoById(c *gin.Context)
	UpdateRoleInfo(c *gin.Context)
	DeleteRoleInfo(c *gin.Context)
}

type RoleInfoServiceModel struct {
	BaseRepository repository.BaseRepositoryInterface
}

func RoleInfoServiceInit(baseRepo repository.BaseRepositoryInterface) *RoleInfoServiceModel {
	return &RoleInfoServiceModel{
		BaseRepository: baseRepo,
	}
}

func (p RoleInfoServiceModel) CreateRoleInfo(c *gin.Context, body dto.RoleInfoCreateRequest) {
	fmt.Println("CreateRoleInfo")
	defer pkg.PanicHandler(c)
	var err error
	model := model.RoleInfo{
		Name:        body.Name,
		Description: body.Description,
	}

	// log.Infof("body = %+v \n", model)

	db := p.BaseRepository.ClientDb()

	if db == nil {
		pkg.PanicException(constant.BadRequest)
	}

	// log.Infof("db = %+v \n", db)

	err = db.Transaction(func(tx *gorm.DB) error {
		// log.Infof("Transaction Running...")

		p.BaseRepository.Find(tx, &model, "name = ?", repository.PaginationModel{
			Limit: 1,
		}, body.Name)

		log.Infof("model = %+v \n", model)

		if model.ID != 0 {
			pkg.PanicException(constant.DataIsExit)
		}

		err := p.BaseRepository.Save(tx, &model)

		if err != nil {
			log.Infof("p.BaseRepository.Save Error")
			return err
		}

		return nil
	})

	if err != nil {
		log.Infof("Transaction error = %+v \n", err)
		pkg.PanicException(constant.BadRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.CreateResponse()))
}

func (p RoleInfoServiceModel) GetPaginationRoleInfo(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.RoleInfo) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get list roleInfo")

	log.Info("start to execute get all data user")
	offset := (page - 1) * pageSize
	limit := pageSize
	fields := DbHandleSelectField(field)
	// fmt.Println("query", search)
	// fmt.Println("fields", fields)
	roleInfos := []model.RoleInfo{}

	paginationModel := repository.PaginationModel{
		Limit:     limit,
		Offset:    offset,
		Search:    search,
		SortField: sortField,
		SortValue: sortValue,
		Field:     fields,
		Dest:      roleInfos,
	}

	log.Infof("Dest = %+v \n", roleInfos)
	log.Infof("paginationModel = %+v \n", paginationModel)
	fmt.Println("Before Pagination")
	data, err := p.BaseRepository.Pagination(paginationModel, nil)
	fmt.Println("After Pagination data", data)

	if err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	fmt.Println("Before TotalPage")
	totalPage, err := p.BaseRepository.TotalPage(&roleInfos, pageSize)
	fmt.Println("After TotalPage", totalPage)
	if err != nil {
		log.Error("Count Data Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	fmt.Println("data", data)

	var res []response.RoleInfo
	pkg.ModelDump(&res, data)

	fmt.Println("res", res)

	c.JSON(http.StatusOK, pkg.BuildPaginationResponse(constant.Success, res, totalPage, page, pageSize))
}

func (p RoleInfoServiceModel) GetRoleInfoById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	roleInfoID, _ := strconv.Atoi(c.Param("roleInfoID"))

	var roleInfo model.RoleInfo
	err := p.BaseRepository.FindOne(nil, &roleInfo, "id = ?", roleInfoID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	var res response.RoleInfo
	pkg.ModelDump(&res, roleInfo)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, res))
}

func (p RoleInfoServiceModel) UpdateRoleInfo(c *gin.Context) {
	defer pkg.PanicHandler(c)

	p.BaseRepository.ClientDb().Transaction(func(tx *gorm.DB) error {
		var err error
		roleInfoID, _ := strconv.Atoi(c.Param("roleInfoID"))
		var request request.UpdateRoleInfo

		err = c.ShouldBindJSON(&request)
		if err != nil {
			pkg.PanicException(constant.BadRequest)
		}

		var roleInfo model.RoleInfo
		err = p.BaseRepository.FindOne(tx, &roleInfo, "id = ?", roleInfoID)
		if err != nil {
			log.Error("Happened error when get data from database. Error", err)
			pkg.PanicException(constant.DataNotFound)
		}

		err = p.BaseRepository.Update(tx, roleInfoID, &roleInfo, &request)
		if err != nil {
			pkg.PanicException(constant.BadRequest)
		}

		c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.UpdateResponse()))
		return nil
	})

}

func (p RoleInfoServiceModel) DeleteRoleInfo(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data user by id")
	roleInfoID, _ := strconv.Atoi(c.Param("roleInfoID"))

	var permissionInfo model.RoleInfo
	err := p.BaseRepository.Delete(&permissionInfo, roleInfoID)
	if err != nil {
		log.Error("Happened Error when try delete data user from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.DeleteResponse()))
}
