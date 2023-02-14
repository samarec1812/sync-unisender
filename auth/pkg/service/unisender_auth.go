package service

import (
	"github.com/samarec1812/sync-unisender/auth/models"
	"github.com/samarec1812/sync-unisender/auth/pkg/repository"
)

type AuthUnisenderService struct {
	repo repository.AuthorizationUnisender
}

func NewAuthUnisenderService(repo repository.AuthorizationUnisender) *AuthUnisenderService {
	return &AuthUnisenderService{repo: repo}
}

func (s *AuthUnisenderService) SaveToken(token models.UnisenderUsers) error {
	return s.repo.SaveToken(token)
}
