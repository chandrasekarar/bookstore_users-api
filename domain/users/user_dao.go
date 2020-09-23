package users

import (
	"fmt"

	"github.com/csrias/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

// Get method
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFound(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.CreatedDate = result.CreatedDate

	return nil
}

// Save method
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]

	if current != nil {
		return errors.NewBadRequest(fmt.Sprintf("user %d already exists", user.ID))
	}
	usersDB[user.ID] = user
	return nil
}
