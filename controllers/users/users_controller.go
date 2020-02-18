package users

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func CreateUser(c *gin.Context) {
  c.String(http.StatusNotImplemented, "Create User Route")
}

func GetUser(c *gin.Context) {
  c.String(http.StatusNotImplemented, "Get User Route")
}

func SearchUser(c *gin.Context) {
  c.String(http.StatusNotImplemented, "Search User Route")
}
