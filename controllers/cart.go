package controllers

import (
	"fmt"
	"net/http"
	"shoppingCart-LI/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
			"Data": cart,
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
