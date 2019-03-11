package myService

import (
	"LoveFamily/models"
	"LoveFamily/myDataBase"
	"errors"
)

// 根据id获取用户信息。
func GetUserInfo(Id string) models.Users {
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
	db.Save(&user)
	var newUser models.Users
	db.Where("id = ?", user.Id).Find(&newUser)
	return nil, user
}
