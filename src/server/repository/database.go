package repository

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type dataCenter interface {
	connectDatabase()
	getAll(any) *gorm.DB
	getAllByPagination(any, int, int, any, ...string) *gorm.DB
	getRowCount(string, *int64) *gorm.DB
	getById(any, any, ...string) *gorm.DB
	insertOne(any) *gorm.DB
	insertAll(any) *gorm.DB
	deleteOneById(any) *gorm.DB
	getAllByPaginationWithCondition(any, int, int, any, any, ...string) *gorm.DB
	getByIdWithCondition(any, string, any, ...string) *gorm.DB
	getOneByStructCondition(any, any) *gorm.DB
	getAllByStructCondition(any, any, any) *gorm.DB

	updateModelByMap(map[string]any, any) *gorm.DB
}

type seapDataCenter struct {
	db *gorm.DB
}

var dc dataCenter

func Init() {
	if dc != nil {
		return
	}
	dc = &seapDataCenter{}
	dc.connectDatabase()
}

func (d *seapDataCenter) connectDatabase() {
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

func (d *seapDataCenter) getAll(dest any) *gorm.DB {
	return d.db.Find(dest)
	//return d.db.Model(&dao.Member{}).Preload("Role").Find(dest)
}

func (d *seapDataCenter) getAllByPagination(dest any, offset, limit int, model any, preloads ...string) *gorm.DB {
	preloadedDb := d.db
	for _, s := range preloads {
		preloadedDb = preloadedDb.Preload(s)
	}
	return preloadedDb.Model(model).Limit(limit).Offset(offset).Find(dest)
}

func (d *seapDataCenter) getAllByPaginationWithCondition(dest any, offset, limit int, condition any, model any, preloads ...string) *gorm.DB {
	preloadedDb := d.db
	for _, s := range preloads {
		preloadedDb = preloadedDb.Preload(s)
	}
	return preloadedDb.Model(model).Where(condition).Limit(limit).Offset(offset).Find(dest)
}

func (d *seapDataCenter) getRowCount(tableName string, count *int64) *gorm.DB {
	return d.db.Table(tableName).Count(count)
}

func (d *seapDataCenter) getById(dest any, model any, preloads ...string) *gorm.DB {
	preloadedDb := d.db
	for _, s := range preloads {
		preloadedDb = preloadedDb.Preload(s)
	}
	return preloadedDb.Model(model).Where(dest).First(dest)
}

func (d *seapDataCenter) insertOne(dest any) *gorm.DB {
	return d.db.Create(dest)
}

func (d *seapDataCenter) insertAll(dest any) *gorm.DB {
	return d.insertOne(dest)
}

func (d *seapDataCenter) deleteOneById(dest any) *gorm.DB {
	return d.db.Delete(dest)
}

func (d *seapDataCenter) getByIdWithCondition(dest any, username string, model any, preloads ...string) *gorm.DB {
	preloadedDb := d.db

	preloadedDb = preloadedDb.Preload("DutiesWithSubmission", "username = (?)", username).Preload("DutiesWithSubmission.Duty_")

	return preloadedDb.Model(model).Where(dest).First(dest)
}

func (d *seapDataCenter) getOneByStructCondition(dest any, condition any) *gorm.DB {
	return d.db.Where(condition).First(dest)
}

func (d *seapDataCenter) getAllByStructCondition(dest any, condition any, model any) *gorm.DB {
	//preloadedDb := d.db
	//for _, s := range preloads {
	//	preloadedDb = preloadedDb.Preload(s)
	//}
	return d.db.Model(model).Where(condition).Find(dest)
}

func (d *seapDataCenter) updateModelByMap(fields map[string]any, model any) *gorm.DB {
	return d.db.Model(model).Updates(fields)
}
