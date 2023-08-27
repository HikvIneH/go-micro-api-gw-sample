# Use a base image with Go installed
FROM golang

# Set the working directory inside the container
WORKDIR /app

# Copy your Go application code to the container
COPY . /app

# Install protoc and grpc-tools
RUN apt-get update && \
    apt-get install -y protobuf-compiler 


RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate Go gRPC code from proto files
RUN make proto

# Compile the Go application
RUN go build -o order-svc-server ./cmd/main.go

# Expose the port the application listens on
EXPOSE 8082

# Command to run the application
CMD ["./order-svc-server"]