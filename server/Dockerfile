FROM golang:1.22-alpine

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

CMD ["./main"]

EXPOSE 8080
