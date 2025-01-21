package controller

import (
	"github.com/Jerasin/app/response"
	"github.com/Jerasin/app/service"
	"github.com/gin-gonic/gin"
)

type ProductCategoryControllerInterface interface {
	GetListProductCategory(c *gin.Context)
	CreateProductCategory(c *gin.Context)
	GetProductCategoryById(c *gin.Context)
	UpdateProductCategoryData(c *gin.Context)
	DeleteProductCategory(c *gin.Context)
}

type ProductCategoryController struct {
	svc service.ProductCategoryServiceInterface
}

func ProductCategoryControllerInit(productCategorySvc service.ProductCategoryServiceInterface) *ProductCategoryController {
	return &ProductCategoryController{
		svc: productCategorySvc,
	}
}

// @Summary Create ProductCategory
// @Schemes
// @Description Create ProductCategory
// @Tags ProductCategory
//
// @Param request body request.ProductCategoryRequest true "query params"
//
//	@Success		200	{object}	response.CreateDataResponse
//
// @Security Bearer
//
// @Router /products/categories [post]
func (p ProductCategoryController) CreateProductCategory(c *gin.Context) {
	p.svc.CreateProductCategory(c)
}

// @Summary Get ProductCategory List
// @Schemes
// @Description Get ProductCategory List
// @Tags ProductCategory
//
// @Param   page         query     int        false  "int valid"
// @Param   pageSize         query     int        false  "int valid"
// @Param   sortField         query     string        false  "string valid"
// @Param   sortValue         query     string        false  "string valid"
//
//	@Success		200	{object}	response.ProductCategoryPagination
//
// @Security Bearer
//
// @Router /products/categories [get]
func (p ProductCategoryController) GetListProductCategory(c *gin.Context) {
	query := CreatePagination(c)
	productCategory := response.ProductCategory{}
	p.svc.GetPaginationProductCategory(c, query.page, query.pageSize, query.search, query.sortField, query.sortValue, productCategory)
}

// @Summary Get ProductCategory By Id
// @Schemes
// @Description Get ProductCategory By Id
// @Tags ProductCategory
// @Param productCategoryID  path int true "User ID"
//
//	@Success		200	{object}	response.ProductCategory
//
// @Security Bearer
//
// @Router /products/categories/{productCategoryID} [get]
func (p ProductCategoryController) GetProductCategoryById(c *gin.Context) {
	p.svc.GetProductCategoryById(c)
}

// @Summary Update ProductCategory By Id
// @Schemes
// @Description Update ProductCategory By Id
// @Tags ProductCategory
// @Param productCategoryID  path int true "User ID"
// @Param request body request.UpdateProductCategory true "query params"
//
//	@Success		200	{object}	response.UpdateDataResponse
//
// @Security Bearer
//
// @Router /products/categories/{productCategoryID} [put]
func (p ProductCategoryController) UpdateProductCategoryData(c *gin.Context) {
	p.svc.UpdateProductCategory(c)
}

// @Summary Delete ProductCategory By Id
// @Schemes
// @Description Delete ProductCategory By Id
// @Tags ProductCategory
// @Param productCategoryID  path int true "User ID"
//
//	@Success		200	{object}	response.DeleteDataResponse
//
// @Security Bearer
//
// @Router /products/categories/{productCategoryID} [delete]
func (p ProductCategoryController) DeleteProductCategory(c *gin.Context) {
	p.svc.DeleteProductCategory(c)
}
