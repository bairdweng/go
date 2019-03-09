package routers

import (
	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string
	Password string
}

func Init() {

	router := gin.Default()
	//获取用户信息。
	router.POST("/getUserInfo", getUserInfo)
	//更新用户信息
	router.POST("/updateUserInfo", updateUserInfo)

	router.Run(":9090") // listen and serve on 0.0.0.0:8080
}
