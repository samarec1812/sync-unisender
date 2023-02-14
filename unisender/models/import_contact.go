package models

type ImportContact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
