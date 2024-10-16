FROM golang:1.21.5-alpine

WORKDIR /server

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/main ./cmd/main.go

EXPOSE 8080

CMD ["./bin/main"]