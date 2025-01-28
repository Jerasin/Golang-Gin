package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/model"
	"github.com/Jerasin/app/pkg"
	"github.com/Jerasin/app/repository"
	"github.com/Jerasin/app/request"
	"github.com/Jerasin/app/response"
	"github.com/Jerasin/app/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrderServiceModel struct {
	BaseRepository repository.BaseRepositoryInterface
}

type OrderServiceInterface interface {
	CreateOrder(c *gin.Context)
	GetPaginationOrder(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.Order)
}

func OrderServiceInit(baseRepo repository.BaseRepositoryInterface) *OrderServiceModel {
	return &OrderServiceModel{
		BaseRepository: baseRepo,
	}
}

// goroutine ใช้ใน transaction Database ไม่ได้
func (o OrderServiceModel) CreateOrder(c *gin.Context) {
	defer pkg.PanicHandler(c)

	o.BaseRepository.ClientDb().Transaction(func(tx *gorm.DB) error {
		var body request.OrderRequest
		var err error
		var productIDS []uint
		var products []model.Product
		var orderDetails []model.OrderDetail
		var totalPrice float64
		var totalAmount int
		currentDate := time.Now()
		// Validate Request Body
		err = c.ShouldBindJSON(&body)
		if err != nil {
			log.Error("error ShouldBindJSON", err)
			pkg.PanicException(constant.BadRequest)
		}

		for _, value := range body.Orders {
			productIDS = append(productIDS, uint(value.ProductId))
		}

		log.Debugf("productIDS = %+v\n", productIDS)

		err = o.BaseRepository.Find(tx, &products, "id IN ? AND ((sale_open_date >= ? AND ? <= sale_close_date) OR (sale_open_date IS NULL AND sale_close_date IS NULL))", productIDS, currentDate, currentDate)
		if err != nil {
			log.Error("Find", err)
			pkg.PanicException(constant.BadRequest)
		}

		fmt.Printf("products = %+v type = %T \n", products, products)
		if len(products) != len(productIDS) || len(products) == 0 {
			log.Error("Error", len(products) != len(productIDS) || len(products) == 0)
			pkg.PanicException(constant.DataNotFound)
		}

		for index, product := range products {

			if product.Amount <= 0 {
				pkg.PanicException(constant.BadRequest)
			}

			item, err := util.FindElementByCondition(body.Orders, func(o request.OrderItem) bool {
				return o.ProductId == int(product.ID)
			})
			if err != nil {
				pkg.PanicException(constant.BadRequest)
			}

			orderDetail := model.OrderDetail{
				ProductID: product.ID,
				Price:     product.Price,
				Amount:    item.Amount,
			}

			orderDetails = append(orderDetails, orderDetail)
			totalPrice += product.Price * float64(item.Amount)
			totalAmount += item.Amount

			products[index].Amount -= item.Amount
		}

		log.Debugf("products 2 = %+v\n", products)

		username, err := util.GetPayloadInToken(c, "username")
		if err != nil {
			pkg.PanicException(constant.BadRequest)
		}
		log.Debug("username", username)
		var user model.User
		walletID := body.WalletID
		options := repository.Options{
			Query:     "username = ? AND wallets.id = ?",
			QueryArgs: []interface{}{username, walletID},
			Join:      "LEFT JOIN wallets ON users.id = wallets.user_id ",
			Preload:   "Wallets",
		}

		o.BaseRepository.FindOneV2(tx, &user, options)
		fmt.Printf("user = %+v", user)

		walletCheck := false
		var wallet model.Wallet
		for _, userWallet := range user.Wallets {
			if userWallet.Value >= totalPrice {
				updateWallet := model.Wallet{
					Value: userWallet.Value - totalPrice,
				}
				err = o.BaseRepository.Update(tx, int(walletID), &wallet, &updateWallet)
				if err != nil {
					pkg.PanicException(constant.BadRequest)
				}

				walletCheck = true
			}
		}

		if !walletCheck {
			pkg.PanicException(constant.BadRequest)
		}

		order := model.Order{
			TotalPrice:  totalPrice,
			TotalAmount: totalAmount,
			CreatedBy:   user.ID,
			WalletID:    walletID,
		}
		err = o.BaseRepository.Save(tx, &order)
		if err != nil {
			log.Error(err)
			pkg.PanicException(constant.BadRequest)
		}
		fmt.Println("create Order Success")

		for i, value := range orderDetails {
			fmt.Println("value", value)
			orderDetails[i].OrderID = order.ID
		}

		err = o.BaseRepository.Save(tx, &orderDetails)
		fmt.Println(err)
		if err != nil {
			log.Error(err)
			pkg.PanicException(constant.BadRequest)
		}

		o.BaseRepository.Save(tx, &products)

		c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.CreateResponse()))
		return nil
	})

}

func (o OrderServiceModel) GetPaginationOrder(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.Order) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get list product")
	username, err := util.GetPayloadInToken(c, "username")

	if err != nil {
		log.Error("Happened error when GetPayloadInToken Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}
	offset := (page - 1) * pageSize
	limit := pageSize
	fields := DbHandleSelectField(field)
	// fields := structs.Map(field)

	var user model.User
	o.BaseRepository.FindOne(nil, &user, "username = ?", username)

	var orders []model.Order
	paginationModel := repository.PaginationModel{
		Limit:     limit,
		Offset:    offset,
		Search:    search,
		SortField: sortField,
		SortValue: sortValue,
		Field:     fields,
		Dest:      orders,
	}

	data, err := o.BaseRepository.Pagination(paginationModel, "created_by = ?", user.ID)

	if err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	totalPage, err := o.BaseRepository.TotalPage(&orders, pageSize)
	if err != nil {
		log.Error("Count Data Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	var res []response.Order

	pkg.ModelDump(&res, data)
	c.JSON(http.StatusOK, pkg.BuildPaginationResponse(constant.Success, res, totalPage, page, pageSize))
}
