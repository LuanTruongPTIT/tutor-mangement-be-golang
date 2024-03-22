package services

import (
	"context"
	"fmt"

	"github.com/LuanTruongPTIT/tutor-be/modules/auth/dto"
	"github.com/LuanTruongPTIT/tutor-be/modules/auth/model"
	"github.com/LuanTruongPTIT/tutor-be/modules/auth/repository"
	"github.com/LuanTruongPTIT/tutor-be/pkg/utils"
	"github.com/go-playground/validator/v10"
)

type IAuthService interface {
	Register(ctx context.Context, req *dto.RegisterReq) (*model.Auth, error)
}
type AuthService struct {
	validator validator.Validate
	repo      repository.IAuthRepository
}

func NewAuthService(repo repository.IAuthRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Register(ctx context.Context, req *dto.RegisterReq) (*model.Auth, error) {
	//  err := s.validator.Struct(req)
	if err := s.validator.Struct(req); err != nil {
		return nil, err
	}
	var auth model.Auth
	utils.Copy(&auth, &req)
	err := s.repo.Create(ctx, &auth)
	if err != nil {
		fmt.Printf("Register.Create fail, email: %s, error: %s", req.Email, err)
		return nil, err
	}
	return &auth, err
}
