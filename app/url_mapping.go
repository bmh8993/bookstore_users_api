package app

import (
	"github.com/bmh8993/bookstore_users_api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
