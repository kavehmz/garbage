package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/kavehmz/garbage/050_grpc/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayBye(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{Message: "Bye " + in.Name}, nil
}

func startGRPC() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi %s!", r.URL.Path[1:])
}

func startHTTP() {
	http.HandleFunc("/test/", handler)
	http.ListenAndServe(":8080", nil)
}

func main() {
	go startGRPC()
	startHTTP()
}
