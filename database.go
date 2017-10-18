package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Database() *gorm.DB {
	db, err := gorm.Open("mysql", "root:password@tcp(localhost:3306)/sos?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	return db
}