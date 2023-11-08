# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . .


# Build the Go application
RUN go build -o containerizedapp .

# Command to run the application
CMD ["./containerizedapp"]
