package controllers

import (
	"net/http"
	"shoppingCart-LI/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "no token provided",
		})
		return
	}

	cart, err := models.UserCartByToken(token)
	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cart": CartResponse(cart),
	})
	return

}

func PostCart(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "no token provided",
		})
		return
	}

	paramProductId := c.Param("productId")
	queryAmount := c.Query("amount")

	productId, _ := strconv.ParseUint(paramProductId, 10, 32)
	amount, _ := strconv.Atoi(queryAmount)

	if amount <= 0 {
		amount = 1
	}

	cart, err := models.UserCartByToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized",
		})
		return
	}

	product, err := models.GetProduct(uint(productId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "product not found",
		})
		return
	}

	models.AddProductToCart(cart, &product, int64(amount))

	c.JSON(http.StatusOK, gin.H{
		"message": "product added to cart",
	})
	return

}

func DeleteCart(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "no token provided",
		})
		return
	}

	cart, err := models.UserCartByToken(token)
	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized",
		})
		return
	}

	models.CleanCart(cart)
	c.JSON(http.StatusOK, gin.H{
		"cart": CartResponse(cart),
	})

}

func DeleteCartProduct(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	paramProductId := c.Param("productId")
	queryAmount := c.Query("amount")

	productId, _ := strconv.ParseUint(paramProductId, 10, 32)
	amount, _ := strconv.Atoi(queryAmount)

	if token == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "no token provided",
		})
		return
	}

	cart, err := models.UserCartByToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized",
		})
	}

	product, err := models.GetProduct(uint(productId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "product not found",
		})
		return
	}

	err = models.RemoveProductFromnCart(cart, &product, int64(amount))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "product not on cart",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product removed from cart",
	})
	return

}

func PostCouponCart(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	paramCouponId := c.Param("couponId")

	couponId, _ := strconv.ParseUint(paramCouponId, 10, 32)

	if token == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "no token provided",
		})
		return
	}

	cart, err := models.UserCartByToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized",
		})
	}

	coupon, err := models.GetDiscountCoupon(uint(couponId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "coupon not found",
		})
		return
	}

	models.AddDiscountCouponToCart(cart, &coupon)

	c.JSON(http.StatusOK, gin.H{
		"message": "coupon added to cart",
	})
	return

}

func DeleteCouponCart(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	paramCouponId := c.Param("couponId")

	couponId, _ := strconv.ParseUint(paramCouponId, 10, 32)

	if token == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "no token provided",
		})
		return
	}

	cart, err := models.UserCartByToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized",
		})
		return
	}

	coupon, err := models.GetDiscountCoupon(uint(couponId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "coupon not found",
		})
		return
	}

	models.DeleteDiscountCouponFromCart(cart, &coupon)

	c.JSON(http.StatusOK, gin.H{
		"message": "coupon deleted from cart",
	})
	return

}
