package services

import (
	"context"
	"errors"

	"github.com/Ngab-Rio/NOCs-API/internal/dto"
	"github.com/Ngab-Rio/NOCs-API/internal/repository"
	"github.com/Ngab-Rio/NOCs-API/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	authRepo   repository.AuthRepository
	jwtManager utils.JWTManager
}

func NewAuthService(authRepo repository.AuthRepository, jwtManager utils.JWTManager) AuthService {
	return &authService{authRepo: authRepo, jwtManager: jwtManager}
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	if !utils.IsRequired(req.Email) {
		return nil, errors.New("email is required")
	}

	if !utils.IsRequired(req.Password) {
		return nil, errors.New("password is required")
	}

	if !utils.IsValidEmail(req.Email) {
		return nil, errors.New("invalid email")
	}

	if !utils.IsValidPassword(req.Password) {
		return nil, errors.New("invalid password")
	}

	user, err := s.authRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := s.jwtManager.Generate(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{Token: token}, nil
}
