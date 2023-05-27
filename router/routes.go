package router

import (
	"github.com/francisko-rezende/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {

	handler.InitHandler()
	// create a router group, kinda like a new instance. Will use instead of r from now on
	v1 := r.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening/:id", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}

}
