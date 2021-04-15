package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"shoppingCart-LI/models"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func GetUser(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "no token provided",
		})
		return
	}

	user, err := models.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func CreateUser(c *gin.Context) {
	message, _ := ioutil.ReadAll(c.Request.Body)
	var u User

	json.Unmarshal(message, &u)

	if u.Name == "" || u.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad payload, check docs",
		})
	} else {
		err := models.CreateUser(u.Name, u.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "username taken",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Creation suceeded!",
				"user":    u,
			})
		}
	}

}

func LoginUser(c *gin.Context) {
	message, _ := ioutil.ReadAll(c.Request.Body)
	var u User

	json.Unmarshal(message, &u)

	if u.Name == "" || u.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "login failed, empty user fields!",
		})
		return
	}

	if !models.CheckUserExists(u.Name, u.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "login failed, user does not exists!",
		})
		return
	}

	token := models.GetUserToken(u.Name, u.Password)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
