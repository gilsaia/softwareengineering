protoc --go_out=../src/ simple_test.proto
protoc --go_out=plugins=grpc:../src/ simple_test.proto
protoc -I. -I=$GOPATH/src -I=../pkg/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:../src/ simple_test.proto
protoc -I. --grpc-gateway_out=logtostderr=true,paths=source_relative:../src/software/simple_test/pb_gen simple_test.proto

protoc --go_out=plugins=grpc:../src/ car_port.proto
protoc -I. --grpc-gateway_out=logtostderr=true,paths=source_relative:../src/software/car_port/pb_gen car_port.proto

protoc --go_out=plugins=grpc:../src/ account.proto
protoc -I. --grpc-gateway_out=logtostderr=true,paths=source_relative:../src/software/account/pb_gen account.proto
