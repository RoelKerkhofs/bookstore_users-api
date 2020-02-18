package users

import (
	"bookstore/bookstore_users-api/domain/users"
	"bookstore/bookstore_users-api/services"
	"bookstore/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   "bad_request",
		}
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO handle save error
		return
	}
	c.JSON(http.StatusCreated, result)
}

/*
func GetUser(c *gin.Context) {
  c.String(http.StatusNotImplemented, "Get User Route")
}
*/

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Search User Route")
}
