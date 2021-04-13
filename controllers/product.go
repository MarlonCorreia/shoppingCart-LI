package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shoppingCart-LI/models"
	"shoppingCart-LI/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type product struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func GetProducts(c *gin.Context) {
	products, err := models.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})

}

func GetProduct(c *gin.Context) {
	paramId := c.Param("id")
	if paramId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "no product id provided",
		})
		return
	}
	productId, _ := strconv.ParseUint(paramId, 10, 64)

	exist, err := models.ProductExists(uint(productId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server problem",
		})
		return
	}

	if exist {
		product, err := models.GetProduct(uint(productId))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "server problem",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"product": product,
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "product not found",
	})

}

func PutProduct(c *gin.Context) {
	prodId := c.Param("id")
	if prodId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "no product id provided",
		})
		return
	}

	message, _ := ioutil.ReadAll(c.Request.Body)
	var p product
	json.Unmarshal(message, &p)

	if p.Name == "" || p.Status == "" || p.Price == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing or empty fields",
		})
		return
	}
	id, _ := strconv.ParseUint(prodId, 10, 64)
	p.ID = uint(id)

	pex, _ := models.ProductExists(p.ID)
	if pex {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "product already created",
		})
		return
	}

	fmt.Println(utils.Jsonfy(p))
	err := models.CreateProduct(p.ID, p.Name, p.Price, p.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product created",
		"product": p,
	})

}

func DeleteProduct(c *gin.Context) {
	paramId := c.Param("id")
	productId, _ := strconv.ParseUint(paramId, 10, 32)

	exists, err := models.ProductExists(uint(productId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server problem",
		})
		return
	}

	if exists {
		err := models.DeleteProduct(uint(productId))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "server problem",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "product deleted",
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "product does not exist",
	})

}
