# Step 1: Build the Go app
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files first for dependency management
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o myapp .

# Step 2: Create a smaller image for running the app
FROM alpine:latest

# Set working directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Expose port (optional, if your app uses a specific port)
EXPOSE 8080

# Command to run the application
CMD ["./myapp"]