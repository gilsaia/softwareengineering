FROM golang:1.14

WORKDIR /go/src/car_port
ADD . /go/src/car_port

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build .

ENTRYPOINT ["./car_port"]