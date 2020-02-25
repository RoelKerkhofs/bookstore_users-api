package users

import (
	"bookstore/bookstore_users-api/domain/users"
	"bookstore/bookstore_users-api/services"
	"bookstore/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUserId(userParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userParam, 10, 64)

	if userErr != nil {
		return 0, errors.NewBadRequestError("User Id should be a number")
	}

	return userId, nil
}

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-public") == "true"))
}

func Get(c *gin.Context) {
	userId, idErr := GetUserId(c.Param("user_id"))

	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-public") == "true"))

}

func Update(c *gin.Context) {
	userId, idErr := GetUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.ID = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-public") == "true"))
}

func Delete(c *gin.Context) {
	userId, idErr := GetUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.DeleteUser(userId); err != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.Search(status)
	if err != nil {
		c.JSON(err.Status, err.Message)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-public") == "true"))
}
