package repository

import (
	"github.com/samarec1812/sync-unisender/auth/models"
	"gorm.io/gorm"
)

type AuthorizationAmo interface {
	GetAuthCode(user models.User) (string, error)
}

type AuthorizationUnisender interface {
	SaveToken(token models.UnisenderUsers) error
}

type Repository struct {
	AuthorizationAmo
	AuthorizationUnisender
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		AuthorizationAmo:       NewAuthAmoMysql(db),
		AuthorizationUnisender: NewAuthUnisenderMysql(db),
	}
}
