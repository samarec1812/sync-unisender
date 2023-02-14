package models

type Token struct {
	Type         string `json:"token_type"`
	Expires      int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
