FROM golang:1.24

# Create app directory
WORKDIR /app

# Copy app
COPY . .

# Build the application
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 8080

# Start the application
CMD ["./main"]