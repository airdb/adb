package web

import (
	"log"

	"github.com/airdb-template/gin-api/model/vo"
	"github.com/airdb/sailor/enum"
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

func ListUser(c *gin.Context) {
	user := c.Param("user")
	userInfo := vo.List(user)

	if userInfo == nil {
		log.Println("user info is nil.")
		middlewares.SetResp(
			c,
			enum.AirdbFailed,
			vo.UserResp{
				Nickname:   user,
				Headimgurl: "null",
			},
		)

		return
	}

	userInfo.Nickname = user

	middlewares.SetResp(
		c,
		enum.AirdbSuccess,
		userInfo,
	)
}
