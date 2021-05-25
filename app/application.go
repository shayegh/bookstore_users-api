package app

import (
	"github.com/gin-gonic/gin"
	"github.com/shayegh/bookstore_users-api/logger"
)

var router = gin.Default()

func StartApplication() {
	mapURLs()
	logger.Info("About to start the app...")
	router.Run(":8080")
}
