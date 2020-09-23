package users

import (
	"net/http"
	"strconv"

	"github.com/csrias/bookstore_users-api/domain/users"
	"github.com/csrias/bookstore_users-api/services"
	"github.com/csrias/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

// CreateUser action
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// GetUser action
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequest("user id should be number")
		c.JSON(err.Status, err)
		return
	}

	user, err := services.GetUser(userID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

// SearchUser action
func SearchUser(c *gin.Context) {

}

// bytes, err := ioutil.ReadAll(c.Request.Body)
// if err != nil {
// 	// TODO: handle error
// 	return
// }
// if err := json.Unmarshal(bytes, &user); err != nil {
// 	// TODO: handle error
// 	return
// }
