package pkg

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/Jerasin/app/constant"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func PanicDatabaseException(err error, c *gin.Context) {
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrDuplicatedKey):
			c.JSON(http.StatusBadRequest, BuildResponse(constant.Duplicated, "username or password is exits"))
			return
		case errors.Is(err, gorm.ErrInvalidDB):
			c.JSON(http.StatusBadRequest, BuildResponse(constant.Duplicated, "invalid database"))
			return
		case errors.Is(err, gorm.ErrInvalidValue):
			c.JSON(http.StatusBadRequest, BuildResponse(constant.Duplicated, "invalid value"))
			return
		default:
			log.Error("Happened error when saving data to database. Error", err)
			PanicException(constant.UnknownError)
			return
		}
	}
}

func PanicException_(key string, message string) {
	err := errors.New(message)
	err = fmt.Errorf("%s: %w", key, err)
	if err != nil {
		panic(err)
	}
}

func PanicException(responseKey constant.ResponseStatus) {
	PanicException_(responseKey.GetResponseStatus(), responseKey.GetResponseMessage())
}

func CustomPanicException(responseKey constant.ResponseStatus, message string) {
	PanicException_(responseKey.GetResponseStatus(), message)
}

func PanicHandler(c *gin.Context) {
	if err := recover(); err != nil {
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")

		key := strArr[0]
		msg := strings.Trim(strArr[1], " ")

		switch key {
		case
			constant.DataNotFound.GetResponseStatus():
			c.JSON(http.StatusNotFound, BuildResponse_(key, msg, Null()))
			c.Abort()
		case
			constant.BadRequest.GetResponseStatus():
			c.JSON(http.StatusBadRequest, BuildResponse_(key, msg, Null()))
			c.Abort()
		case
			constant.Unauthorized.GetResponseStatus():
			c.JSON(http.StatusUnauthorized, BuildResponse_(key, msg, Null()))
			c.Abort()
		default:
			c.JSON(http.StatusInternalServerError, BuildResponse_(key, msg, Null()))
			c.Abort()
		}
	}
}
