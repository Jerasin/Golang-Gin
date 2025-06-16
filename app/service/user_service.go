package service

import (
	"fmt"
	"net/http"
	"net/mail"
	"strconv"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/model"
	"github.com/Jerasin/app/pkg"
	"github.com/Jerasin/app/repository"
	"github.com/Jerasin/app/request"
	"github.com/Jerasin/app/response"
	"github.com/goforj/godump"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	GetPaginationUser(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.User)
	GetUserById(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetUserInfo(c *gin.Context)
	// GetUser(c *gin.Context, user model.User, query map[interface{}]interface{}, field response.User) model.User
}

type UserServiceModel struct {
	BaseRepository repository.BaseRepositoryInterface
	UserRepository repository.UserRepositoryInterface
}

func UserServiceInit(baseRepo repository.BaseRepositoryInterface, userRepo repository.UserRepositoryInterface) *UserServiceModel {
	return &UserServiceModel{
		BaseRepository: baseRepo,
		UserRepository: userRepo,
	}
}

func (u UserServiceModel) CreateUser(c *gin.Context) {
	defer pkg.PanicHandler(c)
	err := u.BaseRepository.ClientDb().Transaction(func(tx *gorm.DB) error {
		var err error
		var request model.User

		log.Info("start to execute program add data user")

		if err = c.ShouldBindJSON(&request); err != nil {
			return err
		}

		_, err = mail.ParseAddress(request.Email)
		if err != nil {
			log.Error("Happened error when mapping request from FE. Error", err)
			return err
		}

		godump.Dump(request)

		hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 15)
		request.Password = string(hash)

		err = u.BaseRepository.Save(tx, &request)
		if err != nil {
			return err
		}

		c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.CreateResponse()))
		return nil
	})

	if err != nil {
		pkg.PanicException(constant.BadRequest)
	}
}

func (u UserServiceModel) UpdateUser(c *gin.Context) {
	defer pkg.PanicHandler(c)

	u.BaseRepository.ClientDb().Transaction(func(tx *gorm.DB) error {
		var err error
		log.Info("start to execute program update user data by id")
		userID, _ := strconv.Atoi(c.Param("userID"))

		var request request.UpdateUser
		if err = c.ShouldBindJSON(&request); err != nil {
			log.Error("Happened error when mapping request from FE. Error", err)
			pkg.PanicException(constant.InvalidRequest)
		}

		var user model.User
		err = u.BaseRepository.FindOne(tx, &user, "id = ?", userID)
		if err != nil {
			log.Error("Happened error when get data from database. Error", err)
			pkg.PanicException(constant.DataNotFound)
		}

		godump.Dump(request)
		updateData := map[string]any{
			"username":  request.Username,
			"fullname":  request.Fullname,
			"avatar":    request.Avatar,
			"is_active": request.IsActive,
		}
		updateError := u.BaseRepository.Update(tx, userID, &user, &updateData)

		if updateError != nil {
			log.Error("Happened error when updating data to database. Error", err)
			pkg.PanicException(constant.UnknownError)
		}

		c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.UpdateResponse()))
		return nil
	})

}

func (u UserServiceModel) GetUserById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	var user response.User
	var wallets []response.Wallet
	err := u.BaseRepository.FindOne(nil, &user, "id = ?", userID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	err = u.BaseRepository.Find(nil, &wallets, "user_id = ?", repository.PaginationModel{}, userID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, map[string]any{"user": user, "wallets": wallets}))
}

func (u UserServiceModel) GetPaginationUser(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.User) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get all data user")
	offset := (page - 1) * pageSize
	limit := pageSize
	fields := DbHandleSelectField(field)

	godump.Dump("fields", fields)

	var users []model.User
	paginationModel := repository.PaginationModel{
		Limit:     limit,
		Offset:    offset,
		Search:    search,
		SortField: sortField,
		SortValue: sortValue,
		Field:     fields,
		Dest:      users,
	}

	data, err := u.BaseRepository.Pagination(paginationModel, nil)
	if err != nil {
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	totalPage, err := u.BaseRepository.TotalPage(&users, pageSize)
	if err != nil {
		log.Error("Count Data Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	total, totalErr := u.BaseRepository.Total(&users)
	if totalErr != nil {
		log.Error("Count Data Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	fmt.Println("count", totalPage)

	var res []response.User
	pkg.ModelDump(&res, data)
	c.JSON(http.StatusOK, pkg.BuildPaginationResponse(constant.Success, res, totalPage, page, pageSize, total))
}

func (u UserServiceModel) DeleteUser(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))
	var user model.User
	err := u.BaseRepository.Delete(&user, userID)
	if err != nil {
		log.Error("Happened Error when try delete data user from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.DeleteResponse()))
}

func (u UserServiceModel) GetUserInfo(c *gin.Context) {
	defer pkg.PanicHandler(c)

	id, ok := c.Get("userID")
	if !ok {
		pkg.PanicException(constant.BadRequest)
	}

	godump.Dump(id)

	userID, ok := id.(uint)
	if !ok {
		pkg.PanicException(constant.BadRequest)
	}

	var user model.User
	var roleInfo model.RoleInfo
	err := u.BaseRepository.FindOne(nil, &user, "id = ?", userID)
	if err != nil {
		pkg.PanicException(constant.UnknownError)
	}
	options := repository.Options{
		Query:     "role_infos.id = ?",
		QueryArgs: []any{user.RoleInfoID},
		Preloads:  []string{"PermissionInfos"},
	}
	err = u.BaseRepository.FindOneV2(nil, &roleInfo, options)
	if err != nil {
		pkg.PanicException(constant.UnknownError)
	}

	// godump.Dump(user)

	res := response.UserInfo{
		BaseModel: model.BaseModel{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Email:      user.Email,
		Username:   user.Username,
		FullName:   user.Fullname,
		Avatar:     user.Avatar,
		RoleInfoID: user.RoleInfoID,
		RoleInfo: response.UserRoleInfo{
			Name:            roleInfo.Name,
			Description:     roleInfo.Description,
			PermissionInfos: roleInfo.PermissionInfos,
		},
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, res))
}

// func (u UserServiceModel) GetUser(c *gin.Context, user model.User, query map[interface{}]interface{}, field response.User) model.User {
// 	defer pkg.PanicHandler(c)

// 	fields := structs.Map(field)

// 	log.Info("start to execute get data user")
// 	result, err := u.UserRepository.FindOneUser(user, query, fields)
// 	if err != nil {
// 		log.Error("Happened Error when try delete data user from DB. Error:", err)
// 		pkg.PanicException(constant.UnknownError)
// 	}

// 	return result
// }
