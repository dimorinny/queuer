package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	driverName   = "sqlite3"
	databasePath = "queuer.db"
	maxIdleConns = 10
	maxOpenConns = 100
)

var (
	Db gorm.DB
)

func Init() {
	var (
		err error
	)

	Db, err = gorm.Open(driverName, databasePath)

	if err != nil {
		log.Fatal(err)
	}

	Db.LogMode(true)
	Db.DB()
	Db.DB().Ping()
	Db.DB().SetMaxIdleConns(maxIdleConns)
	Db.DB().SetMaxOpenConns(maxOpenConns)
}

func Migrate() {
	Db.AutoMigrate(&Queue{})
	Db.AutoMigrate(&Member{})
	Db.AutoMigrate(&User{})
	Db.AutoMigrate(&History{})
}

func Create() {
	Db.CreateTable(&Queue{})
	Db.CreateTable(&Member{})
	Db.CreateTable(&User{})
	Db.CreateTable(&History{})
}
