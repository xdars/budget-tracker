package service

import (
	"context"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/xdars/budget-tracker/auth-service/internal/db"
	pb "github.com/xdars/budget-tracker/auth-service/proto"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	DB *db.InMemoryDB
}

func NewAuthServer(database *db.InMemoryDB) *AuthServer {
	return &AuthServer{DB: database}
}

func (s *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := db.User{
		ID:       uuid.New().String(),
		Username: req.Username,
		Password: string(hashed),
	}

	if err := s.DB.CreateUser(user); err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Id:       user.ID,
		Username: user.Username,
	}, nil
}

func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, exists := s.DB.GetUser(req.Username)
	if !exists {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return &pb.LoginResponse{Token: "token-for-" + user.Username}, nil
}

var ErrInvalidCredentials = &LoginError{"invalid credentials"}

type LoginError struct {
	msg string
}

func (e *LoginError) Error() string {
	return e.msg
}