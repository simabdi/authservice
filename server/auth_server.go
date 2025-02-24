package server

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/simabdi/authservice/config"
	"github.com/simabdi/authservice/models"
	pb "github.com/simabdi/authservice/proto"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
}

// Register new user
func (s *AuthServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := models.User{Name: req.Name, Email: req.Email, Password: string(hashedPassword)}
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to register user: %v", err)
	}

	return &pb.RegisterResponse{Message: "User registered successfully"}, nil
}

func (s *AuthServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("account didnt match")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	tokenString, _ := token.SignedString([]byte(secret))

	return &pb.LoginResponse{Token: tokenString}, nil
}
