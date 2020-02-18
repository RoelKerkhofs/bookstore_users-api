package app

import (
  "bookstore/bookstore_users-api/controllers"
)
func mapUrls() {
  router.GET("/ping", controllers.Ping)
}
