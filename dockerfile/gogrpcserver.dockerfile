# Use Fedora as the base image
FROM registry.fedoraproject.org/fedora:latest as builder

# Enable fastest mirror in dnf configuration.
RUN echo 'fastestmirror=1' >> /etc/dnf/dnf.conf

# Install Go, Protobuf Compiler, Git
RUN dnf -y update && \
    dnf -y install golang protobuf-compiler git

# Set the working directory in the container
WORKDIR /app

# Clone the repository
RUN git clone https://github.com/anilv4/gogrpcserver.git .

# Change to the 'source' directory
WORKDIR /app/source

RUN mkdir pb

RUN go mod init gogrpcserver

# Install the protoc-gen-go and protoc-gen-go-grpc plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Add the Go bin directory to PATH
ENV PATH="${PATH}:/root/go/bin"

RUN go mod tidy

# Generate Go code from the .proto file
RUN protoc --go_out=./pb --go_opt=paths=source_relative \
           --go-grpc_out=./pb --go-grpc_opt=paths=source_relative \
           hostname.proto

# Build the application
RUN go build -o server ./server.go

# Use a smaller base image to run the server
FROM registry.fedoraproject.org/fedora:latest

# Copy the server binary from the builder stage
COPY --from=builder /app/source/server /server

# Expose the port the server listens on
EXPOSE 50051

# Command to run the server binary
CMD ["/server"]
