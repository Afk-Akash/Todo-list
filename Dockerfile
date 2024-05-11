# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory in the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Install dependencies
RUN go mod download

# Build the Go application inside the container
RUN go build -o containerizedapp .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./containerizedapp"]
