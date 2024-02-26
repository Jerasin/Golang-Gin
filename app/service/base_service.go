package service

import (
	"errors"
	"net/http"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/pkg"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func DbHandleError(err error, c *gin.Context) {
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrDuplicatedKey):
			c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.Duplicated, "username or password is exits"))
			return
		default:
			log.Error("Happened error when saving data to database. Error", err)
			pkg.PanicException(constant.UnknownError)
			return
		}
	}
}
