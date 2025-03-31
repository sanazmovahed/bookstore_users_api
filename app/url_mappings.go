package app

import (
	"bookstore_users_api/controllers/ping"
	"bookstore_users_api/controllers/user"
)

func mapUrls() {
	router.Handle("GET", "/ping", ping.Ping)

	router.Handle("GET", "/users/:userid", user.SearchUser)
	// router.Handle("GET", "/users/search", controllers.SearchUser)
	router.Handle("POST", "/users", user.CreateUser)
}
