package service

import (
	"github.com/samarec1812/sync-unisender/auth/models"
	"github.com/samarec1812/sync-unisender/auth/pkg/repository"
)

type AuthAmoService struct {
	repo repository.AuthorizationAmo
}

func NewAuthAmoService(repo repository.AuthorizationAmo) *AuthAmoService {
	return &AuthAmoService{repo: repo}
}

func (s *AuthAmoService) GetAuthCode(user models.User) (string, error) {
	return s.repo.GetAuthCode(user)
}
