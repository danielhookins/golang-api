package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/", home)
	router.GET("/article/:permalink/", getArticleByPermalink)

	router.Run(":8080")
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": "OK",
	})
}

// TODO: Return Article from Database
func getArticleByPermalink(c *gin.Context) {
	permalink := c.Param("permalink")

	c.JSON(200, gin.H{
		"permalink": permalink,
	})
}
