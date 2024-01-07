# Use a multi-stage build for a smaller final image
FROM golang:1.21.3

WORKDIR /usr/src/app

# Install air for live reloading (if needed)
RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy

# Build the Go application
# RUN go build -o /go/bin/mstracker_api ./cmd/main.go

# # Use a smaller image for the final stage
# FROM alpine:latest

# WORKDIR /app

# COPY --from=builder /go/bin/mstracker_api .

# EXPOSE 8080

# CMD ["./mstracker_api"]
