package controller

import (
	"github.com/Jerasin/app/response"
	"github.com/Jerasin/app/service"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	svc service.OrderServiceInterface
}

type OrderControllerInterface interface {
	CreateOrder(c *gin.Context)
	GetAllProducts(c *gin.Context)
}

func OrderControllerInit(orderSvc service.OrderServiceInterface) *OrderController {
	return &OrderController{
		svc: orderSvc,
	}
}

// @Summary Create Order
// @Schemes
// @Description Create Order
// @Tags Order
//
// @Param request body request.OrderRequest true "query params"
//
//	@Success		200	{object}	model.Order
//
// @Security Bearer
//
// @Router /orders [post]
func (o OrderController) CreateOrder(c *gin.Context) {
	o.svc.CreateOrder(c)
}

// @Summary Get Order List
// @Schemes
// @Description Get Order List
// @Tags Order
//
// @Param   page         query     int        false  "int valid"
// @Param   pageSize         query     int        false  "int valid"
// @Param   sortField         query     string        false  "string valid"
// @Param   sortValue         query     string        false  "string valid"
//
//	@Success		200	{object}	response.OrderPagination
//
// @Security Bearer
//
// @Router /orders [get]
func (o OrderController) GetAllProducts(c *gin.Context) {
	query := CreatePagination(c)
	order := response.Order{}

	o.svc.GetPaginationOrder(c, query.page, query.pageSize, query.search, query.sortField, query.sortValue, order)
}
