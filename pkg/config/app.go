package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
