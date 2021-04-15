package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"shoppingCart-LI/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products, _ := models.GetAllProducts()

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

	product, err := models.GetProduct(uint(productId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "product not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
	return

}

func PutProduct(c *gin.Context) {
	prodId := c.Param("id")

	message, _ := ioutil.ReadAll(c.Request.Body)
	var p models.Product
	json.Unmarshal(message, &p)

	if p.Name == "" || p.Status == "" || p.Price == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "request body not accepted",
		})
		return
	}
	id, _ := strconv.ParseUint(prodId, 10, 64)
	p.ID = uint(id)

	_, err := models.GetProduct(p.ID)
	if err == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "product already exists",
		})
		return
	}

	models.CreateProduct(p.ID, p.Name, p.Price, p.Status)
	c.JSON(http.StatusOK, gin.H{
		"message": "product created",
	})

}

func PostProduct(c *gin.Context) {
	paramId := c.Param("id")
	productId, _ := strconv.ParseUint(paramId, 10, 32)

	message, _ := ioutil.ReadAll(c.Request.Body)

	var p models.Product
	json.Unmarshal(message, &p)

	product, err := models.GetProduct(uint(productId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "product not found",
		})
		return
	}
	models.UpdateProduct(&product, p)
	c.JSON(http.StatusOK, gin.H{
		"message": "product updated",
	})

}

func DeleteProduct(c *gin.Context) {
	paramId := c.Param("id")
	productId, _ := strconv.ParseUint(paramId, 10, 32)

	product, err := models.GetProduct(uint(productId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "product not found",
		})
		return
	}
	models.DeleteProduct(&product)

	c.JSON(http.StatusOK, gin.H{
		"message": "product deleted",
	})
	return

}
