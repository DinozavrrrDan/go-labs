package models

import (
	u "contactsBook/utils"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user_id"`
}

func (contact *Contact) ValidateContact() (map[string]interface{}, bool) {

	if contact.Name == "" {
		return u.Message(false, "Name cannot be empty!"), false
	}

	if contact.Phone == "" {
		return u.Message(false, "Phone number cannot be empty!"), false
	}

	if contact.UserId <= 0 {
		return u.Message(false, "User not found!"), false
	}

	return u.Message(true, "success"), true
}

func (contact *Contact) CreateContact() map[string]interface{} {

	if response, ok := contact.ValidateContact(); !ok {
		return response
	}

	GetDB().Create(contact)

	resp := u.Message(true, "contact create")
	resp["contact"] = contact
	return resp
}

func (contact *Contact) UpdateContact(user uint) map[string]interface{} {

	if response, ok := contact.ValidateContact(); !ok {
		return response
	}
	err := GetDB().Table("contacts").Where("user_id = ?", user).Update("phone", contact.Phone).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return u.Message(false, "Contact not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}
	resp := u.Message(true, "contact update")
	resp["contact"] = contact
	return resp
}

func (contact *Contact) DeleteContact() map[string]interface{} {

	if response, ok := contact.ValidateContact(); !ok {
		return response
	}

	GetDB().Delete(contact)

	resp := u.Message(true, "contact delete")
	resp["contact"] = contact
	return resp
}

func GetContact(id uint) *Contact {

	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user uint) []*Contact {

	contactsSlice := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contactsSlice).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contactsSlice
}
