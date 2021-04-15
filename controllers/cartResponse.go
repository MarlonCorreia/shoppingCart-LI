package controllers

import "shoppingCart-LI/models"

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

func CartResponse(cart *models.Cart) ResponseCart {
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
