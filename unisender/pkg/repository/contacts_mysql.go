package repository

import (
	"errors"
	"fmt"
	"github.com/samarec1812/sync-unisender/unisender/models"
	"gorm.io/gorm"
)

type AuthMysql struct {
	db *gorm.DB
}

func NewContactsMysql(db *gorm.DB) *AuthMysql {
	return &AuthMysql{db: db}
}

func (r *AuthMysql) Create(contact models.Contact) (uint64, error) {
	res := r.db.Create(&contact)
	if res.Error != nil {
		return 0, res.Error
	}
	return contact.ContactId, nil
}

func (r *AuthMysql) Update(contact models.Contact) (uint64, error) {
	res := r.db.Model(&contact).Where("contact_id = ?", contact.ContactId).Updates(contact)

	if res.Error != nil {
		return 0, res.Error
	}

	if res.RowsAffected == 0 {
		return 0, errors.New(fmt.Sprintf("contact with contact_id %d not found for update", contact.ContactId))

	}
	fmt.Println(res.RowsAffected)
	fmt.Println(contact.ID)

	return contact.ContactId, nil
}

func (r *AuthMysql) Delete(contactId uint64) (uint64, error) {
	var contact models.Contact
	res := r.db.Where("contact_id = ?", contactId).Delete(&contact)
	if res.Error != nil {
		return 0, res.Error
	}

	return contact.ContactId, nil
}

func (r *AuthMysql) Export(id uint) (models.ImportContact, error) {
	var contact models.ImportContact
	res := r.db.Table("contacts").Select("contacts.email, contacts.name, uu.token").Joins("left join unisender_users uu on contacts.account_id = uu.account_id").Where("contacts.id = ?", id).Find(&contact)
	if res.Error != nil {
		return models.ImportContact{}, res.Error
	}
	return contact, nil
}
