package controller

import (
	"github.com/Jerasin/app/response"
	"github.com/Jerasin/app/service"
	"github.com/gin-gonic/gin"
)

type WalletControllerInterface interface {
	CreateWallet(c *gin.Context)
	GetListWallet(c *gin.Context)
	GetWalletById(c *gin.Context)
	UpdateWalletData(c *gin.Context)
	DeleteWallet(c *gin.Context)
}

type WalletController struct {
	svc service.WalletServiceInterface
}

func WalletControllerInit(walletService service.WalletServiceInterface) *WalletController {
	return &WalletController{
		svc: walletService,
	}
}

// @Summary Create Wallet
// @Schemes
// @Description Create Wallet
// @Tags Wallet
//
// @Param request body request.WalletRequest true "query params"
//
//	@Success		200	{object}	response.CreateDataResponse
//
// @Security Bearer
//
// @Router /wallets [post]
func (p WalletController) CreateWallet(c *gin.Context) {
	p.svc.CreateWallet(c)
}

// @Summary Get Wallet List
// @Schemes
// @Description Get Wallet List
// @Tags Wallet
//
// @Param   page         query     int        false  "int valid"
// @Param   pageSize         query     int        false  "int valid"
// @Param   sortField         query     string        false  "string valid"
// @Param   sortValue         query     string        false  "string valid"
//
//	@Success		200	{object}	response.WalletPagination
//
// @Security Bearer
//
// @Router /wallets [get]
func (p WalletController) GetListWallet(c *gin.Context) {
	query := CreatePagination(c)
	permissionInfo := response.Wallet{}
	p.svc.GetPaginationWallet(c, query.page, query.pageSize, query.search, query.sortField, query.sortValue, permissionInfo)
}

// @Summary Get Wallet By Id
// @Schemes
// @Description Get Wallet By Id
// @Tags Wallet
// @Param walletID  path int true "Wallet ID"
//
//	@Success		200	{object}	response.Wallet
//
// @Security Bearer
//
// @Router /wallets/{walletID} [get]
func (p WalletController) GetWalletById(c *gin.Context) {
	p.svc.GetWalletById(c)
}

// @Summary Update Wallet By Id
// @Schemes
// @Description Update Wallet By Id
// @Tags Wallet
// @Param walletID  path int true "Wallet ID"
// @Param request body request.UpdateWallet true "query params"
//
//	@Success		200	{object}	response.UpdateDataResponse
//
// @Security Bearer
//
// @Router /wallets/{walletID} [put]
func (p WalletController) UpdateWalletData(c *gin.Context) {
	p.svc.UpdateWallet(c)
}

// @Summary Delete Wallet By Id
// @Schemes
// @Description Delete Wallet By Id
// @Tags Wallet
// @Param walletID  path int true "Wallet ID"
//
//	@Success		200	{object}	response.DeleteDataResponse
//
// @Security Bearer
//
// @Router /wallets/{walletID} [delete]
func (p WalletController) DeleteWallet(c *gin.Context) {
	p.svc.DeleteWallet(c)
}
