FROM golang:1.11-alpine3.8 AS build
# Support CGO and SSL
RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR /go/src/app
COPY app/main.go main.go
WORKDIR /go
COPY app/src src
WORKDIR /go/src/app
RUN go get -u github.com/labstack/echo/...
RUN go get -u github.com/lib/pq
RUN go get -u github.com/joho/godotenv/cmd/godotenv
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/test ./main.go

FROM alpine:3.8
# เอาไว้ใช้ตอน Repo หลักพัง
RUN echo http://mirror.yandex.ru/mirrors/alpine/v3.5/main > /etc/apk/repositories; \
    echo http://mirror.yandex.ru/mirrors/alpine/v3.5/community >> /etc/apk/repositories
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
EXPOSE 8081
ENTRYPOINT /go/bin/test --port 8081

