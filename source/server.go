package main

import (
        "context"
        "log"
        "net"
        "os"

        "gogrpcserver/pb"
        "google.golang.org/grpc"
        "google.golang.org/grpc/reflection"
)

type server struct {
        pb.UnimplementedHostnameGetterServer
}

func (s *server) GetHostname(ctx context.Context, in *pb.HostnameRequest) (*pb.HostnameReply, error) {
        hostname, err := os.Hostname()
        if err != nil {
                return nil, err
        }
        return &pb.HostnameReply{Message: "Hostname: " + hostname}, nil
}

func main() {
        lis, err := net.Listen("tcp", ":50051")
        if err != nil {
                log.Fatalf("Failed to listen: %v", err)
        }

        s := grpc.NewServer()
        pb.RegisterHostnameGetterServer(s, &server{})
        reflection.Register(s)

        log.Printf("Server is listening at %v", lis.Addr())
        if err := s.Serve(lis); err != nil {
                log.Fatalf("Failed to serve: %v", err)
        }
}
