package routes

import (
	"github.com/gin-gonic/gin"
	"go-bike-microservice/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	bikes := r.Group("/bikes")
	{
		bikes.GET("/", controllers.FindBikes)
		bikes.GET("/:id", controllers.FindBike)
		bikes.POST("/", controllers.CreateBike)
		bikes.PATCH("/:id", controllers.UpdateBike)
		bikes.DELETE("/:id", controllers.DeleteBike)
	}

	return r
}
