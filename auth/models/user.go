package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ClientId     string `json:"client_id" gorm:"unique"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	RedirectUri  string `json:"redirect_uri"`
}
