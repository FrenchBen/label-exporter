# Start from the latest golang base image
FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

ARG CGO_ENABLED=0
ARG GOOS=linux

# Test and Build the Go app
RUN go test -v ./... && go build -a -installsuffix cgo -o label-exporter cmd/label-exporter/main.go

######## Start a new stage from scratch #######
FROM scratch

# Add Maintainer Info
LABEL maintainer="micnncim, frenchben"

WORKDIR /app/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/label-exporter .

# Entrypoint is the main executable
ENTRYPOINT ["/app/label-exporter"]

# Command to run against the executable
CMD ["--help"]
