# Set the base image to official Golang image
FROM golang:1.16

# Define working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. They will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Expose port 8080 to the outside
EXPOSE 8080

# Command to run the executable
ENTRYPOINT [ "./main" ]