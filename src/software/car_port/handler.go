package car_port

import (
	"context"
	"software/car_port/pb_gen"
)

type carPortServer struct {
}

func newCarPortServer() pb_gen.CarPortServiceServer {
	return new(carPortServer)
}

func (s *carPortServer) AddCarPort(ctx context.Context, req *pb_gen.ReqAddCarPort) (*pb_gen.RespAddCarPort, error) {

}

func (s *carPortServer) UpdateCarPort(ctx context.Context, req *pb_gen.ReqUpdateCarPort) (*pb_gen.RespUpdateCarPort, error) {
	panic("implement me")
}

func (s *carPortServer) GetCarPort(ctx context.Context, req *pb_gen.ReqGetCarPort) (*pb_gen.RespGetCarPort, error) {
	panic("implement me")
}

func (s *carPortServer) MGetCarPort(ctx context.Context, req *pb_gen.ReqMGetCarPort) (*pb_gen.RespMGetCarPort, error) {
	panic("implement me")
}
