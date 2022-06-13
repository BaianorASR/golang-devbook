package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// USERRepository is a struct that contains the user's information.
type USERRepository struct {
	db *sql.DB
}

// UserRepository creates a new user repository.
func UserRepository(db *sql.DB) *USERRepository {
	return &USERRepository{db}
}

// CreateUser creates a new user in the database.
func (ur *USERRepository) CreateUser(user models.User) (uint64, error) {
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

// GetAllUsers gets all users from the database.
func (ur *USERRepository) GetAllUsers() ([]models.User, error) {
	stmt, err := ur.db.Prepare("SELECT id, name, nickname, email, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
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

// GetUser gets a user by name or nickname from the database.
func (ur *USERRepository) GetUser(name string) ([]models.User, error) {
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

// GetUserByID gets a user by ID from the database.
func (ur USERRepository) GetUserByID(id uint64) (models.User, error) {
	stmt, err := ur.db.Prepare("SELECT id, name, nickname, email, created_at FROM users WHERE id = ?")
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	row, err := stmt.Query(id)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// UpdateUser updates a user in the database.
func (ur *USERRepository) UpdateUser(ID uint64, user models.User) error {
	stmt, err := ur.db.Prepare(
		"UPDATE users SET name = ?, nickname = ?, email = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Nickname, user.Email, ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes a user from the database.
func (ur *USERRepository) DeleteUser(ID uint64) error {
	stmt, err := ur.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ID)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByEmail gets a user by email from the database.
func (ur *USERRepository) GetUserByEmail(email string) (models.User, error) {
	stmt, err := ur.db.Prepare("SELECT  email, password FROM users WHERE email = ?")
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	row, err := stmt.Query(email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err := row.Scan(&user.Email, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}
