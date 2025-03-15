# Dockerfile
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Copy the vendor directory
COPY vendor ./vendor

# Copy the source code into the container
COPY . .     

# Build the admin application.
WORKDIR /app/api/tooling/admin
RUN go build -o admin

# Build the application
WORKDIR /app/api/service/geofence
RUN go build -o geofence

# Create a new lightweight image for running the application
FROM alpine:latest

# Copy the pre-built binary file from the builder stage
COPY --from=builder /app/api/service/geofence/geofence /app/geofence
COPY --from=builder /app/api/tooling/admin/admin /app/admin

# Expose the port the application listens on
EXPOSE 3000

# Command to run the executable
CMD ["/app/geofence"]