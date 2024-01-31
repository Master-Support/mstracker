FROM golang:1.21.0

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/cosmtrek/air@latest

RUN go build -o /mstracker ./cmd

EXPOSE 8080

CMD ["air", "cmd/main.go", "-b", "0.0.0.0"]
