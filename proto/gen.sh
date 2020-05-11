protoc --go_out=../src/ simple_test.proto
protoc --go_out=plugins=grpc:../src/ simple_test.proto
protoc -I. -I=$GOPATH/src -I=../pkg/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:../src/ simple_test.proto
