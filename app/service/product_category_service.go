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

type ProductCategoryServiceInterface interface {
	CreateProductCategory(c *gin.Context)
	GetPaginationProductCategory(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.ProductCategory)
	GetProductCategoryById(c *gin.Context)
	UpdateProductCategory(c *gin.Context)
	DeleteProductCategory(c *gin.Context)
}

type ProductCategoryServiceModel struct {
	BaseRepository      repository.BaseRepositoryInterface
	ProductCategoryRepo repository.ProductCategoryRepositoryInterface
}

func ProductCategoryServiceInit(baseRepo repository.BaseRepositoryInterface, productCategoryRepo repository.ProductCategoryRepositoryInterface) *ProductCategoryServiceModel {
	return &ProductCategoryServiceModel{
		ProductCategoryRepo: productCategoryRepo,
		BaseRepository:      baseRepo,
	}
}

func (p ProductCategoryServiceModel) CreateProductCategory(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data product category")

	p.BaseRepository.ClientDb().Transaction(func(tx *gorm.DB) error {
		var request model.ProductCategory
		if err := c.ShouldBind(&request); err != nil {
			log.Error("Happened error when mapping request from FE. Error", err)
			pkg.PanicException(constant.InvalidRequest)
		}

		err := p.BaseRepository.Save(tx, &request)

		if err != nil {
			log.Error("Happened error when mapping request from FE. Error", err)
			pkg.PanicException(constant.InvalidRequest)
		}

		c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.CreateResponse()))
		return nil
	})

}

func (p ProductCategoryServiceModel) GetPaginationProductCategory(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.ProductCategory) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get list product category")

	log.Info("start to execute get all data user")
	offset := (page - 1) * pageSize
	limit := pageSize
	fields := DbHandleSelectField(field)
	fmt.Println("query", search)
	fmt.Println("fields", fields)
	var productCategories []model.ProductCategory

	paginationModel := repository.PaginationModel{
		Limit:     limit,
		Offset:    offset,
		Search:    search,
		SortField: sortField,
		SortValue: sortValue,
		Field:     fields,
		Dest:      productCategories,
	}
	data, err := p.BaseRepository.Pagination(paginationModel, nil)

	if err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	totalPage, err := p.BaseRepository.TotalPage(&productCategories, pageSize)
	if err != nil {
		log.Error("Count Data Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	fmt.Println("data", data)

	var res []response.ProductCategory
	pkg.ModelDump(&res, data)

	c.JSON(http.StatusOK, pkg.BuildPaginationResponse(constant.Success, res, totalPage, page, pageSize))
}

func (p ProductCategoryServiceModel) GetProductCategoryById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	productCategoryID, _ := strconv.Atoi(c.Param("productCategoryID"))

	var productCategory model.ProductCategory
	err := p.BaseRepository.FindOne(nil, &productCategory, "id = ?", productCategoryID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, productCategory))
}

func (p ProductCategoryServiceModel) UpdateProductCategory(c *gin.Context) {
	defer pkg.PanicHandler(c)

	p.BaseRepository.ClientDb().Transaction(func(tx *gorm.DB) error {
		var err error
		productCategoryID, _ := strconv.Atoi(c.Param("productCategoryID"))
		var request request.UpdateProductCategory

		err = c.ShouldBindJSON(&request)
		if err != nil {
			pkg.PanicException(constant.BadRequest)
		}

		var productCategory model.ProductCategory
		err = p.BaseRepository.FindOne(tx, &productCategory, "id = ?", productCategoryID)
		if err != nil {
			log.Error("Happened error when get data from database. Error", err)
			pkg.PanicException(constant.DataNotFound)
		}

		err = p.BaseRepository.Update(tx, productCategoryID, &productCategory, &request)
		if err != nil {
			pkg.PanicException(constant.BadRequest)
		}

		c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.UpdateResponse()))
		return nil
	})

}

func (p ProductCategoryServiceModel) DeleteProductCategory(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data user by id")
	productCategoryID, _ := strconv.Atoi(c.Param("productCategoryID"))

	var productCategory model.ProductCategory
	err := p.BaseRepository.Delete(&productCategory, productCategoryID)
	if err != nil {
		log.Error("Happened Error when try delete data user from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.DeleteResponse()))
}
