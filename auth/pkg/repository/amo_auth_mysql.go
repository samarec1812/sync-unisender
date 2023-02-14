package repository

import (
	"github.com/samarec1812/sync-unisender/auth/models"
	"gorm.io/gorm"
)

type AuthAmoMysql struct {
	db *gorm.DB
}

func NewAuthAmoMysql(db *gorm.DB) *AuthAmoMysql {
	return &AuthAmoMysql{db: db}
}

func (r *AuthAmoMysql) GetAuthCode(user models.User) (string, error) {
	res := r.db.Create(&user)
	if res.Error != nil {
		return "", res.Error
	}
	return user.AccessToken, nil
}
