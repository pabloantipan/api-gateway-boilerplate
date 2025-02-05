package service

import (
	"context"
	"errors"

	"github.com/pabloantipan/go-api-gateway-poc/internal/infrastructure/cloud"
	"github.com/pabloantipan/go-api-gateway-poc/internal/presentation/dto"
)

type AuthService interface {
	ValidateToken(ctx context.Context, token string) (string, error)
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	firebaseClient *cloud.FirebaseClient
}

func NewAuthService(fc *cloud.FirebaseClient) AuthService {
	return &authService{
		firebaseClient: fc,
	}
}

func (s *authService) ValidateToken(ctx context.Context, token string) (string, error) {
	decodedToken, err := s.firebaseClient.VerifyToken(ctx, token)
	if err != nil {
		return "", err
	}
	return decodedToken.UID, nil
}

func (s *authService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email and password are required")
	}

	token, err := s.firebaseClient.SignInWithPassword(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	if token == "" {
		return nil, errors.New("invalid credentials")
	}

	return &dto.LoginResponse{
		Token:   token,
		Message: "Login successful",
	}, nil

}
