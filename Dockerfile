FROM golang:1.21.5

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/main main.go

RUN chmod +x ./run.sh

ENTRYPOINT [ "sh", "./run.sh" ]