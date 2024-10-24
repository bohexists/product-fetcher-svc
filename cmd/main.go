package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bohexists/product-fetcher-svc/api/proto"
	"github.com/bohexists/product-fetcher-svc/internal/adapters/mongo"
	"github.com/bohexists/product-fetcher-svc/internal/config"
	"github.com/bohexists/product-fetcher-svc/internal/services"
	"google.golang.org/grpc"
	"net"
)

func main() {
	cfg := config.LoadConfig()

	mongoClient, err := mongo.NewMongoClient(cfg.MongoURL, cfg.MongoDB)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	grpcServer := grpc.NewServer()
	productService := services.NewProductService(mongoClient)
	proto.RegisterProductServiceServer(grpcServer, productService)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	log.Println("gRPC server is running on port 50051")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	log.Println("Shutting down...")
	grpcServer.GracefulStop()
}
