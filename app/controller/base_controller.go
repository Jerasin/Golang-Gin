package controller

import (
	"strconv"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/pkg"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type BasePaginationModel struct {
	page      int
	pageSize  int
	search    string
	sortField string
	sortValue string
}

func CreatePagination(c *gin.Context) *BasePaginationModel {
	reqPage := c.DefaultQuery("page", "1")
	reqPageSize := c.DefaultQuery("pageSize", "15")
	search := c.DefaultQuery("search", "")
	sortField := c.DefaultQuery("sortField", "updated_at")
	sortValue := c.DefaultQuery("sortValue", "desc")

	page, err := strconv.Atoi(reqPage)

	if err != nil {
		log.Error("PaginationModel Convert Data Error: ", err)
		pkg.PanicException(constant.ValidateError)
	}

	pageSize, err := strconv.Atoi(reqPageSize)
	if err != nil {
		log.Error("PaginationModel Convert Data Error: ", err)
		pkg.PanicException(constant.ValidateError)
	}

	return &BasePaginationModel{
		page:      page,
		pageSize:  pageSize,
		search:    search,
		sortField: sortField,
		sortValue: sortValue,
	}
}
