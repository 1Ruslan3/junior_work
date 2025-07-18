FROM golang:1.24.2 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o junior ./main.go

FROM gcr.io/distroless/base-debian11
WORKDIR /app
COPY --from=builder /app/junior .
CMD ["./junior"]
