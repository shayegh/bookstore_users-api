package app

import (
	"github.com/shayegh/bookstore_users-api/controllers/ping"
	"github.com/shayegh/bookstore_users-api/controllers/users"
)

func mapURLs() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	// router.GET("/users/search",users.SearchUser)
}
