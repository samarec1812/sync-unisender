package repository

import (
	"github.com/samarec1812/sync-unisender/unisender/models"
	"gorm.io/gorm"
)

type Contacts interface {
	Create(models.Contact) (uint64, error)
	Update(models.Contact) (uint64, error)
	Delete(contactId uint64) (uint64, error)
	Export(uint) (models.ImportContact, error)
}

type Repository struct {
	Contacts
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Contacts: NewContactsMysql(db),
	}
}
