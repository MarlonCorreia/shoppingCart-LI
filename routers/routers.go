package routers

import (
	"shoppingCart-LI/controllers"

	"github.com/gin-gonic/gin"
)

func CreateRouters(r *gin.Engine) {

	api := r.Group("/api")
	{
		cart := api.Group("/cart")
		{
			cart.GET("/", controllers.GetCart)
			cart.DELETE("/", controllers.DeleteCart)
			cart.POST("/product/:productId", controllers.PostCart)
			cart.DELETE("/product/:productId", controllers.DeleteCartProduct)
			cart.POST("/coupon/:couponId", controllers.PostCouponCart)
			cart.DELETE("/coupon/:couponId", controllers.DeleteCouponCart)
		}

		user := api.Group("/user")
		{
			user.GET("/", controllers.GetUser)
			user.PUT("/create", controllers.CreateUser)
			user.POST("/login", controllers.LoginUser)
		}

		product := api.Group("/product")
		{
			product.GET("/", controllers.GetProducts)
			product.GET("/:id", controllers.GetProduct)
			product.POST("/:id", controllers.PostProduct)
			product.PUT("/:id", controllers.PutProduct)
			product.DELETE("/:id", controllers.DeleteProduct)
		}

		coupons := api.Group("/coupon")
		{
			coupons.GET("/", controllers.GetAllCoupons)
			coupons.PUT("/:id", controllers.PutCoupon)
			coupons.DELETE("/:id", controllers.DeleteCoupon)

		}
	}
}
