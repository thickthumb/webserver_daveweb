# Use Golang official image
FROM golang:1.23.6 

WORKDIR /app
COPY . .

# Build the Go application
RUN  go build -o main

# Create a minimal final image
# FROM alpine:latest
EXPOSE 8080

CMD ["./main"]
