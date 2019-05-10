package userRouter

import (
	"LoveFamily/helper"
	"LoveFamily/models"
	"LoveFamily/myService"
	"strconv"

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
	idInt, userIdErr := strconv.Atoi(Id)
	if userIdErr != nil {
		c.JSON(200, helper.Error(userIdErr.Error(), nil))
		return
	}
	item := myService.GetUserInfo(idInt)
	c.JSON(200, helper.Successful(item))
}

// 添加用户信息 A是主用户，B表示被添加用户。例如我添加一个儿子，我是A,儿子是B
func AddUser(c *gin.Context) {
	// myService.UpdateUserInfoField(2, "son_id", "12")
	// return
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
	UserIdInt, userIdErr := strconv.Atoi(userId)
	if userIdErr != nil {
		c.JSON(200, helper.Error(userIdErr.Error(), nil))
		return
	}

	auser := myService.GetUserInfo(UserIdInt)
	if auser.Id == 0 {
		c.JSON(200, helper.Error("主用户不存在", nil))
		return
	}
	//模型绑定
	var user models.Users
	berr := c.Bind(&user)
	if berr != nil {
		c.JSON(200, helper.Error(berr.Error(), nil))
		return
	}
	err, addUser := myService.AddUser(user)
	if err != nil {
		c.JSON(200, helper.Error(berr.Error(), nil))
		return
	}

	//A的关系记录。
	relationA := models.Relations{}
	relationA.Id = UserIdInt
	relationA.OtherId = addUser.Id
	relationA.Relation = relationShip
	relationA.RelationCn = helper.GetRelationCn(relationShip)
	relationA.CreateTime = helper.GetCurrentDateTime()
	myService.AddRelations(relationA)

	//B的关系记录。
	reverseRalationShip := helper.ReverseRalationShip(addUser.Gender, relationShip)
	relationB := models.Relations{}
	relationB.Id = addUser.Id
	relationB.OtherId = UserIdInt
	relationB.Relation = reverseRalationShip
	relationB.RelationCn = helper.GetRelationCn(reverseRalationShip)
	relationB.CreateTime = helper.GetCurrentDateTime()
	myService.AddRelations(relationB)

	c.JSON(200, helper.Successful(nil))
}
func GetMyFamily(c *gin.Context) {
	Id := c.PostForm("Id")
	id_int, userIdErr := strconv.Atoi(Id)
	if userIdErr != nil {
		c.JSON(200, helper.Error(userIdErr.Error(), nil))
		return
	}
	relations := myService.GetRelations(id_int)
	var newItems = []models.Relations{}
	for _, model := range relations {
		item := myService.GetUserInfo(model.OtherId)
		model.UserInfo = item
		newItems = append(newItems, model)
	}
	c.JSON(200, helper.Successful(newItems))
}

//删除关系。
func DeleteRelation(c *gin.Context) {
	Id := c.PostForm("Id")
	Relation := c.PostForm("Relation")
	id_int, userIdErr := strconv.Atoi(Id)
	if userIdErr != nil {
		c.JSON(200, helper.Error(userIdErr.Error(), nil))
		return
	}
	OtherId := c.PostForm("OtherId")
	OtherId_int, OtherIdErr := strconv.Atoi(OtherId)
	if OtherIdErr != nil {
		c.JSON(200, helper.Error(OtherIdErr.Error(), nil))
		return
	}
	errA, relationItemA := myService.GetRelation(id_int, OtherId_int, Relation)
	if errA != nil {
		c.JSON(200, helper.Error(errA.Error(), nil))
		return
	}
	otherUser := myService.GetUserInfo(OtherId_int)
	errB, relationItemB := myService.GetRelation(OtherId_int, id_int, helper.ReverseRalationShip(otherUser.Gender, Relation))
	if errB != nil {
		c.JSON(200, helper.Error(errB.Error(), nil))
		return
	}
	myService.DeleRelation(relationItemA.RecordId)
	myService.DeleRelation(relationItemB.RecordId)
	c.JSON(200, helper.Successful(nil))
}
