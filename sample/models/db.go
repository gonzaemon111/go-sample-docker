package models

// https://qiita.com/chan-p/items/cf3e007b82cc7fce2d81
// https://blog.daimori.com/entry/gin-gorm%25e3%2581%25a7%25e7%25b0%25a1%25e5%258d%2598%25e3%2581%25aaweb%25e3%2582%25a2%25e3%2583%2597%25e3%2583%25aa%25e3%2582%2592%25e4%25bd%259c%25e3%2581%25a3%25e3%2581%25a6%25e3%2581%25bf%25e3%2581%259f

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