protoc --go_out=../src/ simple_test.proto
protoc --go_out=plugins=grpc:../src/ simple_test.proto
