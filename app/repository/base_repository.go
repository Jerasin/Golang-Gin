package repository

import (
	"bytes"
	"fmt"
	"math"
	"strings"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/pkg"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Options struct {
	Query     interface{}
	QueryArgs []interface{}
	Join      string
	JoinArgs  []interface{}
	Select    string
	Preload   string
}

type BaseRepositoryInterface interface {
	ClientDb() *gorm.DB
	Pagination(p PaginationModel, query interface{}, args ...interface{}) (result interface{}, Error error)
	Save(tx *gorm.DB, model interface{}) error
	IsExits(tx *gorm.DB, model interface{}, query interface{}, args ...interface{}) error
	FindOne(tx *gorm.DB, model interface{}, query interface{}, args ...interface{}) error
	Find(tx *gorm.DB, model interface{}, query interface{}, args ...interface{}) error
	Update(tx *gorm.DB, id int, model interface{}, update interface{}) error
	TotalPage(model interface{}, pageSize int) (int64, error)
	Delete(model interface{}, id int) error
	FindOneV2(tx *gorm.DB, model interface{}, options Options) error
}

type BaseRepository struct {
	db *gorm.DB
}

type PaginationModel struct {
	Limit     int
	Offset    int
	Search    string
	SortField string
	SortValue string
	Field     map[string]interface{}
	Dest      interface{}
}

func BaseRepositoryInit(db *gorm.DB) *BaseRepository {

	return &BaseRepository{
		db: db,
	}
}

func getField(field map[string]interface{}) string {
	b := new(bytes.Buffer)
	index := 0
	for key := range field {
		// fmt.Println("key", key)
		if index > 0 {
			fmt.Fprintf(b, ",%s", strings.ToLower(key))
		} else {
			fmt.Fprintf(b, "%s", strings.ToLower(key))
		}

		index += 1
	}
	return b.String()

}

func (b BaseRepository) ClientDb() *gorm.DB {
	return b.db
}

func (b BaseRepository) Pagination(p PaginationModel, query interface{}, args ...interface{}) (result interface{}, Error error) {
	var err error
	order := fmt.Sprintf("%s %s", p.SortField, strings.ToUpper(p.SortValue))
	fields := getField(p.Field)
	var db *gorm.DB

	// if fields == "" {
	// 	db = b.db.Order(order).Offset(p.Offset).Limit(p.Limit).Find(&p.Dest)
	// } else {
	// 	db = b.db.Select(fields).Order(order).Offset(p.Offset).Limit(p.Limit).Find(&p.Dest)
	// }
	if fields != "" {
		db = b.db.Select(fields)
	}

	if query != nil {
		db = b.db.Where(query, args...)
	}

	db = db.Order(order).Offset(p.Offset).Limit(p.Limit).Find(&p.Dest)

	if db.Error != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return p.Dest, nil
}

func (b BaseRepository) Save(tx *gorm.DB, model interface{}) error {
	db := b.db

	if tx != nil {
		db = tx
	}

	var err = db.Save(model).Error
	if err != nil {
		log.Error("Got an error when save Error: ", err)
		return err
	}
	return nil
}

func (b BaseRepository) IsExits(tx *gorm.DB, model interface{}, query interface{}, args ...interface{}) error {
	db := b.db

	if tx != nil {
		db = tx
	}

	var err error
	if query != nil {
		err = db.Where(query, args).First(model).Error
	} else {
		err = db.First(model).Error
	}

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}

		log.Error("Got an error when findOne Error: ", err)
		return err
	}

	pkg.PanicException(constant.DataIsExit)
	return nil
}

func (b BaseRepository) FindOne(tx *gorm.DB, model interface{}, query interface{}, args ...interface{}) error {
	db := b.db

	if tx != nil {
		db = tx
	}
	var err error
	if query == nil || args == nil {
		log.Error("Got an error when findOne required query")
		pkg.PanicException(constant.RequiredQuery)
	}

	err = db.Where(query, args).First(model).Error

	if err != nil {
		log.Error("Got an error when findOne Error: ", err)
		return err
	}

	return nil
}

func (b BaseRepository) FindOneV2(tx *gorm.DB, model interface{}, options Options) error {
	db := b.db
	var str string
	if tx != nil {
		db = tx
	}
	var err error
	if options.Query == nil || options.QueryArgs == nil {
		log.Error("Got an error when findOne required query")
		pkg.PanicException(constant.RequiredQuery)
	}

	if options.Select != str {
		db = db.Select(options.Select)
	}

	if options.Join != str && options.Preload != str {
		db = db.Preload(options.Preload).Joins(options.Join)
	}

	err = db.Where(options.Query, options.QueryArgs...).First(model).Error

	if err != nil {
		log.Error("Got an error when findOne Error: ", err)
		return err
	}

	return nil
}

func (b BaseRepository) Find(tx *gorm.DB, model interface{}, query interface{}, args ...interface{}) error {
	db := b.db

	if tx != nil {
		db = tx
	}
	var err error
	if query == nil || args == nil {
		log.Error("Got an error when findOne required query")
		pkg.PanicException(constant.RequiredQuery)
	}

	fmt.Println("args", args)

	err = db.Where(query, args...).Find(model).Error

	if err != nil {
		log.Error("Got an error when find Error: ", err)
		return err
	}

	return nil
}

func (b BaseRepository) Update(tx *gorm.DB, id int, model interface{}, update interface{}) error {
	db := b.db

	if tx != nil {
		db = tx
	}
	var err = db.Model(model).Where(id).Updates(update).Error
	if err != nil {
		log.Error("Got an error when save user. Error: ", err)
		return err
	}
	return nil
}

func (b BaseRepository) TotalPage(model interface{}, pageSize int) (int64, error) {
	var count int64
	err := b.db.Model(model).Count(&count).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return count, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(pageSize)))
	return totalPage, err
}

func (b BaseRepository) Delete(model interface{}, id int) error {
	err := b.db.Delete(model, id).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return err
	}
	return nil
}
