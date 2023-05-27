package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// start router using gin default config
	r := gin.Default()
	// define route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("8080") // listen and serve on 0.0.0.0:8080
}
