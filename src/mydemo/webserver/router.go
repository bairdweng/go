package webserver

import (
	"mydemo/myService"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/getCompanyInfoById", func(c *gin.Context) {
		id := c.PostForm("id")
		item := myService.GetCompanyInfo(id)
		c.JSON(200, item)
	})
	router.Run(":9090") // listen and serve on 0.0.0.0:8080
}
