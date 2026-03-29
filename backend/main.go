package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func api(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "this is the backend")
}

func main() {
	router := gin.Default()
	router.GET("/api", api)

	router.Run(":8080")
}
