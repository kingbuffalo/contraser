package gameutil

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"buffalo/king/common/startcfg"
)

var db *gorm.DB

func init() {
	var err error
	mysqlIp := startcfg.GetMysqlIp()
	db, err = gorm.Open("mysql", mysqlIp)
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func CloseMysql() error {
	return db.Close()
}
