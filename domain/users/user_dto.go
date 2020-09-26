package users

import (
	"strings"

	"github.com/csrias/bookstore_users-api/utils/errors"
)

const (
	//StatusActive indicates active status
	StatusActive = "active"
)

// User struct
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	CreatedDate string `json:"created_date"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

// Users slice of user
type Users []User

// Validate user
func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequest("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequest("invalid password")
	}

	return nil
}
