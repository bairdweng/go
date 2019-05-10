package myService

import (
	"LoveFamily/helper"
	"LoveFamily/models"
	"LoveFamily/myDataBase"
	"errors"
)

// 根据id获取用户信息。
func GetUserInfo(Id int) models.Users {
	var db = myDataBase.Gdb
	var user models.Users
	db.Where("id = ?", Id).Find(&user)
	return user
}

//更新用户信息。
func UpdateUserInfo(user models.Users) error {
	var db = myDataBase.Gdb
	var usr models.Users
	db.Where("id = ?", user.Id).Find(&usr)
	usr.UpdateTime = helper.GetCurrentDateTime()
	if usr.Id != 0 {
		db.Save(&user)
		return nil
	} else {
		return errors.New("用户不存在")
	}
}

//添加用户信息, 返回添加后的用户信息。
func AddUser(user models.Users) (error, models.Users) {
	var db = myDataBase.Gdb
	user.CreateTime = helper.GetCurrentDateTime()
	user.UpdateTime = helper.GetCurrentDateTime()
	db.NewRecord(user)
	db.Create(&user)
	var newUser models.Users
	db.Where("id = ?", user.Id).Find(&newUser)
	return nil, user
}

//添加关系,返回添加后的用户信息。
func AddRelations(relation models.Relations) (error, models.Relations) {
	var db = myDataBase.Gdb
	// db.Save(&relation)
	db.NewRecord(relation)
	db.Create(&relation)
	var obj models.Relations
	db.Where("id = ?", obj.Id).Find(&obj)
	return nil, obj
}

//更新用户局部信息
func UpdateUserInfoField(id int, key string, value interface{}) error {
	var db = myDataBase.Gdb
	var newUser = models.Users{}
	newUser.Id = id
	err2 := db.Model(&newUser).Update(key, value).Error
	if err2 != nil {
		return err2
	}
	return nil
}

//获取关系列表。
func GetRelations(id int) []models.Relations {
	var relations []models.Relations
	var db = myDataBase.Gdb
	db.Where("id = ?", id).Find(&relations)
	return relations
}

//获取关系
func GetRelation(id int, other_id int, relation string) (error, models.Relations) {
	var db = myDataBase.Gdb
	var relationItem models.Relations
	db.Where("id = ? And other_id = ? And relation = ?", id, other_id, relation).Find(&relationItem)
	if relationItem.Id == 0 {
		return errors.New("关系不存在"), relationItem
	}
	return nil, relationItem
}

//删除关系
func DeleRelation(id int) {
	var db = myDataBase.Gdb
	db.Where("record_id = ?", id).Delete(models.Relations{})
}
