FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go mod init microservice && CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server .

FROM scratch
WORKDIR /
COPY --from=builder /app/server .
EXPOSE 8080
ENTRYPOINT ["/server"]
