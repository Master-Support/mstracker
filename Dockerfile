FROM golang:alpine as builder
WORKDIR /go/src/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/mstracker_api main.go
CMD ["/go/bin/mstracker_api"]
# FROM scratch
# COPY --from=builder /go/bin/server /mstracker_api
# EXPOSE 8080
# CMD ["/mstracker_api"]

