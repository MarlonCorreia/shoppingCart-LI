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
		c.JSON(http.StatusForbidden, gin.H{
			"message": "no token provided",
		})
		return
	}
	cartId := c.Param("id")
	if cartId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "no cartId provided",
		})
		return
	}

	authorized, err := models.CheckTokenExists(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server problem",
		})
		return
	}

	if authorized {
		id, err := strconv.ParseUint(cartId, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "pass a valid cartId",
			})
			return
		}

		cart, err := models.GetCart(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "server problem",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Data": cartResponse(cart),
		})
		return

	}

	c.JSON(http.StatusForbidden, gin.H{
		"message": "Not authorized",
	})
}

func PostCart(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "no token provided",
		})
		return
	}

	paramCartId := c.Param("id")
	paramProductId := c.Param("productId")
	queryAmount := c.Query("amount")

	productId, _ := strconv.ParseUint(paramProductId, 10, 32)
	cartId, _ := strconv.ParseUint(paramCartId, 10, 32)
	amount, _ := strconv.Atoi(queryAmount)

	if amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "amount must be greater than zero",
		})
		return
	}

	cart, err := models.GetCart(uint(cartId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "cart not found",
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

	models.AddProductToCart(&cart, &product, int64(amount))

	c.JSON(http.StatusOK, gin.H{
		"message": "product added to cart",
	})
	return

}

func DeleteCart(c *gin.Context) {
	paramId := c.Param("id")

	cartId, _ := strconv.ParseUint(paramId, 0, 32)

	cart, err := models.GetCart(uint(cartId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "cart not found",
		})
		return
	}

	models.CleanCart(&cart)
	c.JSON(http.StatusOK, gin.H{
		"cart": cart,
	})

}

func DeleteCartProduct(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	paramCartId := c.Param("id")
	paramProductId := c.Param("productId")
	queryAmount := c.Query("amount")

	productId, _ := strconv.ParseUint(paramProductId, 10, 32)
	cartId, _ := strconv.ParseUint(paramCartId, 10, 32)
	amount, _ := strconv.Atoi(queryAmount)

	if amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "amount must be greater than zero",
		})
		return
	}
	if token == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "no token provided",
		})
		return
	}
	authorized, err := models.CheckTokenExists(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server problem",
		})
		return
	}

	if authorized {

		cart, err := models.GetCart(uint(cartId))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "cart not found",
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

		err = models.RemoveProductFromnCart(&cart, &product, int64(amount))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "product not on cart",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "product removed",
		})
		return

	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "not authorized",
	})
}

func PostCouponCart(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	paramCartId := c.Param("id")
	paramCouponId := c.Param("couponId")

	cartId, _ := strconv.ParseUint(paramCartId, 10, 32)
	couponId, _ := strconv.ParseUint(paramCouponId, 10, 32)

	if token == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "no token provided",
		})
		return
	}

	authorized, err := models.CheckTokenExists(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server problem",
		})
		return
	}

	if authorized {
		cart, err := models.GetCart(uint(cartId))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "cart not found",
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

		models.AddDiscountCouponToCart(&cart, &coupon)

		c.JSON(http.StatusOK, gin.H{
			"message": "coupon added to cart",
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "not authorized",
	})
}

func DeleteCouponCart(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	paramCartId := c.Param("id")
	paramCouponId := c.Param("couponId")

	cartId, _ := strconv.ParseUint(paramCartId, 10, 32)
	couponId, _ := strconv.ParseUint(paramCouponId, 10, 32)

	if token == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "no token provided",
		})
		return
	}

	authorized, err := models.CheckTokenExists(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server problem",
		})
		return
	}

	if authorized {
		cart, err := models.GetCart(uint(cartId))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "cart not found",
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

		models.DeleteDiscountCouponFromCart(&cart, &coupon)

		c.JSON(http.StatusOK, gin.H{
			"message": "coupon deleted",
		})
		return

	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "not authorized",
	})

}

func cartResponse(cart models.Cart) ResponseCart {
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
