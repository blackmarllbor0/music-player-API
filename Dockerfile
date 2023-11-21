# Use an official Go runtime as a base image
FROM golang:1.17-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Go application with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o myapp .

# Create a minimal runtime image
FROM scratch
COPY --from=builder /app/myapp /app/myapp

# Expose the port on which the application will run
EXPOSE 8080

# Define the command to run the executable
CMD ["/app/myapp"]
