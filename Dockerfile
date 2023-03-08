FROM golang:latest

ENV GO111MODULE="on" GOPROXY=https://goproxy.cn,https://goproxy.io,direct

WORKDIR /app

COPY . /app

RUN go build .



ENTRYPOINT ["./snake"]