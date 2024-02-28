package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"name":    "api server 3",
			"version": "v1.0",
		})
	})
	r.Run()
}
