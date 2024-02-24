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
}

type SeapDataCenter struct {
	db *gorm.DB
}

func (d SeapDataCenter) ConnectDatabase() {
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
