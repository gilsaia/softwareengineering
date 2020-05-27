FROM golang:1.14

WORKDIR /go/src/account
ADD /account /go/src/account
ADD /common /go/src/common

RUN GOPROXY="https://goproxy.io" go get -d -v ./...
RUN GOPROXY="https://goproxy.io" go install -v ./...
RUN GOPROXY="https://goproxy.io" go build .

ENTRYPOINT ["./account"]