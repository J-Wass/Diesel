# diesel/Dockerfile
FROM golang:1.19.8

WORKDIR /app

# Copy project files and build the app
COPY . .
RUN go mod download
RUN go build -o diesel-server

# Expose the application port
EXPOSE 8080

# Start the application
CMD ["./diesel-server"]
