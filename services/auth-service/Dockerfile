FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/api/main.go

FROM alpine:3.19
WORKDIR /app
RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta
COPY --from=builder /app/app .
EXPOSE 8081
CMD ["./app"]
