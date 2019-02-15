package myService

import (
	"fmt"
	"myDataBase"
)

type CompanyItem struct {
	Id       int
	Title    string
	Content  string
	Module   string
	ImageURL string
	ShowType string
	Ctime    string
}

func GetCompanyInfo(id string) CompanyItem {
	var DB = myDataBase.DB
	var item = CompanyItem{}
	DB.Begin()
	rows, err := DB.Query("SELECT *FROM info where id = ?", id)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&item.Id, &item.Title, &item.Content, &item.Module, &item.ImageURL, &item.ShowType, &item.Ctime)
			fmt.Println("=========:", item.Content)
		}
	}
	return item
}
