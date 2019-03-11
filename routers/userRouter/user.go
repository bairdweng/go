package userRouter

import (
	"LoveFamily/helper"
	"LoveFamily/models"
	"LoveFamily/myService"

	"github.com/gin-gonic/gin"
)

// 更新
func UpdateUserInfo(c *gin.Context) {
	var user models.Users
	err := c.Bind(&user)
	if err != nil {
		c.JSON(200, helper.Error(err.Error(), nil))
		return
	}
	uerr := myService.UpdateUserInfo(user)
	if uerr != nil {
		c.JSON(200, helper.Error(uerr.Error(), nil))
		return
	}
	c.JSON(200, helper.Successful(nil))
}

// 获取用户信息
func GetUserInfo(c *gin.Context) {
	Id := c.PostForm("Id")
	item := myService.GetUserInfo(Id)
	c.JSON(200, helper.Successful(item))
}

// 添加用户信息
func AddUser(c *gin.Context) {
	// 检查关系是否合法。
	relationShip := c.PostForm("RelationShip")
	err1 := helper.CheckRelationShip(relationShip)
	if err1 != nil {
		c.JSON(200, helper.Error(err1.Error(), nil))
		return
	}
	// 检查性别是否合法。
	gender := c.PostForm("Gender")
	err2 := helper.CheckGender(gender)
	if err2 != nil {
		c.JSON(200, helper.Error(err2.Error(), nil))
		return
	}
	// 检查用户是否存在。
	userId := c.PostForm("UserId")
	auser := myService.GetUserInfo(userId)
	if auser.Id == 0 {
		c.JSON(200, helper.Error("主用户不存在", nil))
		return
	}
	//绑定
	var user models.Users
	berr := c.Bind(&user)
	if berr != nil {
		c.JSON(200, helper.Error(berr.Error(), nil))
		return
	}
	//将用户添加到数据库。
	err, u := myService.AddUser(user)
	if err != nil {
		c.JSON(200, helper.Error(berr.Error(), nil))
	}
	//更新当前用户的关系。
	print("u=================", u.Id)

	// var user models.Users
	// err := c.Bind(&user)
	// if err != nil {
	// 	c.JSON(200, helper.Error(err.Error(), nil))
	// } else {
	// 	myService.AddUser(user)
	// 	c.JSON(200, helper.Successful(nil))
	// }
}
