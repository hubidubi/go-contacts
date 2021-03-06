package app

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

var DB *gorm.DB

func ListContacts(w http.ResponseWriter, r *http.Request) {
	log.Print("Listing contacts")
	contacts := []Contact{}
	err := DB.Find(&contacts).Error
	if errors.Is(err, gorm.ErrEmptySlice) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No records"))
	}
	json.NewEncoder(w).Encode(contacts)
}

func GetContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Print("Getting Contact by ID: ", vars["id"])
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal("Invalid id: ", vars["id"])
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Invalid id: " + vars["id"]))
		return
	}
	var contact Contact
	err = DB.First(&contact, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No record with id: " + vars["id"]))
		return
	}
	json.NewEncoder(w).Encode(contact)
}

func AddContact(w http.ResponseWriter, r *http.Request) {
	log.Print("Add contact ")
	contact := Contact{}
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		log.Fatal("error during decoding", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Provided data was incorrect"))
		return
	} else {
		err = DB.Create(&contact).Error
		json.NewEncoder(w).Encode(&contact)
	}
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	log.Print("Updating contact")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal("Invalid id ", id)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Invalid id: " + vars["id"]))
		return
	}
	contact := Contact{}
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		log.Fatal("error during decoding", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Provided data was incorrect"))
		return
	}
	contact.ID = uint(id)
	err = DB.Save(&contact).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No record with id: " + vars["id"]))
		return
	}

	json.NewEncoder(w).Encode(&contact)
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Print("Deleting Contact by ID: ", vars["id"])
	if id, err := strconv.Atoi(vars["id"]); err != nil {
		log.Fatal("Invalid id: ", id)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Invalid id: " + vars["id"]))
	} else {
		contact := Contact{}
		err := DB.First(&contact, id).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No record with id: " + vars["id"]))
			return
		} else {
			DB.Delete(&contact, id)
		}
	}
}

func SearchContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	log.Print("Searching contact: ", name)
	contacts := []Contact{}
	err := DB.Where("name LIKE ?", "%"+name+"%").Find(&contacts).Error
	if errors.Is(err, gorm.ErrEmptySlice) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No record with name: " + vars["name"]))
		return
	}
	json.NewEncoder(w).Encode(&contacts)
}
