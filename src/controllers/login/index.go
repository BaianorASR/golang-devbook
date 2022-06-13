package loginController

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login is a function that handles the login.
func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Get the user from the request.
	user := models.User{}
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.UserRepository(db)
	userSavedInDB, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	// Check if the password is correct.
	if err = security.CheckPasswordHash(userSavedInDB.Password, user.Password); err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := auth.GenerateToken(userSavedInDB.ID)
	response.JSON(w, http.StatusOK, map[string]string{"token": token})
}
