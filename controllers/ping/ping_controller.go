package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping action
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}
