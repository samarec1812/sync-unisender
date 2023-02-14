FROM golang:1.18-alpine as builder
WORKDIR /src/app

COPY auth/ ./auth/
COPY proto/ ./proto/
COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o application ./auth/cmd/main.go
# Финальный этап, копируем собранное приложение
FROM alpine:3.15
WORKDIR /
COPY --from=builder /src/app ./app
EXPOSE 8081

CMD ["./app/application"]