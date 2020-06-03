package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type CarPort struct {
	Id        int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	State     int       `json:"state" xorm:"default 0 comment('车位状态 0-未定义 1-可用 2-在用 3-废弃') INT(11) 'state'"`
	UserId    int64     `json:"user_id" xorm:"default 0 comment('当车位被使用时，表示使用者userid') BIGINT(20) 'user_id'"`
	Longitude float64   `json:"longitude" xorm:"default 0 comment('车位经度 x坐标') DOUBLE 'longitude'"`
	Latitude  float64   `json:"latitude" xorm:"default 0 comment('车位纬度 y坐标') DOUBLE 'latitude'"`
	Park      int64     `json:"park" xorm:"default 0 comment('所属停车场id') BIGINT(20) 'park'"`
	LastUsed  time.Time `json:"last_used" xorm:"comment('最近一次使用开始时间') TIMESTAMP 'last_used'"`
	CreatedAt time.Time `json:"created_at" xorm:"default 'current_timestamp()' created TIMESTAMP 'created_at'"`
	UpdatedAt time.Time `json:"updated_at" xorm:"default 'current_timestamp()' updated TIMESTAMP 'updated_at'"`
}

func (CarPort) TableName() string {
	return "car_port"
}

func CreateCarPort(db *gorm.DB, port CarPort) error {
	db = db.Create(&port)
	return db.Error
}

func UpdateCarPort(db *gorm.DB, port CarPort) error {
	db = db.Model(&port).Updates(port)
	return db.Error
}

func GetCarPort(db *gorm.DB, portId int64) (CarPort, error) {
	port := CarPort{}
	db = db.Where("id = ?", portId).First(&port)
	return port, db.Error
}

func MGetCarPort(db *gorm.DB, offset int32, count int32) ([]CarPort, int32, error) {
	var carPorts []CarPort
	var tableCount int32
	//var tableCount, nextOffset int32
	//hasMore := false
	db = db.Table("car_port").Offset(offset).Limit(count).Find(&carPorts)
	if db.Error != nil {
		return nil, 0, db.Error
	}
	db = db.Table("car_port").Offset(-1).Count(&tableCount)
	//if offset+count < tableCount {
	//	hasMore = true
	//	nextOffset = offset + count
	//} else {
	//	nextOffset = tableCount
	//}
	return carPorts, tableCount, db.Error
}
