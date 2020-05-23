package model

import (
	"github.com/jinzhu/gorm"
	"software/car_port/common"
	"time"
)

type CarPort struct {
	Id        int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	State     int       `json:"state" xorm:"comment('车位状态 0-未定义 1-可用 2-在用 3-废弃') INT(11) 'state'"`
	DrawNode  int       `json:"draw_node" xorm:"comment('绘图结点') INT(11) 'draw_node'"`
	UserId    int64     `json:"user_id" xorm:"comment('当车位被使用时，表示使用者userid') BIGINT(20) 'user_id'"`
	CreatedAt time.Time `json:"created_at" xorm:"default 'current_timestamp()' created TIMESTAMP 'created_at'"`
	UpdatedAt time.Time `json:"updated_at" xorm:"default 'current_timestamp()' updated TIMESTAMP 'updated_at'"`
}

func (CarPort) TableName() string {
	return "car_port"
}

func CreateCarPort(db *gorm.DB) common.BgErr {
	db
}
