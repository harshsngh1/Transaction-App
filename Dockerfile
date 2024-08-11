FROM golang:1.22.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o transaction-app .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/transaction-app .

EXPOSE 8080

CMD ["./transaction-app"]