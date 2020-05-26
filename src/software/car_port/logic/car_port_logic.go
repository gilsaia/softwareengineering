package logic

import (
	"context"
	"software/car_port/common"
	"software/car_port/model"
	"software/car_port/pb_gen"
)

type CarPortLogic struct {
	ctx context.Context
}

func NewCarPortLogic(ctx context.Context) CarPortLogic {
	return CarPortLogic{ctx: ctx}
}

func (logic CarPortLogic) CreateCarPort(port *pb_gen.CarPort) common.BgErr {
	if port.Id < 0 {
		return common.ParamErr
	}
	db, err := model.NewDbConnection()
	if err != nil {
		return common.DbErr
	}
	defer db.Close()
	carPort := model.CarPort{
		Id:       port.Id,
		State:    port.State,
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
	defer db.Close()
	carPort := model.CarPort{
		Id:       port.Id,
		State:    port.State,
		DrawNode: port.DrawNode,
	}
	err = model.UpdateCarPort(db, carPort.Id, carPort)
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
	defer db.Close()
	carPort, err := model.GetCarPort(db, portId)
	if err != nil {
		return nil, common.DbErr
	}
	port := &pb_gen.CarPort{
		Id:       carPort.Id,
		State:    carPort.State,
		DrawNode: carPort.DrawNode,
	}
	return port, common.Success
}

func (logic CarPortLogic) MGetCarPort(count int32, offset int32) (map[int64]*pb_gen.CarPort, bool, int32, common.BgErr) {
	db, err := model.NewDbConnection()
	if err != nil {
		return nil, false, 0, common.DbErr
	}
	defer db.Close()
	carports, hasMore, nextOffset, err := model.MGetCarPort(db, offset, count)
	if err != nil {
		return nil, false, 0, common.CustomErr(common.DbErr, err)
	}
	carPortMap := map[int64]*pb_gen.CarPort{}
	for _, value := range carports {
		carPortMap[value.Id] = packCarPort(value)
	}
	return carPortMap, hasMore, nextOffset, common.Success
}

func packCarPort(port model.CarPort) *pb_gen.CarPort {
	return &pb_gen.CarPort{
		Id:       port.Id,
		State:    port.State,
		DrawNode: port.DrawNode,
	}
}
