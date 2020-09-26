package users

import "encoding/json"

// PublicUser struct
type PublicUser struct {
	ID          int64  `json:"id"`
	CreatedDate string `json:"created_date"`
	Status      string `json:"status"`
}

// PrivateUser struct
type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	CreatedDate string `json:"created_date"`
	Status      string `json:"status"`
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for i, user := range users {
		result[i] = user.Marshall(isPublic)
	}
	return result
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:          user.ID,
			CreatedDate: user.CreatedDate,
			Status:      user.Status,
		}
	}
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}
