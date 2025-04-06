package app

import (
	"bookstore_users_api/controllers/ping"
	"bookstore_users_api/controllers/user"
)

func mapUrls() {
	router.Handle("GET", "/ping", ping.Ping)

	router.Handle("GET", "/users/:userid", user.Get)
	router.Handle("POST", "/users", user.Create)
	router.Handle("PUT", "/users/:userid", user.Update)
	router.Handle("PATCH", "/users/:userid", user.Update)
	router.Handle("DELETE", "/users/:userid", user.Delete)
	router.Handle("GET", "/internal/users/search", user.Search)
}
