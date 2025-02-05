package service

import (
	"context"

	"github.com/pabloantipan/go-api-gateway-poc/internal/infrastructure/cloud"
)

type AuthService interface {
	ValidateToken(ctx context.Context, token string) (string, error)
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
