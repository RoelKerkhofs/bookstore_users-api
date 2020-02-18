package users

import (
	"bookstore/bookstore_users-api/domain/users"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO handle error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		//TODO handle json error
		return
	}

	c.String(http.StatusNotImplemented, string(bytes))
}

/*
func GetUser(c *gin.Context) {
  c.String(http.StatusNotImplemented, "Get User Route")
}
*/

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Search User Route")
}
