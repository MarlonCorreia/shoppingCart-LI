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
			cart.GET("/:id", controllers.GetCart)
			cart.DELETE("/:id", controllers.DeleteCart)
			cart.POST("/:id/product/:productId", controllers.PostCart)
			cart.DELETE("/:id/product/:productId", controllers.DeleteCartProduct)
			cart.POST("/:id/coupon/:couponId", controllers.PostCouponCart)
			cart.DELETE("/:id/coupon/:couponId", controllers.DeleteCouponCart)
		}

		user := api.Group("/user")
		{
			user.GET("/:id", controllers.GetUser)
			user.PUT("/create", controllers.CreateUser)
			user.POST("/login", controllers.LoginUser)
		}

		product := api.Group("/product")
		{
			product.GET("/", controllers.GetProducts)
			product.GET("/:id", controllers.GetProduct)
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
