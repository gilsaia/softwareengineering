FROM golang:1.14

WORKDIR /go/src/simple_test
ADD . /go/src/simple_test

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build .

ENTRYPOINT ["./simple_test"]