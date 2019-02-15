package myDataBase

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/Go-SQL-Driver/MySQL"
)

//数据库配置
const (
	userName = "root"
	password = "root"
	ip       = "127.0.0.1"
	port     = "8889"
	dbName   = "FxxInfotech"
)

var DB *sql.DB

// func main() {
// 	InitDB()
// 	search()
// }

//注意方法名大写，就是public
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
}
