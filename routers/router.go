package routers

import (
	"LoveFamily/routers/userRouter"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string
	Password string
}

func Init() {

	router := gin.Default()
	//获取用户信息。
	router.POST("/getUserInfo", userRouter.GetUserInfo)
	//更新用户信息
	router.POST("/updateUserInfo", userRouter.UpdateUserInfo)
	//添加用户信息
	router.POST("/addUserInfo", userRouter.AddUser)

	router.Run(":9090") // listen and serve on 0.0.0.0:8080
}
