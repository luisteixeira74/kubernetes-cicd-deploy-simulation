# Etapa 1: build com Go
FROM golang:1.24.2 AS builder

WORKDIR /app
COPY . .

RUN go mod download

# ⚠️ Compila um binário estático compatível com Alpine (musl)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

# Etapa 2: imagem final mínima
FROM alpine:latest
WORKDIR /root/

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
