package router

import "github.com/gin-gonic/gin"

func InitRouter() {
	// start router using gin default config
	r := gin.Default()
	// initialize routes
	initializeRoutes(r)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
