package service

import (
	"github.com/samarec1812/sync-unisender/auth/models"
	"github.com/samarec1812/sync-unisender/auth/pkg/repository"
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
