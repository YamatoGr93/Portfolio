# Use Go 1.23
FROM golang:1.23

# Set working directory
WORKDIR /app

# Copy module files first
COPY go.mod go.sum ./

# Download dependencies explicitly
RUN go mod download

# Copy all project files
COPY . .

# Build the Go binary
RUN go build -o xss-scanner main.go

# Set the entrypoint
ENTRYPOINT ["/app/xss-scanner"]
