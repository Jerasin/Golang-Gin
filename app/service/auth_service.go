package service

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/model"
	"github.com/Jerasin/app/pkg"
	"github.com/Jerasin/app/repository"
	"github.com/Jerasin/app/request"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServiceInterface interface {
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
	Register(c *gin.Context)
}

type AuthServiceModel struct {
	BaseRepository repository.BaseRepositoryInterface
	UserRepository repository.UserRepositoryInterface
	UserService    UserServiceInterface
}

func AuthServiceInit(baseRepo repository.BaseRepositoryInterface, userRepo repository.UserRepositoryInterface, userSvc UserServiceInterface) *AuthServiceModel {
	return &AuthServiceModel{
		BaseRepository: baseRepo,
		UserRepository: userRepo,
		UserService:    userSvc,
	}
}

func (authSvc AuthServiceModel) Register(c *gin.Context) {
	defer pkg.PanicHandler(c)

	authSvc.UserService.CreateUser(c)
}

func (authSvc AuthServiceModel) Login(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var request request.LoginRequest
	var user model.User
	var err error

	fmt.Println("request", request)

	if err = c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.CustomPanicException(constant.InvalidRequest, err.Error())
	}

	query := make(map[interface{}]interface{})
	query["username"] = request.Username

	err = authSvc.BaseRepository.FindOne(nil, &user, "username = ?", query["username"])

	if err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			pkg.PanicException(constant.DataNotFound)
		}

		pkg.PanicException(constant.InvalidRequest)
	}

	fmt.Println("user", user)

	isError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	fmt.Println("isError", isError)

	if isError != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	// var res response.User
	// copier.Copy(&res, &result)

	jwt := pkg.NewAuthService()

	token := jwt.GenerateToken(user.Username)

	var response = make(map[string]interface{})
	response["token"] = token
	response["refresh_token"] = jwt.GenerateRefreshToken(user.Username)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}

func (authSvc AuthServiceModel) RefreshToken(c *gin.Context) {
	defer pkg.PanicHandler(c)

	tokenReq := request.TokenReqBody{}

	if err := c.ShouldBindJSON(&tokenReq); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.CustomPanicException(constant.InvalidRequest, err.Error())
	}

	jwtService := pkg.NewAuthService()

	token, err := jwtService.ValidateToken(tokenReq.RefreshToken)
	if err != nil {
		panic(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		pkg.CustomPanicException(constant.InvalidRequest, "token claims is not of type jwt.MapClaims")
	}

	username, ok := claims["username"].(string)
	if !ok {
		// Handle case where username is not a string
		panic("username claim is not a string")
	}

	fmt.Println("claims", claims["username"])
	fmt.Printf("username %T\n", username)

	refreshToken := jwtService.GenerateToken(username)

	var response = make(map[string]interface{})
	response["refreshToken"] = refreshToken
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}
