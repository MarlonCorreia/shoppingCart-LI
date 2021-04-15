package controllers

import (
	"net/http"
	"shoppingCart-LI/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResponseCart struct {
	ID              uint                    `json:"id"`
	Orders          []ResponseOrder         `json:"orders,omitempty"`
	Coupon          []models.DiscountCoupon `json:"coupon,omitempty"`
	DiscountedPrice float64                 `json:"discounted,omitempty" `
	Total           float64                 `json:"total"`
}

type ResponseOrder struct {
	ID       uint    `json:"productId"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantity"`
	SubTotal float64 `json:"subtotal"`
}

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
		"cart": cartResponse(cart),
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
		"cart": cartResponse(cart),
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

func cartResponse(cart *models.Cart) ResponseCart {
	var responseOrders []ResponseOrder
	var total float64

	for _, v := range cart.Orders {
		order := ResponseOrder{
			ID:       v.ProductID,
			Name:     v.Product.Name,
			Price:    v.Product.Price,
			Quantity: v.Quantity,
			SubTotal: v.Product.Price * float64(v.Quantity),
		}
		total = total + order.SubTotal

		responseOrders = append(responseOrders, order)
	}

	discountedPrice := discountedPrice(cart.DiscountCoupons)
	if total-discountedPrice <= 0 {
		discountedPrice = total
		total = 0
	} else {
		total = total - discountedPrice
	}

	responseCart := ResponseCart{
		ID:              cart.ID,
		Orders:          responseOrders,
		Coupon:          cart.DiscountCoupons,
		DiscountedPrice: discountedPrice,
		Total:           total,
	}

	return responseCart
}

func discountedPrice(coupons []models.DiscountCoupon) float64 {
	var total float64
	for _, v := range coupons {
		total = total + v.Price
	}

	return total
}
