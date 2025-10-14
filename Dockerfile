FROM golang:1.25-alpine

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o ./main main.go

EXPOSE 3000

CMD ["./main"]
