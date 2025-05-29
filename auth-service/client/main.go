package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/xdars/budget-tracker/auth-service/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	registerResp, err := client.Register(context.Background(), &pb.RegisterRequest{
		Username: "testuser",
		Password: "secret123",
	})
	if err != nil {
		log.Fatalf("Register error: %v", err)
	}
	log.Printf("Registered user: ID=%s Username=%s", registerResp.Id, registerResp.Username)

	loginResp, err := client.Login(context.Background(), &pb.LoginRequest{
		Username: "testuser",
		Password: "secret123",
	})
	if err != nil {
		log.Fatalf("Login error: %v", err)
	}
	log.Printf("Login success: Token=%s", loginResp.Token)
}