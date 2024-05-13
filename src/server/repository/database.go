package repository

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DataCenter interface {
	ConnectDatabase()
	GetAll(any) *gorm.DB
	GetAllByPagination(any, int, int, any, ...string) *gorm.DB
	GetRowCount(string, *int64) *gorm.DB
	GetById(any, any, ...string) *gorm.DB
	InsertOne(any) *gorm.DB
	InsertAll(any) *gorm.DB
	DeleteOneById(any) *gorm.DB
	GetAllByPaginationWithCondition(any, int, int, any, any, ...string) *gorm.DB
	GetByIdWithCondition(any, string, any, ...string) *gorm.DB
	GetOneByStructCondition(any, any) *gorm.DB
	GetAllByStructCondition(any, any, any) *gorm.DB

	UpdateModelByMap(map[string]any, any) *gorm.DB
}

type seapDataCenter struct {
	db *gorm.DB
}

var dc DataCenter

func Init() {
	if dc != nil {
		return
	}
	dc = &seapDataCenter{}
	dc.ConnectDatabase()
}

func NewDataCenter() DataCenter {
	return &seapDataCenter{}
}

func InitializeDataCenter(ds DataCenter) {
	dc = ds
}

func (d *seapDataCenter) ConnectDatabase() {
	dbDriver := os.Getenv("DB_DRIVER")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	DB_URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	d.db, err = gorm.Open(mysql.Open(DB_URL), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to database ", dbDriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database", dbDriver)
	}
}

func (d *seapDataCenter) GetAll(dest any) *gorm.DB {
	return d.db.Find(dest)
	//return d.db.Model(&dao.Member{}).Preload("Role").Find(dest)
}

func (d *seapDataCenter) GetAllByPagination(dest any, offset, limit int, model any, preloads ...string) *gorm.DB {
	preloadedDb := d.db
	for _, s := range preloads {
		preloadedDb = preloadedDb.Preload(s)
	}
	return preloadedDb.Model(model).Limit(limit).Offset(offset).Find(dest)
}

func (d *seapDataCenter) GetAllByPaginationWithCondition(dest any, offset, limit int, condition any, model any, preloads ...string) *gorm.DB {
	preloadedDb := d.db
	for _, s := range preloads {
		preloadedDb = preloadedDb.Preload(s)
	}
	return preloadedDb.Model(model).Where(condition).Limit(limit).Offset(offset).Find(dest)
}

func (d *seapDataCenter) GetRowCount(tableName string, count *int64) *gorm.DB {
	return d.db.Table(tableName).Count(count)
}

func (d *seapDataCenter) GetById(dest any, model any, preloads ...string) *gorm.DB {
	preloadedDb := d.db
	for _, s := range preloads {
		preloadedDb = preloadedDb.Preload(s)
	}
	return preloadedDb.Model(model).Where(dest).First(dest)
}

func (d *seapDataCenter) InsertOne(dest any) *gorm.DB {
	return d.db.Create(dest)
}

func (d *seapDataCenter) InsertAll(dest any) *gorm.DB {
	return d.InsertOne(dest)
}

func (d *seapDataCenter) DeleteOneById(dest any) *gorm.DB {
	return d.db.Delete(dest)
}

func (d *seapDataCenter) GetByIdWithCondition(dest any, username string, model any, preloads ...string) *gorm.DB {
	preloadedDb := d.db

	preloadedDb = preloadedDb.Preload("DutiesWithSubmission", "username = (?)", username).Preload("DutiesWithSubmission.Duty_")

	return preloadedDb.Model(model).Where(dest).First(dest)
}

func (d *seapDataCenter) GetOneByStructCondition(dest any, condition any) *gorm.DB {
	return d.db.Where(condition).First(dest)
}

func (d *seapDataCenter) GetAllByStructCondition(dest any, condition any, model any) *gorm.DB {
	//preloadedDb := d.db
	//for _, s := range preloads {
	//	preloadedDb = preloadedDb.Preload(s)
	//}
	return d.db.Model(model).Where(condition).Find(dest)
}

func (d *seapDataCenter) UpdateModelByMap(fields map[string]any, model any) *gorm.DB {
	return d.db.Model(model).Updates(fields)
}
