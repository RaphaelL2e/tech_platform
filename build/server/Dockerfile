FROM golang:1.16 AS builder
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /tech_platform
COPY go.mod go.sum ./
RUN go mod download

COPY server server

RUN GO111MODULE=on CGO_ENABLED=0 go install ./...

FROM alpine:3.12
WORKDIR /usr/bin
COPY --from=builder /go/bin .

EXPOSE 8080
CMD ["cmd"]