package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/xdars/budget-tracker/auth-service/internal/db"
	"github.com/xdars/budget-tracker/auth-service/internal/service"
	pb "github.com/xdars/budget-tracker/auth-service/proto"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	database := db.NewInMemoryDB()
	authServer := service.NewAuthServer(database)

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authServer)

	log.Println("AuthService is running on port :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}