# Stage 1: Build the application
FROM golang:1.21-alpine AS builder
WORKDIR /app

# Copy only necessary files first to leverage Docker caching
COPY App/go.mod App/go.sum ./
RUN go mod download

# Copy the application source code
COPY App/ ./

# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o main .

# Stage 2: Create the smallest runtime image
FROM scratch
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./main"]
