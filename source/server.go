package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "gogrpcserver/pb" // Make sure this import path matches your setup
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedHostnameGetterServer
}

func (s *server) GetHostname(ctx context.Context, in *pb.HostnameRequest) (*pb.HostnameReply, error) {
	// Extract client IP from context
	p, ok := peer.FromContext(ctx)
	if ok {
		clientIP := p.Addr.String()
		log.Printf("Received GetHostname request from client IP: %s", clientIP)
	} else {
		log.Printf("Received GetHostname request but could not determine client IP")
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Error getting hostname: %v", err)
		return nil, err
	}

	log.Printf("Sending hostname: %s", hostname)
	return &pb.HostnameReply{Message: "Hostname: " + hostname}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHostnameGetterServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	log.Printf("Server is listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
