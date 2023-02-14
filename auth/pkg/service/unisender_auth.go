package service

import (
	"git.amocrm.ru/akondratev/amo-project/auth/models"
	"git.amocrm.ru/akondratev/amo-project/auth/pkg/repository"
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
