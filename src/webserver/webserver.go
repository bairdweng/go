package webserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":9090") // listen and serve on 0.0.0.0:8080

	// fmt.Println("webserver 初始化")
	// //第一个参数为客户端发起http请求时的接口名，第二个参数是一个func，负责处理这个请求。
	// http.HandleFunc("/login", loginTask)

	// //服务器要监听的主机地址和端口号
	// err := http.ListenAndServe(":2078", nil)

	// if err != nil {
	// 	fmt.Println("ListenAndServe error: ", err.Error())
	// }
	// fmt.Println("服务器连接成功")
}
func loginTask(w http.ResponseWriter, req *http.Request) {
	fmt.Println("loginTask is running...")
}
