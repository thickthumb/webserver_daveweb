# Use Golang official image
FROM golang:1.23.6 

WORKDIR /app
COPY . .

# Build the Go application
RUN go mod tidy && go build -o main

# Create a minimal final image
# FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8081

CMD ["./main"]
