package models

import (
	"gorm.io/gorm"
)

// swagger:model Contact
type Contact struct {
	gorm.Model
	// The ContactId of a thing
	// example: 11225727
	ContactId uint64 `json:"contact_id"`

	// The Name of a thing
	// example: Artur Dolgih
	Name string `json:"name"`

	// The Email of a thing
	// example: artur201@mail.ru
	Email string `json:"email"`

	// The AccountId of a number person who create this contact
	AccountId uint64 `json:"account_id"`
}
