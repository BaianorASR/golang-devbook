package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CreateUser creates a new user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.PrepareUser(); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.UserRepository(db)
	user.ID, err = repository.CreateUser(user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

// GetAllUsers gets all users.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// println("\nusers")
}

// GetUserByID gets a user by ID.
func GetUserByID(w http.ResponseWriter, r *http.Request) {

}

// UpdateUser updates a user.
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// DeleteUser deletes a user.
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
