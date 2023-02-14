package service

import (
	"git.amocrm.ru/akondratev/amo-project/auth/models"
	"git.amocrm.ru/akondratev/amo-project/auth/pkg/repository"
)

type AuthorizationAmo interface {
	GetAuthCode(user models.User) (string, error)
}

type AuthorizationUnisender interface {
	SaveToken(token models.UnisenderUsers) error
}

type Service struct {
	AuthorizationAmo
	AuthorizationUnisender
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		AuthorizationAmo:       NewAuthAmoService(repos.AuthorizationAmo),
		AuthorizationUnisender: NewAuthUnisenderService(repos.AuthorizationUnisender),
	}
}
