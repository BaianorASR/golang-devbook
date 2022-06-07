package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateUser creates a new user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewUserRepository(db)
	userID, err := repository.CreateUser(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("User created with ID: %d", userID)))
}

// GetAllUsers gets all users.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get All Users"))
}

// GetUserByID gets a user by ID.
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User By ID"))
}

// UpdateUser updates a user.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

// DeleteUser deletes a user.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
