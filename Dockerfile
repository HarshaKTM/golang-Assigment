FROM golang:1.20-alpine AS build

WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /bookapi

# Use a smaller image for the final container
FROM alpine:latest  

WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /bookapi .

# Create a data directory for the JSON file
RUN mkdir -p /app/data

# Expose port 5000
EXPOSE 5000

# Command to run the executable
CMD ["./bookapi"] 