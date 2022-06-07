package repositories

import (
	"api/src/models"
	"database/sql"
)

// UserRepository is a struct that contains the user's information.
type userRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new user repository.
func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

// CreateUser creates a new user in the database.
func (ur *userRepository) CreateUser(user models.User) (uint64, error) {
	stmt, err := ur.db.Prepare(
		"INSERT INTO users (name, nickname, email, password) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ID), nil
}
