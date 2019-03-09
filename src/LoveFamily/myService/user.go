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
	if user.Id != 0 {
		db.Save(&user)
	}
	return errors.New("卧槽")
}
