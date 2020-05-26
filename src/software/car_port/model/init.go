package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"software/car_port/common"
)

func NewDbConnection() (*gorm.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(mysql)/%s?charset=utf8&parseTime=True&loc=Local", common.DbRoot, common.DbPassword, common.DbName)
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return db, nil
}
