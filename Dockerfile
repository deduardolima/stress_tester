FROM golang:1.20-buster as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /stress-tester ./cmd/stress_tester

FROM debian:buster-slim

WORKDIR /root/

RUN apt-get update && apt-get install -y apache2-utils iputils-ping curl && apt-get clean

COPY --from=builder /stress-tester /stress-tester

ENTRYPOINT ["/stress-tester"]

