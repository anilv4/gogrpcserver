Go gRPC Server Example
======================

This repository contains a simple gRPC server written in Go. 

The server implements a basic service that returns the hostname of the machine it's running on.


Overview
--------

The gRPC service is defined in the 'hostname.proto' file, which is compiled to Go code using the Protocol Buffer Compiler (protoc). 

The server listens on port 50051 and responds to 'GetHostname' requests with the hostname.


Getting Started
---------------

Prerequisites:

- Go (1.18 or later)

- Docker (for containerized deployment)

- Protocol Buffer Compiler (protoc)


Installation:

1. Clone the Repository:

   git clone https://github.com/anilv4/gogrpcserver.git

   cd gogrpcserver


2. Build the Server (Optional):

   If you have Go installed and want to build the server locally:

   cd source

   go build server.go


Running the Server
------------------

There are two ways to run the server: natively or using Docker.


Running Natively:

If you have built the server natively, you can run it directly:

./server


Running with Docker:

1. Build the Docker Image:

   docker build -t gogrpcserver .


2. Run the Docker Container:

   docker run -p 50051:50051 gogrpcserver


Usage
-----

Once the server is running, it will listen for 'GetHostname' gRPC calls on port 50051. 

You can interact with the server using a gRPC client or tools like grpcurl.


Example using grpcurl:

grpcurl -plaintext -d '{}' localhost:50051 hostname.HostnameGetter/GetHostname


Contributing
------------

Contributions to this project are welcome. 

Please feel free to open issues or submit pull requests.


License
-------

This project is open-source and available under the MIT License.
