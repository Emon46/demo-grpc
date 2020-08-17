package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"

	pb "github.com/Emon331046/grpcDemo/helloToDemo"
	"os"
	"time"
)

const add = "localhost:8000"

func main() {
	flag.Parse()
	connection , err := grpc.Dial(add, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error :%v", err)
		return
	}
	defer connection.Close()

	gc := pb.NewGreeterClient(connection)

	name := "world of demo"
	if len(os.Args) >1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	r ,err := gc.SayHelloToDemo(ctx, &pb.HelloRequest{Name: name})

	if err != nil {
		log.Fatalf("err %v",err)
	}
	log.Fatalf("greetings : %v",r.GetMessage() )
	return

}

