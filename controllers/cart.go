package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"get": "route",
	})
}

func PostCart(c *gin.Context) {
	message, _ := ioutil.ReadAll(c.Request.Body)

	println(string(message))

	c.JSON(200, gin.H{
		"post": "route",
	})
}
