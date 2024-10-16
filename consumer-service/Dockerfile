FROM golang:1.21.5-alpine

WORKDIR /consumer-service

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/main ./cmd/main.go

EXPOSE 8000

CMD ["./bin/main"]