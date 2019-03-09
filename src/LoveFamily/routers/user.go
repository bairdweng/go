package routers

import (
	// "LoveFamily/models"
	// "LoveFamily/myService"

	"LoveFamily/models"
	"LoveFamily/myService"

	"github.com/gin-gonic/gin"
)

func updateUserInfo(c *gin.Context) {
	var user models.Users
	err := c.Bind(&user)
	if err != nil {
		c.JSON(401, Error(err.Error(), nil))
	}
	myService.UpdateUserInfo(user)
	c.JSON(200, Successful(nil))
}
func getUserInfo(c *gin.Context) {
	Id := c.PostForm("Id")
	item := myService.GetUserInfo(Id)
	c.JSON(200, Successful(item))
}
