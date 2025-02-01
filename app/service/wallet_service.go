package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/model"
	"github.com/Jerasin/app/pkg"
	"github.com/Jerasin/app/repository"
	"github.com/Jerasin/app/request"
	"github.com/Jerasin/app/response"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type WalletServiceInterface interface {
	CreateWallet(c *gin.Context)
	GetPaginationWallet(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.Wallet)
	GetWalletById(c *gin.Context)
	UpdateWallet(c *gin.Context)
	DeleteWallet(c *gin.Context)
}

type WalletServiceModel struct {
	BaseRepository repository.BaseRepositoryInterface
}

func WalletServiceInit(baseRepo repository.BaseRepositoryInterface) *WalletServiceModel {
	return &WalletServiceModel{
		BaseRepository: baseRepo,
	}
}

func (p WalletServiceModel) CreateWallet(c *gin.Context) {
	defer pkg.PanicHandler(c)
	var err error
	var body model.Wallet

	err = c.ShouldBindJSON(&body)
	if err != nil {
		pkg.PanicException(constant.BadRequest)
	}

	fmt.Printf("body = %+v", body)

	p.BaseRepository.ClientDb().Transaction(func(tx *gorm.DB) error {
		err := p.BaseRepository.Save(tx, &body)

		if err != nil {
			pkg.PanicException(constant.BadRequest)
		}

		c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.CreateResponse()))
		return nil
	})
}

func (p WalletServiceModel) GetPaginationWallet(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.Wallet) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get list wallet")

	log.Info("start to execute get all data user")
	offset := (page - 1) * pageSize
	limit := pageSize
	fields := DbHandleSelectField(field)
	fmt.Println("query", search)
	fmt.Println("fields", fields)
	var wallets []model.Wallet

	paginationModel := repository.PaginationModel{
		Limit:     limit,
		Offset:    offset,
		Search:    search,
		SortField: sortField,
		SortValue: sortValue,
		Field:     fields,
		Dest:      wallets,
	}
	data, err := p.BaseRepository.Pagination(paginationModel, nil)

	if err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	totalPage, err := p.BaseRepository.TotalPage(&wallets, pageSize)
	if err != nil {
		log.Error("Count Data Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	fmt.Println("data", data)

	var res []response.Wallet
	pkg.ModelDump(&res, data)

	c.JSON(http.StatusOK, pkg.BuildPaginationResponse(constant.Success, res, totalPage, page, pageSize))
}

func (p WalletServiceModel) GetWalletById(c *gin.Context) {
	fmt.Println("FindOneV2")
	defer pkg.PanicHandler(c)

	walletID, _ := strconv.Atoi(c.Param("walletID"))

	var wallet model.Wallet
	options := repository.Options{
		Query:     "wallets.id = ?",
		QueryArgs: []interface{}{walletID},
		Join:      "LEFT JOIN users ON users.id = wallets.user_id ",
		Preload:   "User",
	}
	err := p.BaseRepository.FindOneV2(nil, &wallet, options)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	fmt.Printf("wallet = %+v", wallet)

	var res response.WalletDetail
	pkg.ModelDump(&res, wallet)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, res))
}

func (p WalletServiceModel) UpdateWallet(c *gin.Context) {
	defer pkg.PanicHandler(c)

	p.BaseRepository.ClientDb().Transaction(func(tx *gorm.DB) error {
		var err error
		walletID, _ := strconv.Atoi(c.Param("walletID"))
		var request request.UpdateWallet

		err = c.ShouldBindJSON(&request)
		if err != nil {
			pkg.PanicException(constant.BadRequest)
		}

		var wallet model.Wallet
		err = p.BaseRepository.FindOne(tx, &wallet, "id = ?", walletID)
		if err != nil {
			log.Error("Happened error when get data from database. Error", err)
			pkg.PanicException(constant.DataNotFound)
		}

		err = p.BaseRepository.Update(tx, walletID, &wallet, &request)
		if err != nil {
			pkg.PanicException(constant.BadRequest)
		}

		c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.UpdateResponse()))
		return nil
	})

}

func (p WalletServiceModel) DeleteWallet(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data user by id")
	walletID, _ := strconv.Atoi(c.Param("walletID"))

	var wallet model.Wallet
	err := p.BaseRepository.Delete(&wallet, walletID)
	if err != nil {
		log.Error("Happened Error when try delete data user from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.DeleteResponse()))
}
