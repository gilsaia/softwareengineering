package main

import (
	"context"
	"software/car_port/common"
	"software/car_port/logic"
	"software/car_port/pb_gen"
)

type carPortServer struct {
}

func newCarPortServer() pb_gen.CarPortServiceServer {
	return new(carPortServer)
}

func (s *carPortServer) AddCarPort(ctx context.Context, req *pb_gen.ReqAddCarPort) (*pb_gen.RespAddCarPort, error) {
	carPortLogic := logic.NewCarPortLogic(ctx)
	resp := &pb_gen.RespAddCarPort{}
	logicErr := carPortLogic.CreateCarPort(req.Port)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	return resp, nil
}

func (s *carPortServer) UpdateCarPort(ctx context.Context, req *pb_gen.ReqUpdateCarPort) (*pb_gen.RespUpdateCarPort, error) {
	carPortLogic := logic.NewCarPortLogic(ctx)
	resp := &pb_gen.RespUpdateCarPort{}
	logicErr := carPortLogic.UpdateCarPort(req.Port)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	return resp, nil
}

func (s *carPortServer) GetCarPort(ctx context.Context, req *pb_gen.ReqGetCarPort) (*pb_gen.RespGetCarPort, error) {
	carPortLogic := logic.NewCarPortLogic(ctx)
	resp := &pb_gen.RespGetCarPort{}
	carPort, logicErr := carPortLogic.GetCarPort(req.CarPortId)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	resp.Port = carPort
	return resp, nil
}

func (s *carPortServer) MGetCarPort(ctx context.Context, req *pb_gen.ReqMGetCarPort) (*pb_gen.RespMGetCarPort, error) {
	carPortLogic := logic.NewCarPortLogic(ctx)
	resp := &pb_gen.RespMGetCarPort{}
	carPorts, hasMore, nextOffset, logicErr := carPortLogic.MGetCarPort(req.Count, req.Offset)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	resp.Ports = carPorts
	resp.HasMore = hasMore
	resp.NextOffset = nextOffset
	return resp, nil
}
