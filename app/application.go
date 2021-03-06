package app

import (
	"github.com/csrias/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication inits the app
func StartApplication() {
	mapUrls()
	logger.Info("about to start the application...")
	router.Run(":8080")
}
