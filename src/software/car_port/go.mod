module software/car_port

go 1.14

require (
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.224
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.14.5
	github.com/jinzhu/gorm v1.9.12
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
	google.golang.org/genproto v0.0.0-20200521103424-e9a78aa275b7
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.23.0
	software/common v0.0.0-00010101000000-000000000000
)

replace software/common => ../common
