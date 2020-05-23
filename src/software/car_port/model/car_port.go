package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type CarPort struct {
	Id        int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	State     int32     `json:"state" xorm:"comment('车位状态 0-未定义 1-可用 2-在用 3-废弃') INT(11) 'state'"`
	DrawNode  int32     `json:"draw_node" xorm:"comment('绘图结点') INT(11) 'draw_node'"`
	UserId    int64     `json:"user_id" xorm:"comment('当车位被使用时，表示使用者userid') BIGINT(20) 'user_id'"`
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

func UpdateCarPort(db *gorm.DB, portId int64, port CarPort) error {
	db = db.Where("id = ?", portId).Save(&port)
	return db.Error
}

func GetCarPort(db *gorm.DB, portId int64) (CarPort, error) {
	port := CarPort{}
	db = db.Where("id = ?", portId).First(&port)
	return port, db.Error
}

func MGetCarPort(db *gorm.DB, offset int32, count int32) ([]CarPort, bool, int32, error) {
	var carPorts []CarPort
	var tableCount, nextOffset int32
	hasMore := false
	db = db.Table("car_port").Offset(offset).Count(count).Find(&carPorts)
	if db.Error != nil {
		return nil, false, 0, db.Error
	}
	db = db.Table("car_port").Count(&tableCount)
	if offset+count < tableCount {
		hasMore = true
		nextOffset = offset + count
	} else {
		nextOffset = tableCount
	}
	return carPorts, hasMore, nextOffset, db.Error
}
