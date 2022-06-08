package models

import (
	"errors"
	"strings"
	"time"
)

// User is a struct that contains the user's information.
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// PrepareUser is a function that prepares user and validate.
func (u User) PrepareUser() error {
	u.formatUser()
	if err := u.validateUser(); err != nil {
		return err
	}
	return nil
}

func (user *User) validateUser() error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Nickname == "" {
		return errors.New("nickname is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

func (user *User) formatUser() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
}
