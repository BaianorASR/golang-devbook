package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User is a struct that contains the u's information.
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// PrepareUser is a function that prepares u and validate.
func (u User) PrepareUser(stage string) error {
	if err := u.validateUser(stage); err != nil {
		return err
	}

	// user, err := u.formatUser(stage)
	// if err != nil {
	// 	return nil, err
	// }

	return nil
}

// validateUser is a function that validates the u's information.
func (u *User) validateUser(stage string) error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Nickname == "" {
		return errors.New("nickname is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("invalid email format")
	}
	if stage == "signUp" && u.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

// FormatUser is a function that formats the u's information.
func (u *User) FormatUser(stage string) (User, error) {
	u.Password = strings.TrimSpace(u.Password)
	if stage == "signUp" {
		hash, err := security.HashPassword(u.Password)
		if err != nil {
			return User{}, err
		}

		u.Password = string(hash)
	}

	u.Name = strings.TrimSpace(u.Name)
	u.Nickname = strings.TrimSpace(u.Nickname)
	u.Email = strings.TrimSpace(u.Email)

	return *u, nil
}
