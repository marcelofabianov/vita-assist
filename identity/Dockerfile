FROM golang:1.23-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest
RUN go install github.com/jackc/tern/v2@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

RUN go build -o /app/main/api ./cmd/api/main.go

EXPOSE 8081

CMD ["air", "-c", "/app/.air.toml"]
