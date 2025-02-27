package controller

import (
	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/dto"
	"github.com/Jerasin/app/pkg"
	"github.com/Jerasin/app/service"
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

type AuthControllerInterface interface {
	Test(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type AuthController struct {
	svc service.AuthServiceInterface
}

func AuthControllerInit(authService service.AuthServiceInterface) *AuthController {
	return &AuthController{svc: authService}
}

func DeferTest(c *gin.Context) {
	if err := recover(); err != nil {
		c.JSON(200, gin.H{
			"message": err,
		})
	}

}

func (auth AuthController) Test(c *gin.Context) {
	defer DeferTest(c)
	panic("Test")
	// fmt.Println("Test")
}

// @Summary Register
// @Schemes
// @Description Register
// @Tags Auth
//
// @Param request body request.UserRequest true "query params"
//
//	@Success		200	{object}	model.User
//
// @Router /auth/register [post]
func (auth AuthController) Register(c *gin.Context) {
	auth.svc.Register(c)
}

// @Summary Login
// @Schemes
// @Description Login
// @Tags Auth
//
// @Param request body dto.LoginDtoRequest true "query params"
//
//	@Success		200	{object}	dto.LoginDtoResponse
//
// @Router /auth/login [post]
func (auth AuthController) Login(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var loginDto dto.LoginDtoRequest
	var err error
	err = c.ShouldBindJSON(&loginDto)

	if err != nil {
		pkg.PanicException(constant.BadRequest)
	}

	if err = validator.Validate(&loginDto); err != nil {
		pkg.PanicException(constant.BadRequest)
	}

	auth.svc.Login(c, loginDto)
}

// @Summary RefreshToken
// @Schemes
// @Description RefreshToken
// @Tags Auth
//
// @Param request body request.TokenReqBody true "query params"
//
//	@Success		200	{object}	model.User
//
// @Router /auth/refresh/token [post]
func (auth AuthController) RefreshToken(c *gin.Context) {
	auth.svc.RefreshToken(c)
}
