package main

import (
	"context"
	"flag"
	"github.com/golang/glog"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"software/car_port/pb_gen"
	"software/common"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, err := net.Listen("tcp", ":30010")
	if err != nil {
		return err
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(common.AuthToken)))
	pb_gen.RegisterCarPortServiceServer(s, newCarPortServer())

	go func() {
		if sErr := s.Serve(lis); sErr != nil {
			return
		}
	}()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))

	carportEndpoint := flag.String("car_port_endpoint", "localhost:30010", "endpoint of car_port")
	opt := []grpc.DialOption{grpc.WithInsecure()}
	err = pb_gen.RegisterCarPortServiceHandlerFromEndpoint(ctx, mux, *carportEndpoint, opt)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8082", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
