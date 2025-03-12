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

# Build the application
WORKDIR /app/api/service/geofence
RUN go build -o geofence

# Create a new lightweight image for running the application
FROM alpine:latest

# Copy the pre-built binary file from the builder stage
COPY --from=builder /app/api/service/geofence/geofence /app/geofence

# Expose the port the application listens on
EXPOSE 3000

# # Set environment variables
# ENV GOMAXPROCS=${GOMAXPROCS:-1}
# ENV GOGC=${GOGC:-off}
# ENV GOMEMLIMIT=${GOMEMLIMIT:-8000MiB}

# Command to run the executable
CMD ["/app/geofence"]