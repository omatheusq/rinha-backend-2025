FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app-binary ./cmd/server

FROM alpine:latest
COPY --from=builder /app-binary /app-binary

EXPOSE 8080
CMD ["/app-binary"]
