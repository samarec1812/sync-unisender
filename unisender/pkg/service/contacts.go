package service

import (
	"github.com/samarec1812/sync-unisender/unisender/models"
	"github.com/samarec1812/sync-unisender/unisender/pkg/repository"
)

type ContactsService struct {
	repo repository.Contacts
}

func NewContactsService(repo repository.Contacts) *ContactsService {
	return &ContactsService{repo: repo}
}

func (s *ContactsService) Create(contact models.Contact) (uint64, error) {
	return s.repo.Create(contact)
}

func (s *ContactsService) Update(contact models.Contact) (uint64, error) {
	return s.repo.Update(contact)
}

func (s *ContactsService) Delete(contactId uint64) (uint64, error) {
	return s.repo.Delete(contactId)
}

func (s *ContactsService) Export(id uint) (models.ImportContact, error) {
	return s.repo.Export(id)
}
