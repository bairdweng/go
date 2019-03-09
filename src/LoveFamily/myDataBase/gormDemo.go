package myDataBase

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var Gdb *gorm.DB

func InitDataBaseWithDataBase(dataBaseName string) {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:8889)/"+dataBaseName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		println("数据库连接失败", err)
	}
	db.SingularTable(true)
	println("数据库连接成功", dataBaseName)
	Gdb = db
	// defer db.Close()
}

type Info struct {
	Id       int
	Title    string
	Content  string
	Module   string
	ImageURL string
	ShowType string
}
