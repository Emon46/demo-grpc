package main
import (
	"context"
	"google.golang.org/grpc"
	pb "github.com/Emon331046/grpcDemo/helloToDemo"
	"log"
	"net"
)
const (
	port = ":8000"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHelloToDemo(ctx context.Context,in *pb.HelloRequest) ( *pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
func main() {
	lis , err := net.Listen("tcp",port)
	if err != nil {
		log.Fatalf("error listening to tcp port : %v",err)
	}
	serv := grpc.NewServer()
	pb.RegisterGreeterServer(serv,&server{})

	err= serv.Serve(lis)
	if err != nil {
		log.Fatalf("error serving the server : %v",err)
	}
	log.Fatalf("server is running ......")

}