# Build stage
FROM golang:1.24.5 AS builder

# Goes to the app directory
WORKDIR /app

# Copy go mod files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy rest of the app into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o calculator .

# Final stage
FROM alpine:latest

# Goes to the app directory (like cd)
WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/calculator .

# Copy static files
COPY --from=builder /app/index.html .
COPY --from=builder /app/dev.index.html .
COPY --from=builder /app/static/style.css ./static/

# Expose port
EXPOSE 8080

# Set enviroment to development
ENV ENVIROMENT=dev

# Run the application
CMD ["./calculator"]