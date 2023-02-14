package service

import (
	"github.com/samarec1812/sync-unisender/unisender/models"
	"github.com/samarec1812/sync-unisender/unisender/pkg/repository"
)

type Contacts interface {
	Create(models.Contact) (uint64, error)
	Update(models.Contact) (uint64, error)
	Delete(contactId uint64) (uint64, error)
	Export(uint) (models.ImportContact, error)
}

type Service struct {
	Contacts
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Contacts: NewContactsService(repos.Contacts),
	}
}
