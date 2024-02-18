package userController

import (
	"net/http"

	"github.com/Jerasin/models"
	"github.com/Jerasin/utils"
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

// GET BY ID
func (db *DBController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	result := db.Database.First(&user, id)

	if result.Error != nil {
		// Handle error...
		c.JSON(http.StatusNotFound, gin.H{"results": "Data Not Found"})
		return
	}

	// db.Database.Model(&user)

	c.JSON(http.StatusOK, gin.H{"results": &user})
}

// POST
func (db *DBController) CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)

	// log.Fatal("ShouldBind Error = ", err != nil)

	if err != nil {
		// log.Fatal("ShouldBind Error")
		utils.LoggerInfo(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// log.Fatalln("test")

	result := db.Database.First(&user, "username = ?", user.Username)

	// Check FindOne User If Error 'record not found' that mean Not Have User In DB
	if result.Error != nil {
		utils.LoggerInfo(result.Error.Error())
		hash_password, _ := utils.HashPassword(user.Password)
		user.Password = hash_password

		result = db.Database.Create(&user)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request."})

		} else {
			c.JSON(http.StatusOK, gin.H{"results": &user})
		}

	} else {
		// Handle error...
		c.JSON(http.StatusBadRequest, gin.H{"message": "Duplicated User"})
		return
	}

}

// PUT
func (db *DBController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	err := c.ShouldBind(&user)
	body := user

	result := db.Database.First(&user, id)

	if result.Error != nil || err != nil {
		// Handle error...
		c.JSON(http.StatusNotFound, gin.H{"results": "Data Not Found"})
		return
	}

	result = db.Database.Where("id = ?", id).Updates(body)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request."})
	} else {
		c.JSON(http.StatusOK, gin.H{"results": &body})
	}
}

// DELETE
func (db *DBController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	db.Database.Delete(&user, id)

	c.JSON(http.StatusOK, gin.H{"message": http.StatusOK})
}
