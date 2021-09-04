package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", home)
	router.GET("/ping/:name/:action", ping)

	router.Run(":8080")
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": "OK",
	})
}

func ping(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action

	// c.String(http.StatusOK, message)
	c.JSON(200, gin.H{
		"name":    name,
		"action":  action,
		"message": message,
	})
}
