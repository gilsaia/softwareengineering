package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Bill struct {
	Id        int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	UserId    int64     `json:"user_id" xorm:"not null comment('停车用户id') BIGINT(20) 'user_id'"`
	ParkId    int64     `json:"park_id" xorm:"not null comment('停车场id') BIGINT(20) 'park_id'"`
	Duration  string    `json:"duration" xorm:"not null comment('停车时长') CHAR(255) 'duration'"`
	Charge    int       `json:"charge" xorm:"not null comment('停车费用') INT(11) 'charge'"`
	Status    int       `json:"status" xorm:"not null comment('账单状态 0-未知 1-待付款 2-已付款') INT(11) 'status'"`
	CreatedAt time.Time `json:"created_at" xorm:"default 'current_timestamp()' created TIMESTAMP 'created_at'"`
	UpdatedAt time.Time `json:"updated_at" xorm:"default 'current_timestamp()' updated TIMESTAMP 'updated_at'"`
}

func (Bill) TableName() string {
	return "bill"
}

func CreateBill(db *gorm.DB, bill Bill) error {
	db = db.Create(&bill)
	return db.Error
}

func UpdateBill(db *gorm.DB, bill Bill) error {
	db = db.Model(&bill).Updates(bill)
	return db.Error
}

func GetBill(db *gorm.DB, billId int64) (Bill, error) {
	bill := Bill{}
	db = db.Where("id = ?", billId).First(&bill)
	return bill, db.Error
}

func MGetBill(db *gorm.DB, offset int32, count int32) ([]Bill, int32, error) {
	var bills []Bill
	var tableCount int32
	db = db.Table("bill").Offset(offset).Limit(count).Find(&bills)
	if db.Error != nil {
		return nil, 0, db.Error
	}
	db = db.Table("bill").Offset(-1).Count(&tableCount)
	return bills, tableCount, db.Error
}

func GetBillByUserId(db *gorm.DB, userId int64) ([]Bill, error) {
	var bills []Bill
	db = db.Table("bill").Where("user_id = ? and status = ?", userId, 1).Find(&bills)
	return bills, db.Error
}
