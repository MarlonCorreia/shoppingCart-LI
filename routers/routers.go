package routers

import (
	"shoppingCart-LI/controllers"

	"github.com/gin-gonic/gin"
)

func CreateRouters(r *gin.Engine) {

	api := r.Group("/api")
	{
		api.GET("/cart/:id", controllers.GetCart)
		api.POST("/cart/:id", controllers.PostCart)
	}

	users := r.Group("/user")
	{
		users.GET("/:id", controllers.GetUser)
		users.POST("/create", controllers.CreateUser)
		users.POST("/login", controllers.LoginUser)
	}
}
