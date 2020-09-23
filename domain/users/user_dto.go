package users

import (
	"strings"

	"github.com/csrias/bookstore_users-api/utils/errors"
)

// User struct
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"lst_name"`
	Email       string `json:"email"`
	CreatedDate string `json:"created_date"`
}

// Validate user
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequest("invalid email address")
	}

	return nil
}
