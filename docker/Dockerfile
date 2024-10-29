FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY ../go.mod ../go.sum ./

RUN go mod download

COPY . .

RUN go build -o product-fetcher-svc ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/product-fetcher-svc .

CMD ["./product-fetcher-svc"]