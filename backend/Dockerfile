FROM golang:1.26.1-alpine AS builder
COPY go.mod go.sum /app/
WORKDIR /app/
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:3.23
WORKDIR /app/
COPY --from=builder /app/main .
CMD ["./main"]