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

func (s *carPortServer) Park(ctx context.Context, req *pb_gen.ReqPark) (*pb_gen.RespPark, error) {
	resp := &pb_gen.RespPark{}
	clientLogic, logicErr := logic.NewClientLogic(ctx)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	logicErr = clientLogic.Park(req.CarPortId)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	return resp, nil
}

func (s *carPortServer) PickUp(ctx context.Context, req *pb_gen.ReqPickUp) (*pb_gen.RespPickUp, error) {
	resp := &pb_gen.RespPickUp{}
	clientLogic, logicErr := logic.NewClientLogic(ctx)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	billId, logicErr := clientLogic.PickUp(req.CarPortId)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.BillId = billId
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	return resp, nil
}

func (s *carPortServer) Pay(ctx context.Context, req *pb_gen.ReqPay) (*pb_gen.RespPay, error) {
	resp := &pb_gen.RespPay{}
	clientLogic, logicErr := logic.NewClientLogic(ctx)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	logicErr = clientLogic.Pay(req.BillId)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	return resp, nil
}

func (s *carPortServer) BillInfo(ctx context.Context, req *pb_gen.ReqBillInfo) (*pb_gen.RespBillInfo, error) {
	resp := &pb_gen.RespBillInfo{}
	clientLogic, logicErr := logic.NewClientLogic(ctx)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	bill, logicErr := clientLogic.BillInfo(req.BillId)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.Bill = bill
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	return resp, nil
}

func (s *carPortServer) UserInfo(ctx context.Context, req *pb_gen.ReqUserInfo) (*pb_gen.RespUserInfo, error) {
	resp := &pb_gen.RespUserInfo{}
	clientLogic, logicErr := logic.NewClientLogic(ctx)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	nickname, carPortId, logicErr := clientLogic.UserInfo(req.Cellphone)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.Nickname = nickname
	resp.CarPortId = carPortId
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	return resp, nil
}

func (s *carPortServer) MetaInfo(ctx context.Context, _ *pb_gen.ReqMetaInfo) (*pb_gen.RespMetaInfo, error) {
	resp := &pb_gen.RespMetaInfo{}
	clientLogic, logicErr := logic.NewClientLogic(ctx)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	count, parks, logicErr := clientLogic.MetaInfo()
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.Count = count
	resp.Parks = parks
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	return resp, nil
}

func (s *carPortServer) ParkInfo(ctx context.Context, req *pb_gen.ReqParkInfo) (*pb_gen.RespParkInfo, error) {
	resp := &pb_gen.RespParkInfo{}
	clientLogic, logicErr := logic.NewClientLogic(ctx)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	parkName, carPorts, logicErr := clientLogic.ParkInfo(req.ParkId)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ParkName = parkName
	resp.Ports = carPorts
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	return resp, nil
}

func (s *carPortServer) GetBill(ctx context.Context, req *pb_gen.ReqGetBill) (*pb_gen.RespGetBill, error) {
	resp := &pb_gen.RespGetBill{}
	billLogic, logicErr := logic.NewBillLogic(ctx)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	bill, logicErr := billLogic.GetBill(req.BillId)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	resp.Bill = bill
	return resp, nil

}

func (s *carPortServer) MGetBill(ctx context.Context, req *pb_gen.ReqMGetBill) (*pb_gen.RespMGetBill, error) {
	resp := &pb_gen.RespMGetBill{}
	billLogic, logicErr := logic.NewBillLogic(ctx)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	bills, tableCount, logicErr := billLogic.MGetBill(req.Count, req.Num)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	resp.Bills = bills
	resp.Count = tableCount
	return resp, nil
}

func (s *carPortServer) AddPark(ctx context.Context, req *pb_gen.ReqAddPark) (*pb_gen.RespAddPark, error) {
	resp := &pb_gen.RespAddPark{}
	parkLogic, logicErr := logic.NewParkLogic(ctx)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	logicErr = parkLogic.CreatePark(req.Park)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	return resp, nil
}

func (s *carPortServer) UpdatePark(ctx context.Context, req *pb_gen.ReqUpdatePark) (*pb_gen.RespUpdatePark, error) {
	resp := &pb_gen.RespUpdatePark{}
	parkLogic, logicErr := logic.NewParkLogic(ctx)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	logicErr = parkLogic.UpdatePark(req.Park)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	return resp, nil
}

func (s *carPortServer) GetPark(ctx context.Context, req *pb_gen.ReqGetPark) (*pb_gen.RespGetPark, error) {
	resp := &pb_gen.RespGetPark{}
	parkLogic, logicErr := logic.NewParkLogic(ctx)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	park, logicErr := parkLogic.GetPark(req.ParkId)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	resp.Park = park
	return resp, nil
}

func (s *carPortServer) MGetPark(ctx context.Context, req *pb_gen.ReqMGetPark) (*pb_gen.RespMGetPark, error) {
	resp := &pb_gen.RespMGetPark{}
	parkLogic, logicErr := logic.NewParkLogic(ctx)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	parks, tableCount, logicErr := parkLogic.MGetPark(req.Count, req.Num)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	resp.Parks = parks
	resp.Count = tableCount
	//resp.HasMore = hasMore
	//resp.NextOffset = nextOffset
	return resp, nil
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
	user, logicErr := userLogic.GetUser(req.Id, req.Cellphone)
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
	users, tableCount, logicErr := userLogic.MGetUser(req.Count, req.Num)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	resp.Users = users
	resp.Count = tableCount
	//resp.HasMore = hasMore
	//resp.NextOffset = nextOffset
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
	carPorts, tableCount, logicErr := carPortLogic.MGetCarPort(req.Count, req.Num)
	if !logicErr.Is(common.Success) {
		resp.ErrNo, resp.ErrTips = logicErr.ErrNoMsg()
		return resp, nil
	}
	resp.ErrNo, resp.ErrTips = common.Success.ErrNoMsg()
	resp.Ports = carPorts
	resp.Count = tableCount
	//resp.HasMore = hasMore
	//resp.NextOffset = nextOffset
	return resp, nil
}
