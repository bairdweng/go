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
	//获取我的家人
	router.POST("/getMyFamily", userRouter.GetMyFamily)
	//删除关系
	router.POST("/deleteRelation", userRouter.DeleteRelation)
	router.Run(":9090") // listen and serve on 0.0.0.0:8080
}
