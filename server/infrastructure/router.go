package infrastructure

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello world",
		})
	})

	Router = router
}