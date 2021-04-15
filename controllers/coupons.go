package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"shoppingCart-LI/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCoupons(c *gin.Context) {
	coupons := models.GetallDiscountCoupon()

	c.JSON(http.StatusOK, gin.H{
		"coupons": coupons,
	})
}

func PutCoupon(c *gin.Context) {
	message, _ := ioutil.ReadAll(c.Request.Body)
	paramId := c.Param("id")

	couponId, _ := strconv.ParseUint(paramId, 0, 32)

	var coupon models.DiscountCoupon

	json.Unmarshal(message, &coupon)
	coupon.ID = uint(couponId)

	if coupon.Name == "" || coupon.Price == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "request body not accepted",
		})
		return
	}

	models.CreateDiscountCoupon(coupon)

	c.JSON(http.StatusOK, gin.H{
		"message": "coupon created",
	})
}

func DeleteCoupon(c *gin.Context) {
	paramId := c.Param("id")
	couponId, _ := strconv.ParseUint(paramId, 10, 32)

	coupon, err := models.GetDiscountCoupon(uint(couponId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "coupon not found",
		})
		return
	}

	models.DeleteDiscountCoupon(&coupon)

	c.JSON(http.StatusOK, gin.H{
		"message": "coupon deleted",
	})

}
