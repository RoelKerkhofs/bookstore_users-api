package app

import (
  "bookstore/bookstore_users-api/controllers"
)
func mapUrls() {
  router.GET("/ping", controllers.Ping)

  router.POST("/users", controllers.CreateUser)
  router.GET("/users/search", controllers.SearchUser)
  router.GET("/users/:user_id", controllers.GetUser)
}
