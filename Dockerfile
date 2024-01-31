FROM golang:1.21.0

# Create the working directory and set it as the working directory
WORKDIR /app

# Copy the Go dependency files go.mod and go.sum
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy all Go files recursively
COPY . .

# Install 'air' binary
RUN go install github.com/cosmtrek/air@latest

# Build the application
RUN go build -o /mstracker ./cmd

# Optional: Expose the port that the application will listen on
EXPOSE 8080

# Command to run the application using 'air'
CMD ["air", "cmd/main.go", "-b", "0.0.0.0"]
