# Use a smaller Alpine image as the base image
FROM golang:1.21.0-alpine AS build

# Install required dependencies
RUN apk update && apk add --no-cache git

# Set the working directory
WORKDIR /go/src/go-clean-architecture

# Copy only the Go module files and download dependencies separately
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app ./cmd

# Create a minimal image
FROM alpine:3.14
LABEL maintainer="ptan <thanhanphan17@gmail.com>"

# Copy the built binary from the build image
COPY --from=build /app /app
COPY config/env /config/env

# Expose the port the application will run on
EXPOSE 8888

# Run the application
CMD ["/app", "-config", "./config/env", "-env=local", "-upgrade=false"]
