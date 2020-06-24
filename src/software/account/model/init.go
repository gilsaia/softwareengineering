package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"software/common"
)

var DB *gorm.DB

func init() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", common.DbRoot, common.DbPassword, common.DbHost, common.DbName)
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		return
	}
	db.LogMode(true)
	DB = db
}

func NewDbConnection() (*gorm.DB, error) {
	if DB != nil {
		return DB, nil
	}
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", common.DbRoot, common.DbPassword, common.DbHost, common.DbName)
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	DB = db
	return db, nil
}
