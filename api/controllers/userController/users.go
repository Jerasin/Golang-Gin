package userController

import (
	"net/http"

	"github.com/Jerasin/models"
	"github.com/gin-gonic/gin"
)

// GET
func (db *DBController) GetUsers(c *gin.Context) {
	_type := c.Query("type")
	_where := map[string]interface{}{}

	if _type != "" {
		_where["type"] = _type
	}

	var user []models.User
	db.Database.Where(_where).Find(&user)

	c.JSON(http.StatusOK, gin.H{"results": &user})
}

// POST
func (db *DBController) CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)

	result := db.Database.Create(&user)

	if result.Error != nil || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meassage": "Bad request."})
	} else {
		c.JSON(http.StatusOK, gin.H{"results": &user})
	}
}
