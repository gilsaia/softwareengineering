package logic

import (
	"context"
	"software/car_port/model"
	"software/car_port/pb_gen"
	"software/common"
)

type CarPortLogic struct {
	ctx context.Context
}

func NewCarPortLogic(ctx context.Context) (*CarPortLogic, common.BgErr) {
	if err := common.AuthPermission(ctx, common.PermissionAdmin); !err.Is(common.Success) {
		return nil, err
	}
	return &CarPortLogic{ctx: ctx}, common.Success
}

func (logic CarPortLogic) CreateCarPort(port *pb_gen.CarPort) common.BgErr {
	if port.Id < 0 {
		return common.ParamErr
	}
	db, err := model.NewDbConnection()
	if err != nil {
		return common.DbErr
	}
	carPort := model.CarPort{
		Id:       port.Id,
		State:    int32(port.State),
		DrawNode: port.DrawNode,
	}
	err = model.CreateCarPort(db, carPort)
	if err != nil {
		return common.DbErr
	}
	return common.Success
}

func (logic CarPortLogic) UpdateCarPort(port *pb_gen.CarPort) common.BgErr {
	if port.Id < 0 {
		return common.ParamErr
	}
	db, err := model.NewDbConnection()
	if err != nil {
		return common.DbErr
	}
	carPort := model.CarPort{
		Id:       port.Id,
		State:    int32(port.State),
		DrawNode: port.DrawNode,
	}
	err = model.UpdateCarPort(db, carPort)
	if err != nil {
		return common.DbErr
	}
	return common.Success
}

func (logic CarPortLogic) GetCarPort(portId int64) (*pb_gen.CarPort, common.BgErr) {
	if portId < 0 {
		return nil, common.ParamErr
	}
	db, err := model.NewDbConnection()
	if err != nil {
		return nil, common.DbErr
	}
	carPort, err := model.GetCarPort(db, portId)
	if err != nil {
		return nil, common.DbErr
	}
	port := packCarPort(carPort)
	return port, common.Success
}

func (logic CarPortLogic) MGetCarPort(count int32, num int32) ([]*pb_gen.CarPort, int32, common.BgErr) {
	db, err := model.NewDbConnection()
	if err != nil {
		return nil, 0, common.DbErr
	}
	offset := num * count
	carports, tableCount, err := model.MGetCarPort(db, offset, count)
	if err != nil {
		return nil, 0, common.CustomErr(common.DbErr, err)
	}
	var carPortList []*pb_gen.CarPort
	for _, value := range carports {
		carPortList = append(carPortList, packCarPort(value))
	}
	return carPortList, tableCount, common.Success
}

func packCarPort(port model.CarPort) *pb_gen.CarPort {
	return &pb_gen.CarPort{
		Id:       port.Id,
		State:    pb_gen.CarPortStatus(port.State),
		DrawNode: port.DrawNode,
	}
}
