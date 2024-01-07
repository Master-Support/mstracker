FROM golang:alpine as builder

WORKDIR /go/src/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/mstracker_api main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /go/bin/mstracker_api .

EXPOSE 8080

ENV DB_USER=postgres
ENV DB_PASSWORD=postgres
ENV DB_HOST=172.17.0.2
ENV DB_PORT=5432
ENV DB_NAME=mstracker

CMD ["./mstracker_api"]
