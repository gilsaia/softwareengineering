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

func (s *carPortServer) UpdateUser(ctx context.Context, req *pb_gen.ReqUpdateUser) (*pb_gen.RespUpdateUser, error) {
	resp := &pb_gen.RespUpdateUser{}
	userLogic, logicErr := logic.NewUserLogic(ctx)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	logicErr = userLogic.UpdateUser(req.User, req.Password)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	return resp, nil
}

func (s *carPortServer) GetUser(ctx context.Context, req *pb_gen.ReqGetUser) (*pb_gen.RespGetUser, error) {
	resp := &pb_gen.RespGetUser{}
	userLogic, logicErr := logic.NewUserLogic(ctx)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	user, logicErr := userLogic.GetUser(req.Id)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.User = user
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	return resp, nil
}

func (s *carPortServer) MGetUser(ctx context.Context, req *pb_gen.ReqMGetUser) (*pb_gen.RespMGetUser, error) {
	resp := &pb_gen.RespMGetUser{}
	userLogic, logicErr := logic.NewUserLogic(ctx)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	users, hasMore, nextOffset, logicErr := userLogic.MGetUser(req.Count, req.Offset)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	resp.Users = users
	resp.HasMore = hasMore
	resp.NextOffset = nextOffset
	return resp, nil

}

func (s *carPortServer) AddCarPort(ctx context.Context, req *pb_gen.ReqAddCarPort) (*pb_gen.RespAddCarPort, error) {
	resp := &pb_gen.RespAddCarPort{}
	carPortLogic, logicErr := logic.NewCarPortLogic(ctx)
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
	carPortLogic, logicErr := logic.NewCarPortLogic(ctx)
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
	carPortLogic, logicErr := logic.NewCarPortLogic(ctx)
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
	carPortLogic, logicErr := logic.NewCarPortLogic(ctx)
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
