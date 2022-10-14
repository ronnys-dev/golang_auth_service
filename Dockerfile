FROM golang:1.19-alpine as builder

RUN apk update && apk add --no-cache  \
        gcc               \
        musl-dev

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN go build main.go

FROM golang:1.19-alpine
ENV APP_PORT=8080
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app ./
EXPOSE ${APP_PORT}
ENTRYPOINT ["./main"]