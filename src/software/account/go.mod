module software/account

go 1.14

require (
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.224
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.14.6
	github.com/jinzhu/gorm v1.9.12
	github.com/jmespath/go-jmespath v0.3.0 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	golang.org/x/crypto v0.1.0
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.24.0
	gopkg.in/ini.v1 v1.56.0 // indirect
	software/common v0.0.0-00010101000000-000000000000
)

replace software/common => ../common
