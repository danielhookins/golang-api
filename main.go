package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	router.GET("/ping/:name/:action", ping)
	
	router.Run(":8080")
}

func ping(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message:= name + " is " + action

	// c.String(http.StatusOK, message)
	c.JSON(200, gin.H{
		"name": name,
		"action": action,
		"message": message,
	})
}
