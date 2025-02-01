package controller

import (
	"github.com/Jerasin/app/response"
	"github.com/Jerasin/app/service"
	"github.com/gin-gonic/gin"
)

type ProductControllerInterface interface {
	CreateProduct(c *gin.Context)
	GetAllProducts(c *gin.Context)
	GetProductById(c *gin.Context)
	UpdateProductData(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

type ProductController struct {
	svc service.ProductServiceInterface
}

func ProductControllerInit(productSvc service.ProductServiceInterface) *ProductController {
	return &ProductController{
		svc: productSvc,
	}
}

// @Summary Create Product
// @Schemes
// @Description Create Product
// @Tags Product
//
// @Param request body request.Product true "query params"
//
//	@Success		200	{object}	response.CreateDataResponse
//
// @Security Bearer
//
// @Router /products [post]
func (p ProductController) CreateProduct(c *gin.Context) {
	p.svc.CreateProduct(c)
}

// @Summary Get Product List
// @Schemes
// @Description Get Product List
// @Tags Product
//
// @Param   page         query     int        false  "int valid"
// @Param   pageSize         query     int        false  "int valid"
// @Param   sortField         query     string        false  "string valid"
// @Param   sortValue         query     string        false  "string valid"
//
//	@Success		200	{object}	response.ProductPagination
//
// @Security Bearer
//
// @Router /products [get]
func (p ProductController) GetAllProducts(c *gin.Context) {
	query := CreatePagination(c)
	Product := response.Product{}

	p.svc.GetPaginationProduct(c, query.page, query.pageSize, query.search, query.sortField, query.sortValue, Product)
}

// @Summary Get Product By Id
// @Schemes
// @Description Get Product By Id
// @Tags Product
// @Param productID  path int true "Product ID"
//
//	@Success		200	{object}	response.Product
//
// @Security Bearer
//
// @Router /products/{productID} [get]
func (p ProductController) GetProductById(c *gin.Context) {
	p.svc.GetProductById(c)
}

// @Summary Update Product By Id
// @Schemes
// @Description Update Product By Id
// @Tags Product
// @Param productID  path int true "Product ID"
// @Param request body request.UpdateProduct true "query params"
//
//	@Success		200	{object}	response.UpdateDataResponse
//
// @Security Bearer
//
// @Router /products/{productID} [put]
func (p ProductController) UpdateProductData(c *gin.Context) {
	p.svc.UpdateProduct(c)
}

// @Summary Delete Product By Id
// @Schemes
// @Description Delete Product By Id
// @Tags Product
// @Param productID  path int true "Product ID"
//
//	@Success		200	{object}	response.DeleteDataResponse
//
// @Security Bearer
//
// @Router /products/{productID} [delete]
func (p ProductController) DeleteProduct(c *gin.Context) {
	p.svc.DeleteProduct(c)
}
