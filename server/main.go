package main

import (
	"context"
	"log"
	"net"

	pb "github.com/fe1t/calculator-gRPC-golang/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	protocol = "tcp"
	port     = ":50051"
)

type server struct{}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddReply, error) {
	return &pb.AddReply{N1: in.N1 + in.N2}, nil
}

func (s *server) Subtract(ctx context.Context, in *pb.SubtractRequest) (*pb.SubtractReply, error) {
	return &pb.SubtractReply{N1: in.N1 - in.N2}, nil
}

func (s *server) Multiply(ctx context.Context, in *pb.MultiplyRequest) (*pb.MultiplyReply, error) {
	return &pb.MultiplyReply{N1: in.N1 * in.N2}, nil
}

func (s *server) Divide(ctx context.Context, in *pb.DivideRequest) (*pb.DivideReply, error) {
	return &pb.DivideReply{N1: in.N1 / in.N2}, nil
}

func main() {
	lis, err := net.Listen(protocol, port)
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
