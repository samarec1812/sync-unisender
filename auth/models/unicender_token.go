package models

import (
	"errors"

	"gorm.io/gorm"
)

type UnisenderUsers struct {
	gorm.Model
	Token     string `json:"token"`
	Uname     string `json:"uname"`
	AccountId uint64 `json:"account_id"`
}

func (u UnisenderUsers) Validate() error {
	if u.Uname == "" {
		return errors.New("have not unisender uname info")
	}
	if u.Token == "" {
		return errors.New("have not unisender token info")
	}
	if u.AccountId <= 0 {
		return errors.New("have not unisender account_id  info")
	}

	return nil
}
