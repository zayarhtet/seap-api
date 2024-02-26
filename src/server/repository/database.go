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
}

type seapDataCenter struct {
	db *gorm.DB
}

var dc dataCenter

func Init() {
	if dc != nil { return }
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
}
