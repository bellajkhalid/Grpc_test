package main

import (
	"context"
	"log"
	"net"

	pb "github.com/bellajkhalid/Grpc_test/pb_go_grpc"
	"google.golang.org/grpc"
)

type Server_1 struct {
	pb.UnimplementedService_1Server
}

func (S *Server_1) SResponse(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.Response{
		Pdf:  []byte("test1"),
		word: []byte("test1"),
	}, nil
}
func main() {

	// create socket to listen to requests
	Lis, err := net.Listen("tcp", "localhost:8765")
	if err != nil {
		log.Fatal("failed to listen : %v", err)

	}
	// create new gRPC server
	servre := grpc.NewServer()

	pb.RegisterService_1Server(servre, &Server_1{})
	log.Printf("server listening at %v", Lis.Addr())
	if erro := servre.serve(Lis); erro != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
