package controllers

import (
	"contactsBook/models"
	u "contactsBook/utils"
	"encoding/json"
	"net/http"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error!"))
		return
	}

	contact.UserId = user
	resp := contact.CreateContact()
	u.Respond(w, resp)
}

var GetContacts = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetContacts(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var UpdateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}
	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Ошибка при декодировании запроса для обновления контакта"))
		return
	}
	contact.UserId = user
	resp := contact.UpdateContact(user)
	u.Respond(w, resp)
}

var DeleteContact = func(w http.ResponseWriter, r *http.Request) {
	// Получение данных для удаления контакта
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}
	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Ошибка при декодировании запроса для удаления контакта"))
		return
	}
	contact.UserId = user
	resp := contact.DeleteContact()
	u.Respond(w, resp)
}
