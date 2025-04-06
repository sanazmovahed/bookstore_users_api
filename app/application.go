package app

import (
	"bookstore_users_api/logger"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	logger.Info("about to start the application...")
	mapUrls()
	router.Run(":8080")
}
