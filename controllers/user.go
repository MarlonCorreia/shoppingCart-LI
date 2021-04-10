package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": "name",
	})
}

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": "created",
	})
}

func LoginUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"login": "sucess",
	})
}
