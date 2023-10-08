FROM golang:1.21

RUN mkdir -p $GOPATH/app/
WORKDIR $GOPATH/app

COPY . $GOPATH/app
RUN go mod tidy
