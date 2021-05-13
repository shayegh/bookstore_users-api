package app

import (
	"github.com/shayegh/bookstore_users-api/controllers/ping"
	"github.com/shayegh/bookstore_users-api/controllers/users"
)

func mapURLs() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	// router.GET("/users/search",users.SearchUser)
}
