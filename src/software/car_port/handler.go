package main

import (
	"context"
	"software/car_port/logic"
	"software/car_port/pb_gen"
	"software/common"
)

type carPortServer struct {
}

func newCarPortServer() pb_gen.CarPortServiceServer {
	return new(carPortServer)
}

func (s *carPortServer) AddCarPort(ctx context.Context, req *pb_gen.ReqAddCarPort) (*pb_gen.RespAddCarPort, error) {
	resp := &pb_gen.RespAddCarPort{}
	carPortLogic, logicErr := logic.NewCarPortLogic(ctx, common.PermissionAdmin)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	logicErr = carPortLogic.CreateCarPort(req.Port)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	return resp, nil
}

func (s *carPortServer) UpdateCarPort(ctx context.Context, req *pb_gen.ReqUpdateCarPort) (*pb_gen.RespUpdateCarPort, error) {
	resp := &pb_gen.RespUpdateCarPort{}
	carPortLogic, logicErr := logic.NewCarPortLogic(ctx, common.PermissionAdmin)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	logicErr = carPortLogic.UpdateCarPort(req.Port)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	return resp, nil
}

func (s *carPortServer) GetCarPort(ctx context.Context, req *pb_gen.ReqGetCarPort) (*pb_gen.RespGetCarPort, error) {
	resp := &pb_gen.RespGetCarPort{}
	carPortLogic, logicErr := logic.NewCarPortLogic(ctx, common.PermissionAdmin)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
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
	resp := &pb_gen.RespMGetCarPort{}
	carPortLogic, logicErr := logic.NewCarPortLogic(ctx, common.PermissionAdmin)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
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
