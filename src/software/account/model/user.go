package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	Id        int64     `json:"id" xorm:"pk BIGINT(20) 'id'"`
	Cellphone string    `json:"cellphone" xorm:"comment('用户电话号码') CHAR(32) 'cellphone'"`
	Nickname  string    `json:"nickname" xorm:"comment('用户昵称') VARCHAR(4096) 'nickname'"`
	Password  string    `json:"password" xorm:"not null comment('用户密码') VARCHAR(4096) 'password'"`
	BindCode  string    `json:"bind_code" xorm:"not null comment('用户验证码') CHAR(8) 'bind_code'"`
	CreatedAt time.Time `json:"created_at" xorm:"default 'current_timestamp()' created TIMESTAMP 'created_at'"`
	UpdatedAt time.Time `json:"updated_at" xorm:"default 'current_timestamp()' updated TIMESTAMP 'updated_at'"`
}

func (User) TableName() string {
	return "user"
}

func CreateUser(db *gorm.DB, user User) error {
	db = db.Create(&user)
	return db.Error
}

func UpdateUser(db *gorm.DB, cellphone string, user User) error {
	db = db.Where("cellphone = ?", cellphone).Update(&user)
	return db.Error
}

func GetUser(db *gorm.DB, cellphone string) (User, error) {
	user := User{}
	db = db.Table("user").Where("cellphone = ?", cellphone).First(&user)
	return user, db.Error
}
