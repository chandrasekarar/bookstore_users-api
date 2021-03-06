package services

import (
	"github.com/csrias/bookstore_users-api/domain/users"
	"github.com/csrias/bookstore_users-api/utils/cryptoutils"
	"github.com/csrias/bookstore_users-api/utils/dateutils"
	"github.com/csrias/bookstore_users-api/utils/errors"
)

type userService struct{}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
	LoginUser(request users.LoginRequest) (*users.User, *errors.RestErr)
}

var (
	UserService usersServiceInterface = &userService{}
)

// CreateUser service
func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.CreatedDate = dateutils.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = cryptoutils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser service
func (s *userService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateUser service
func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current := &users.User{ID: user.ID}
	if err := current.Get(); err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.FirstName != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

//DeleteUser service
func (s *userService) DeleteUser(userID int64) *errors.RestErr {
	current := &users.User{ID: userID}
	if err := current.Delete(); err != nil {
		return err
	}
	return nil
}

// Search service
func (s *userService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *userService) LoginUser(r users.LoginRequest) (*users.User, *errors.RestErr) {
	dao := &users.User{
		Email:    r.Email,
		Password: cryptoutils.GetMd5(r.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
