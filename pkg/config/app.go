package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func ConnectDB() {
	d, err := gorm.Open("mysql", "root:root@123@tcp(localhost)/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	log.Println("db connected")
	db = d
}

func GetDB() *gorm.DB {
	return db
}
