FROM golang:1.14

WORKDIR /go/src/car_port
ADD /car_port /go/src/car_port
ADD /common /go/src/common

RUN GOPROXY="https://goproxy.io" go get -d -v ./...
RUN GOPROXY="https://goproxy.io" go install -v ./...
RUN GOPROXY="https://goproxy.io" go build .

ENTRYPOINT ["./car_port"]