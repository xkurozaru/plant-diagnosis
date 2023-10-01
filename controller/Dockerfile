FROM golang:1.21

WORKDIR /go/src/github.com/xkurozaru/plant-diagnosis/controller

COPY go.mod go.sum ./
RUN go mod download

COPY . .
