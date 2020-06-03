package model

import (
	"time"
)

type Bill struct {
	Id        int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	UserId    int64     `json:"user_id" xorm:"not null comment('停车用户id') BIGINT(20) 'user_id'"`
	ParkId    int64     `json:"park_id" xorm:"not null comment('停车场id') BIGINT(20) 'park_id'"`
	Duration  time.Time `json:"duration" xorm:"not null comment('停车时长') TIME 'duration'"`
	Charge    int       `json:"charge" xorm:"not null comment('停车费用') INT(11) 'charge'"`
	Status    int       `json:"status" xorm:"not null comment('账单状态 0-未知 1-待付款 2-已付款') INT(11) 'status'"`
	CreatedAt time.Time `json:"created_at" xorm:"default 'current_timestamp()' created TIMESTAMP 'created_at'"`
	UpdatedAt time.Time `json:"updated_at" xorm:"default 'current_timestamp()' updated TIMESTAMP 'updated_at'"`
}

func (Bill) TableName() string {
	return "bill"
}
