package main

import (
	"LoveFamily/myDataBase"
	"LoveFamily/routers"
)

func main() {
	myDataBase.InitDB()
	myDataBase.InitDataBaseWithDataBase("LoveFamily")
	routers.Init()
}
