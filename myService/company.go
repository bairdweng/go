package myService

import "LoveFamily/myDataBase"

type Infoee struct {
	Id       int
	Title    string
	Content  string
	Module   string
	ImageURL string
	ShowType string
	Ctime    string
}
type CompanyItem struct {
	Id       int
	Title    string
	Content  string
	Module   string
	ImageURL string
	ShowType string
	Ctime    string
}

//设置InfoEE为info
func (Infoee) TableName() string {
	return "info"
}

/*
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
}*/

func GetCompanyInfo(id string) Infoee {
	var db = myDataBase.Gdb
	var info Infoee
	db.First(&info)
	return info
}
