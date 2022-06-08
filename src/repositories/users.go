package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// UserRepository is a struct that contains the user's information.
type userRepository struct {
	db *sql.DB
}

// UserRepository creates a new user repository.
func UserRepository(db *sql.DB) *userRepository {
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

// GetUser gets a user by name or nickname from the database.
func (ur *userRepository) GetUser(name string) ([]models.User, error) {
	nameOrNickname := fmt.Sprintf("%%%s%%", name)

	stmt, err := ur.db.Prepare(
		"SELECT id, name, nickname, email, created_at FROM users WHERE name LIKE ? OR nickname LIKE ?",
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(nameOrNickname, nameOrNickname)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		if err := rows.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
