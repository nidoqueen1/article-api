FROM golang:1.20 AS builder
COPY . /app
WORKDIR /app
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/config.yml .

RUN chmod +x main
CMD ["./main"]
