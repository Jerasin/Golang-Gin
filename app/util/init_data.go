package util

import (
	"fmt"

	"github.com/Jerasin/app/model"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type InitDataClient struct {
	db *gorm.DB
}

func InitDataClientInit(db *gorm.DB) *InitDataClient {
	return &InitDataClient{
		db: db,
	}
}

func (i InitDataClient) InitPermissionInfo() []model.PermissionInfo {
	var err error

	data := ReadFile("app/default_data/permission_infos.json")

	var newPermissionInfoList []model.PermissionInfo
	var permissionInfoNameList []string
	for _, item := range data.([]map[string]any) {
		name, ok := item["name"].(string)

		if !ok {
			fmt.Println("err", ok)
		}
		fmt.Println("name", name)

		newPermissionInfo := model.PermissionInfo{
			Name: name,
		}

		newPermissionInfoList = append(newPermissionInfoList, newPermissionInfo)
		permissionInfoNameList = append(permissionInfoNameList, name)
	}

	var permissionInfoList []model.PermissionInfo
	fmt.Printf("newPermissionInfoList = %v Type = %T \n", newPermissionInfoList, newPermissionInfoList)
	fmt.Printf("permissionInfoList = %v Type = %T \n", permissionInfoList, permissionInfoList)

	err = i.db.Where("name IN ?", permissionInfoNameList).Find(&permissionInfoList).Error
	if err != nil {
		log.Error(err)
		panic("Find Error")
	}

	if len(newPermissionInfoList) != len(permissionInfoList) {
		err = i.db.Create(&newPermissionInfoList).Error
		if err != nil {
			log.Error(err)
			panic("Find Error")
		}

		fmt.Printf("newPermissionInfoList = %v", newPermissionInfoList)

		return newPermissionInfoList
	} else {
		return permissionInfoList
	}

}

func (i InitDataClient) InitRoleInfo(permissionInfos []model.PermissionInfo) []model.RoleInfo {
	var err error

	data := ReadFile("app/default_data/role_infos.json")

	var newRoleInfoList []model.RoleInfo
	var newRoleInfoNameList []string
	for _, item := range data.([]map[string]any) {
		name, ok := item["name"].(string)

		if !ok {
			fmt.Println("err", ok)
		}
		fmt.Println("name", name)

		newRoleInfo := model.RoleInfo{
			Name:            name,
			PermissionInfos: permissionInfos,
		}
		newRoleInfoList = append(newRoleInfoList, newRoleInfo)
		newRoleInfoNameList = append(newRoleInfoNameList, name)
	}

	var roleInfoList []model.RoleInfo
	fmt.Printf("newRoleInfoList = %v Type = %T \n", newRoleInfoList, newRoleInfoList)
	fmt.Printf("roleInfoList = %v Type = %T \n", roleInfoList, roleInfoList)

	err = i.db.Where("name IN ?", newRoleInfoNameList).Find(&roleInfoList).Error
	if err != nil {
		log.Error(err)
		panic("Find Error")
	}

	if len(newRoleInfoList) != len(roleInfoList) {
		err = i.db.Create(&newRoleInfoList).Error
		if err != nil {
			log.Error(err)
			panic("Find Error")
		}

		return newRoleInfoList
	} else {
		return roleInfoList
	}

}

func (i InitDataClient) InitUser() []model.User {
	var err error
	data := ReadFile("app/default_data/users.json")
	var newUserList []model.User
	var newUserNameList []string
	var newEmailList []string
	for _, item := range data.([]map[string]any) {
		var (
			username        string
			password        string
			fullname        string
			email           string
			avatar          string
			roleInfoId      uint
			roleInfoFloatId float64
		)
		var ok bool

		username, ok = item["username"].(string)
		if !ok {
			fmt.Println("Error: username is invalid")
			panic("username error")
		}

		password, ok = item["password"].(string)
		if !ok {
			fmt.Println("err", ok)
			panic("password error")
		}

		fullname, ok = item["fullname"].(string)
		if !ok {
			fmt.Println("err", ok)
			panic("fullname error")
		}

		email, ok = item["email"].(string)
		if !ok {
			fmt.Println("err", ok)
			panic("email error")
		}
		avatar, ok = item["avatar"].(string)
		if !ok {
			fmt.Println("err", ok)
			panic("avatar error")
		}
		roleInfoFloatId, ok = item["roleInfoId"].(float64)
		if ok {
			roleInfoId = uint(roleInfoFloatId)

		} else {
			fmt.Println("err", ok)
			panic("roleInfoId error")
		}

		fmt.Println("username", username)

		hash, _ := bcrypt.GenerateFromPassword([]byte(password), 15)
		user := model.User{
			Username:   username,
			Password:   string(hash),
			Fullname:   fullname,
			Email:      email,
			Avatar:     avatar,
			RoleInfoID: roleInfoId,
		}

		newUserNameList = append(newUserNameList, username)
		newUserList = append(newUserList, user)
		newEmailList = append(newEmailList, email)
	}

	var users []model.User

	err = i.db.Where("username IN ? OR email IN ?", newUserNameList, newEmailList).Find(&users).Error
	if err != nil {
		log.Error(err)
		panic("Find Error")
	}

	if len(newUserList) != len(users) {
		err = i.db.Create(&newUserList).Error
		if err != nil {
			log.Error(err)
			panic("Find Error")
		}

		return newUserList
	} else {
		return users
	}
}

func (i InitDataClient) InitProductCategory() []model.ProductCategory {
	var err error
	data := ReadFile("app/default_data/product_categories.json")
	var newProductCategoryList []model.ProductCategory
	var newProductCategoryNameList []string

	for _, item := range data.([]map[string]any) {
		name, ok := item["name"].(string)

		if !ok {
			fmt.Println("err", ok)
		}
		fmt.Println("name", name)

		newProductCategory := model.ProductCategory{
			Name: name,
		}
		newProductCategoryList = append(newProductCategoryList, newProductCategory)
		newProductCategoryNameList = append(newProductCategoryNameList, name)
	}

	var productCategoryList []model.ProductCategory
	err = i.db.Where("name IN ?", newProductCategoryNameList).Find(&productCategoryList).Error
	if err != nil {
		log.Error(err)
		panic("Find Error")
	}

	if len(newProductCategoryList) != len(productCategoryList) {
		err = i.db.Create(&newProductCategoryList).Error
		if err != nil {
			log.Error(err)
			panic("Find Error")
		}

		return newProductCategoryList
	} else {
		return productCategoryList
	}

}

func (i InitDataClient) InitProduct() []model.Product {
	var err error
	data := ReadFile("app/default_data/products.json")
	var newProductList []model.Product
	var newProductNameList []string

	for _, item := range data.([]map[string]any) {
		var (
			name         string
			price        float64
			amount       float64
			categoryName string
			countAmount  int
		)
		var ok bool

		name, ok = item["name"].(string)
		if !ok {
			fmt.Println("err", ok)
			panic("name error")
		}

		price, ok = item["price"].(float64)
		if !ok {
			fmt.Println("err", ok)
			panic("price error")
		}

		amount, ok = item["amount"].(float64)
		if ok {
			countAmount = int(amount)
		} else {
			panic("amount error")
		}

		categoryName, ok = item["categoryName"].(string)
		if !ok {
			fmt.Println("err", ok)
			panic("categoryName error")
		}

		var productCategory model.ProductCategory
		err = i.db.Where("name = ?", categoryName).First(&productCategory).Error
		if err != nil {
			panic(err)
		}

		fmt.Println("productCategory ID", productCategory.ID)

		newProduct := model.Product{
			Name:              name,
			Price:             price,
			Amount:            countAmount,
			ProductCategoryID: productCategory.ID,
			SaleOpenDate:      nil,
		}
		newProductList = append(newProductList, newProduct)
		newProductNameList = append(newProductNameList, name)

		fmt.Printf("newProduct value = %v type = %T", newProduct, newProduct)
	}

	var productList []model.Product
	err = i.db.Where("name IN ?", newProductNameList).Find(&productList).Error
	if err != nil {
		log.Error(err)
		panic("Find Error")
	}

	if len(newProductList) != len(productList) {
		fmt.Printf("value = %v type = %T", newProductList, newProductList)
		err = i.db.Create(&newProductList).Error
		if err != nil {
			log.Error(err)
			panic("Create Error")
		}

		return newProductList
	} else {
		return productList
	}

}
