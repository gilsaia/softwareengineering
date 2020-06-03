package logic

import (
	"context"
	"software/car_port/model"
	"software/car_port/pb_gen"
	"software/common"
)

type ClientLogic struct {
	ctx context.Context
}

func NewClientLogic(ctx context.Context) (*ClientLogic, common.BgErr) {
	if err := common.AuthPermission(ctx, common.PermissionVisitor); !err.Is(common.Success) {
		return nil, err
	}
	return &ClientLogic{ctx: ctx}, common.Success
}

func (logic ClientLogic) UserInfo(cellphone string) (nickname string, bgErr common.BgErr) {
	db, err := model.NewDbConnection()
	if err != nil {
		return "", common.CustomErr(common.DbErr, err)
	}
	user, err := model.GetUserByCellphone(db, cellphone)
	if err != nil {
		return "", common.CustomErr(common.DbErr, err)
	}
	return user.Nickname, common.Success
}

func (logic ClientLogic) MetaInfo() (count int32, parks []*pb_gen.Park, bgErr common.BgErr) {
	db, err := model.NewDbConnection()
	if err != nil {
		return 0, nil, common.CustomErr(common.DbErr, err)
	}
	count, err = model.ParkNum(db)
	if err != nil {
		return 0, nil, common.CustomErr(common.DbErr, err)
	}
	mParks, err := model.MetaInfo(db)
	if err != nil {
		return 0, nil, common.CustomErr(common.DbErr, err)
	}
	var parkIds []int64
	for _, value := range mParks {
		parkIds = append(parkIds, value.Id)
	}
	emptyPortMap, totalPortMap, err := model.GetParkCarPortNum(db, parkIds)
	if err != nil {
		return 0, nil, common.CustomErr(common.DbErr, err)
	}
	var parkList []*pb_gen.Park
	for _, value := range mParks {
		emptyPort := int32(0)
		totalPort := int32(0)
		if port, exist := emptyPortMap[value.Id]; exist {
			emptyPort = port
		}
		if port, exist := totalPortMap[value.Id]; exist {
			totalPort = port
		}
		parkList = append(parkList, packPark(value, emptyPort, totalPort))
	}
	return count, parkList, common.Success
}

func (logic ClientLogic) ParkInfo(parkId int64) (parkName string, carPorts []*pb_gen.CarPort, bgErr common.BgErr) {
	db, err := model.NewDbConnection()
	if err != nil {
		return "", nil, common.CustomErr(common.DbErr, err)
	}
	park, err := model.GetPark(db, parkId)
	if err != nil {
		return "", nil, common.CustomErr(common.DbErr, err)
	}
	ports, err := model.GetParkCarPort(db, parkId)
	if err != nil {
		return "", nil, common.CustomErr(common.DbErr, err)
	}
	for _, value := range ports {
		carPorts = append(carPorts, packCarPort(value))
	}
	return park.Name, carPorts, common.Success
}
