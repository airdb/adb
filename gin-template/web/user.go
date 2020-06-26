package web

import (
	"log"
	"net/http"

	"github.com/airdb-template/gin-api/model/vo"
	"github.com/gin-gonic/gin"
)

func ListUser(c *gin.Context) {
	user := c.Param("user")
	userInfo := vo.List(user)

	if userInfo == nil {
		log.Println("user info is nil.")
		c.String(http.StatusOK, "hello")

		return
	}
}
