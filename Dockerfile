# Use the official Golang image as a base image
FROM golang:1.24

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port your application listens on
EXPOSE 8080

# Command to run the application
CMD ["./main"]