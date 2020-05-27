package main

import (
	"context"
	"software/account/logic"
	"software/account/pb_gen"
	"software/common"
)

type accountServer struct {
}

func (a accountServer) BackendAuth(ctx context.Context, req *pb_gen.ReqBackendAuth) (*pb_gen.RespBackendAuth, error) {
	accountLogic := logic.NewAccountLogic(ctx)
	resp := &pb_gen.RespBackendAuth{}
	logicErr := accountLogic.BackendAuth(req.KeyId, req.KeySecret)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	return resp, nil
}

func (a accountServer) Register(ctx context.Context, req *pb_gen.ReqRegister) (*pb_gen.RespRegister, error) {
	accountLogic := logic.NewAccountLogic(ctx)
	resp := &pb_gen.RespRegister{}
	logicErr := accountLogic.Register(req.Cellphone, req.Password, req.Nickname, req.BindCode)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	return resp, nil
}

func (a accountServer) Verify(ctx context.Context, req *pb_gen.ReqVerify) (*pb_gen.RespVerify, error) {
	accountLogic := logic.NewAccountLogic(ctx)
	resp := &pb_gen.RespVerify{}
	logicErr := accountLogic.Verify(req.Cellphone)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	return resp, nil
}

func (a accountServer) Login(ctx context.Context, req *pb_gen.ReqLogin) (*pb_gen.RespLogin, error) {
	accountLogic := logic.NewAccountLogic(ctx)
	resp := &pb_gen.RespLogin{}
	token, logicErr := accountLogic.Login(req.Cellphone, req.Password)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.Token = token
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	return resp, nil
}

func newAccountServer() pb_gen.AccountServiceServer {
	return new(accountServer)
}
