package controller

import (
	"github.com/Jerasin/app/response"
	"github.com/Jerasin/app/service"
	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	GetAllUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	GetUserById(c *gin.Context)
	UpdateUserData(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserController struct {
	svc service.UserServiceInterface
}

func UserControllerInit(userService service.UserServiceInterface) *UserController {
	return &UserController{
		svc: userService,
	}
}

// @Summary Get User List
// @Schemes
// @Description Get User List
// @Tags User
//
// @Param   page         query     int        false  "int valid"
// @Param   pageSize         query     int        false  "int valid"
// @Param   sortField         query     string        false  "string valid"
// @Param   sortValue         query     string        false  "string valid"
//
//	@Success		200	{object}	response.UserPagination
//
// @Security Bearer
//
// @Router /users [get]
func (u UserController) GetAllUsers(c *gin.Context) {
	query := CreatePagination(c)
	user := response.User{}

	u.svc.GetPaginationUser(c, query.page, query.pageSize, query.search, query.sortField, query.sortValue, user)
}

// @Summary Create User
// @Schemes
// @Description Create User
// @Tags User
//
// @Param request body request.UserRequest true "query params"
//
//	@Success		200	{object}	response.CreateDataResponse
//
// @Security Bearer
//
// @Router /users [post]
func (u UserController) CreateUser(c *gin.Context) {
	u.svc.CreateUser(c)
}

// @Summary Get User By Id
// @Schemes
// @Description Get user By Id
// @Tags User
// @Param userID  path int true "User ID"
//
//	@Success		200	{object}	response.User
//
// @Security Bearer
//
// @Router /users/{userID} [get]
func (u UserController) GetUserById(c *gin.Context) {
	u.svc.GetUserById(c)
}

// @Summary Update User By Id
// @Schemes
// @Description Update user By Id
// @Tags User
// @Param userID  path int true "User ID"
// @Param request body request.UserRequest true "query params"
//
//	@Success		200	{object}	response.UpdateDataResponse
//
// @Security Bearer
//
// @Router /users/{userID} [put]
func (u UserController) UpdateUserData(c *gin.Context) {
	u.svc.UpdateUser(c)
}

// @Summary Delete User By Id
// @Schemes
// @Description Delete user By Id
// @Tags User
// @Param userID  path int true "User ID"
//
//	@Success		200	{object}	response.DeleteDataResponse
//
// @Security Bearer
//
// @Router /users/{userID} [delete]
func (u UserController) DeleteUser(c *gin.Context) {
	u.svc.DeleteUser(c)
}
