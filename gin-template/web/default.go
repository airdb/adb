package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRoot(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to use adb Generate.")
}
