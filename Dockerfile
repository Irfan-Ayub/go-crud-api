# Start from a base image containing Go runtime
FROM golang:1.17 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Verify the contents of /app
RUN ls -la /app

# Start from a base image containing minimal runtime
FROM alpine:latest

# Install necessary packages
RUN apk add --no-cache libc6-compat

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Ensure the binary has executable permissions
RUN chmod +x main

# Verify the contents of /root/
RUN ls -la /root/

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./main"]
