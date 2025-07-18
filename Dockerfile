FROM golang:1.24.2 

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init
RUN go build -o main .

CMD ["./main"]