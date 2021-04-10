package server

import (
	"shoppingCart-LI/routers"

	"github.com/gin-gonic/gin"
)

func RunServer() {

	r := gin.Default()

	routers.CreateRouters(r)

	r.Run()

}
