package models

// https://qiita.com/chan-p/items/cf3e007b82cc7fce2d81

import (
	"github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
)

var db = NewDBConn()

func NewDBConn() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(go_sample_mysql:3306)/go_sample?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	//マイグレーション
	//db.AutoMigrate(&Book{})


	return db
}

func GetDBConn() *gorm.DB {
	return db
}

// func GetDBConfig() (string, string) {
// 	return config.Env("dialect"), config.Env("datasource")
// }