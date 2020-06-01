package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	Id         int64     `json:"id" xorm:"pk BIGINT(20) 'id'"`
	Cellphone  string    `json:"cellphone" xorm:"comment('用户电话号码') CHAR(32) 'cellphone'"`
	Nickname   string    `json:"nickname" xorm:"comment('用户昵称') VARCHAR(4096) 'nickname'"`
	Password   string    `json:"password" xorm:"not null comment('用户密码') VARCHAR(4096) 'password'"`
	BindCode   string    `json:"bind_code" xorm:"not null comment('用户验证码') CHAR(8) 'bind_code'"`
	Permission int       `json:"permission" xorm:"default 0 comment('0-用户1-管理员留扩展以备后用') INT(11) 'permission'"`
	Status     int       `json:"status" xorm:"default 0 comment('0-不清楚/废弃1-等待验证2-验证通过') INT(11) 'status'"`
	CreatedAt  time.Time `json:"created_at" xorm:"default 'current_timestamp()' created TIMESTAMP 'created_at'"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"default 'current_timestamp()' updated TIMESTAMP 'updated_at'"`
}

func (User) TableName() string {
	return "user"
}

func CreateUser(db *gorm.DB, user User) error {
	db = db.Create(&user)
	return db.Error
}

func UpdateUser(db *gorm.DB, cellphone string, user User) error {
	dbUser := User{}
	db = db.Where("cellphone = ?", cellphone).First(&dbUser)
	if db.Error != nil {
		return db.Error
	}
	user.Id = dbUser.Id
	db = db.Model(&user).Updates(user)
	return db.Error
}

func GetUser(db *gorm.DB, cellphone string) (User, error) {
	user := User{}
	db = db.Table("user").Where("cellphone = ?", cellphone).First(&user)
	return user, db.Error
}
