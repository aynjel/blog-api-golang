FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . /app
RUN go mod download
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main /app/main
EXPOSE 8080
CMD ["/app/main"]