package model

import (
	"github.com/jinzhu/gorm"
	"software/common"
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

type parkPort struct {
	Id  int64
	Num int32
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

func GetParkCarPortNum(db *gorm.DB, parkIds []int64) (map[int64]int32, map[int64]int32, error) {
	var totalParkPorts, emptyParkPorts []parkPort
	db = db.Table("car_port").Select("park as id,COUNT(id) as num").Where("(state=1 or state=2) AND park in (?)", parkIds).Group("park").Scan(&totalParkPorts)
	if db.Error != nil {
		return nil, nil, db.Error
	}
	db = db.Table("car_port").Select("park as id,COUNT(id) as num").Where("state=1 and park in (?)", parkIds).Group("park").Scan(&emptyParkPorts)
	if db.Error != nil {
		return nil, nil, db.Error
	}
	totalMap := map[int64]int32{}
	emptyMap := map[int64]int32{}
	for _, value := range totalParkPorts {
		totalMap[value.Id] = value.Num
	}
	for _, value := range emptyParkPorts {
		emptyMap[value.Id] = value.Num
	}
	return emptyMap, totalMap, db.Error
}

func GetParkCarPort(db *gorm.DB, parkId int64) ([]CarPort, error) {
	var carPorts []CarPort
	db = db.Table("car_port").Where("park = ?", parkId).Find(&carPorts)
	return carPorts, db.Error
}

func UserPark(db *gorm.DB, portId int64, userId int64) error {
	port := CarPort{}
	db = db.Where("id = ?", portId).First(&port)
	if db.Error != nil {
		return db.Error
	}
	if port.State == 3 || port.State == 0 {
		return common.AbandonErr
	}
	if port.UserId != 0 || port.State == 2 {
		return common.OccupyErr
	}
	port.State = 2
	port.UserId = userId
	port.LastUsed = time.Now()
	db = db.Model(&port).Updates(port)
	return db.Error
}

func UserPickUp(db *gorm.DB, portId int64, userId int64) (time.Duration, int64, error) {
	port := CarPort{}
	db = db.Where("id = ?", portId).First(&port)
	if db.Error != nil {
		return time.Since(time.Now()), 0, db.Error
	}
	if port.State != 2 || port.UserId != userId {
		return time.Since(time.Now()), 0, common.ParamErr
	}
	if port.LastUsed.After(time.Now()) {
		return time.Since(time.Now()), 0, common.ParamErr
	}
	duration := time.Now().Sub(port.LastUsed)
	parkId := port.Park
	port.State = 1
	port.UserId = 0
	db = db.Model(&port).Updates(port)
	return duration, parkId, db.Error
}
