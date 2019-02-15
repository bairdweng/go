package webserver

import (
	// "myService"

	"myService"

	"github.com/gin-gonic/gin"
)

func Init() {

	router := gin.Default()

	// myService.getCompanyInfoById()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/getCompanyInfoById", func(c *gin.Context) {
		id := c.PostForm("id")

		// item := myService.CompanyItem{}
		// print(item.Id)

		// var a myService.CompanyItem
		// company.getCompanyInfoById()
		item := myService.GetCompanyInfo(id)

		// fmt.Println("------------", item.content)
		c.JSON(200, gin.H{
			"content": item.Content,
		})
	})
	router.Run(":9090") // listen and serve on 0.0.0.0:8080
}
