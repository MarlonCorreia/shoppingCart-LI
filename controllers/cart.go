package controllers

import (
	"fmt"
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
	fmt.Println("cartid: ", paramCartId, "productid: ", paramProductId)

	productId, _ := strconv.ParseUint(paramProductId, 10, 32)
	cartId, _ := strconv.ParseUint(paramCartId, 10, 32)

	exists, err := models.ProductExists(uint(productId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	if exists {
		_, err := models.GetCart(uint(cartId))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "server error",
			})
			return
		}

		err = models.AddProductToCart(uint(cartId), uint(productId), 1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "server error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "product added to cart",
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "product not found",
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
	if total == 0 {
		discountedPrice = 0
	}

	responseCart := ResponseCart{
		ID:              cart.ID,
		Orders:          responseOrders,
		Coupon:          cart.DiscountCoupons,
		DiscountedPrice: discountedPrice,
		Total:           total - discountedPrice,
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
