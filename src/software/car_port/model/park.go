package model

import (
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
