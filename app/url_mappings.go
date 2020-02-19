package app

import (
	"bookstore/bookstore_users-api/controllers/ping"
	"bookstore/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/search", users.SearchUser)
	router.GET("/user/:user_id", users.GetUser)
}
