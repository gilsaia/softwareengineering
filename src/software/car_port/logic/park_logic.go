package logic

import (
	"context"
	"software/car_port/model"
	"software/car_port/pb_gen"
	"software/common"
)

type ParkLogic struct {
	ctx context.Context
}

func NewParkLogic(ctx context.Context) (*ParkLogic, common.BgErr) {
	if err := common.AuthPermission(ctx, common.PermissionAdmin); !err.Is(common.Success) {
		return nil, err
	}
	return &ParkLogic{ctx: ctx}, common.Success
}

func (logic ParkLogic) CreatePark(park *pb_gen.Park) common.BgErr {
	if park.Id < 0 {
		return common.ParamErr
	}
	db, err := model.NewDbConnection()
	if err != nil {
		return common.CustomErr(common.DbErr, err)
	}
	mPark := model.Park{
		Id:        park.Id,
		Name:      park.Name,
		Latitude:  park.Latitude,
		Longitude: park.Longitude,
		Charge:    int(park.Charge),
	}
	err = model.CreatePark(db, mPark)
	if err != nil {
		return common.CustomErr(common.DbErr, err)
	}
	return common.Success
}

func (logic ParkLogic) UpdatePark(park *pb_gen.Park) common.BgErr {
	if park.Id < 0 {
		return common.ParamErr
	}
	db, err := model.NewDbConnection()
	if err != nil {
		return common.DbErr
	}
	mPark := model.Park{
		Id:        park.Id,
		Name:      park.Name,
		Latitude:  park.Latitude,
		Longitude: park.Longitude,
		Charge:    int(park.Charge),
	}
	err = model.UpdatePark(db, mPark)
	if err != nil {
		return common.DbErr
	}
	return common.Success
}

func (logic ParkLogic) GetPark(parkId int64) (*pb_gen.Park, common.BgErr) {
	if parkId < 0 {
		return nil, common.ParamErr
	}
	db, err := model.NewDbConnection()
	if err != nil {
		return nil, common.DbErr
	}
	mPark, err := model.GetPark(db, parkId)
	if err != nil {
		return nil, common.DbErr
	}
	parkIds := []int64{parkId}
	emptyPortMap, totalPortMap, err := model.GetParkCarPortNum(db, parkIds)
	if err != nil {
		return nil, common.CustomErr(common.DbErr, err)
	}
	emptyPort := int32(0)
	totalPort := int32(0)
	if value, exist := emptyPortMap[parkId]; exist {
		emptyPort = value
	}
	if value, exist := totalPortMap[parkId]; exist {
		totalPort = value
	}
	park := packPark(mPark, emptyPort, totalPort)
	return park, common.Success
}

func (logic ParkLogic) MGetPark(count int32, num int32) ([]*pb_gen.Park, int32, common.BgErr) {
	db, err := model.NewDbConnection()
	if err != nil {
		return nil, 0, common.DbErr
	}
	offset := num * count
	parks, tableCount, err := model.MGetPark(db, offset, count)
	if err != nil {
		return nil, 0, common.CustomErr(common.DbErr, err)
	}
	var parkIds []int64
	for _, value := range parks {
		parkIds = append(parkIds, value.Id)
	}
	emptyPortMap, totalPortMap, err := model.GetParkCarPortNum(db, parkIds)
	if err != nil {
		return nil, 0, common.CustomErr(common.DbErr, err)
	}
	var parkList []*pb_gen.Park
	for _, value := range parks {
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
	return parkList, tableCount, common.Success
}

func packPark(park model.Park, emptyPort int32, totalPort int32) *pb_gen.Park {
	return &pb_gen.Park{
		Id:         park.Id,
		Name:       park.Name,
		Longitude:  park.Longitude,
		Latitude:   park.Latitude,
		Charge:     int32(park.Charge),
		EmptyPorts: emptyPort,
		TotalPorts: totalPort,
	}
}
