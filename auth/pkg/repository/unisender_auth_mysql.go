package repository

import (
	"fmt"

	"github.com/samarec1812/sync-unisender/auth/models"
	"gorm.io/gorm"
)

type AuthUnisenderMysql struct {
	db *gorm.DB
}

func NewAuthUnisenderMysql(db *gorm.DB) *AuthUnisenderMysql {
	return &AuthUnisenderMysql{db: db}
}

func (r *AuthUnisenderMysql) SaveToken(token models.UnisenderUsers) error {
	res := r.db.Create(&token)
	fmt.Println("DB: ERROR: ", res.Error)
	return res.Error
}
