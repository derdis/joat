package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func fetchData() gin.H {
	return gin.H{
		"message": "Hello, this is the backend!",
	}
}

func events(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.SSEvent("event", fetchData())
}

func main() {
	router := gin.Default()
	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, "this is the backend")
	})

	router.GET("/events", events)

	router.Run(":8080")
}
