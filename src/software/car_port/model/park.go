package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Park struct {
	Id        int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	Name      string    `json:"name" xorm:"comment('停车场名字') VARCHAR(512) 'name'"`
	Longitude float64   `json:"longitude" xorm:"default 0 comment('停车场经度') DOUBLE 'longitude'"`
	Latitude  float64   `json:"latitude" xorm:"default 0 comment('停车场纬度') DOUBLE 'latitude'"`
	Charge    int       `json:"charge" xorm:"default 0 comment('停车费用 按小时计费') INT(11) 'charge'"`
	CreatedAt time.Time `json:"created_at" xorm:"default 'current_timestamp()' created TIMESTAMP 'created_at'"`
	UpdatedAt time.Time `json:"updated_at" xorm:"default 'current_timestamp()' updated TIMESTAMP 'updated_at'"`
}

func (Park) TableName() string {
	return "park"
}

func CreatePark(db *gorm.DB, park Park) error {
	db = db.Create(&park)
	return db.Error
}

func UpdatePark(db *gorm.DB, park Park) error {
	db = db.Model(&park).Updates(park)
	return db.Error
}

func GetPark(db *gorm.DB, parkId int64) (Park, error) {
	park := Park{}
	db = db.Where("id = ?", parkId).First(&park)
	return park, db.Error
}

func MGetPark(db *gorm.DB, offset int32, count int32) ([]Park, int32, error) {
	var parks []Park
	var tableCount int32
	db = db.Table("park").Offset(offset).Limit(count).Find(&parks)
	if db.Error != nil {
		return nil, 0, db.Error
	}
	db = db.Table("park").Offset(-1).Count(&tableCount)
	return parks, tableCount, db.Error
}

func MetaInfo(db *gorm.DB) ([]Park, error) {
	var parks []Park
	db = db.Table("park").Find(&parks)
	return parks, db.Error
}

func ParkNum(db *gorm.DB) (int32, error) {
	var tableCount int32
	db = db.Table("park").Count(&tableCount)
	return tableCount, db.Error
}
