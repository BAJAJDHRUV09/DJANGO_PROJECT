package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "grpcservice/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()

	// Set up a connection to the server
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewModuleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	log.Println("Testing ListModules...")
	listResp, err := client.ListModules(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not list modules: %v", err)
	}
	log.Printf("Found %d modules:", len(listResp.Modules))
	for _, module := range listResp.Modules {
		log.Printf("- Module %d: %s (Grade %s, Chapter %s)",
			module.Id, module.Subject, module.Grade, module.Chapter)
	}

	if len(listResp.Modules) > 0 {
		firstModuleID := listResp.Modules[0].Id
		log.Printf("\nTesting GetModule for ID %d...", firstModuleID)

		getResp, err := client.GetModule(ctx, &pb.ModuleRequest{Id: firstModuleID})
		if err != nil {
			log.Fatalf("could not get module: %v", err)
		}
		log.Printf("Module details:")
		log.Printf("- ID: %d", getResp.Id)
		log.Printf("- Subject: %s", getResp.Subject)
		log.Printf("- Grade: %s", getResp.Grade)
		log.Printf("- Chapter: %s", getResp.Chapter)
		log.Printf("- Content: %s", getResp.ContentJson)
	}
}
