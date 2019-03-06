package main

import (
	"fmt"
	"mydemo/myDataBase"
	"mydemo/webserver"
)

func main() {
	myDataBase.InitDB()
	// webserver.Init()
	// myDataBase.DemoInit()
	myDataBase.InitDataBaseWithDataBase("FxxInfotech")
	webserver.Init()

	// search()
}
func search() {
	var DB = myDataBase.DB
	DB.Begin()
	rows, err := DB.Query("SELECT *FROM info")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		for rows.Next() {
			var id int
			var title string
			var content string
			var module string
			var imageURL string
			var showType string
			var ctime string
			err = rows.Scan(&id, &title, &content, &module, &imageURL, &showType, &ctime)
			// print("err", err.Error())
			fmt.Println("=========:", content)
		}
	}
}
