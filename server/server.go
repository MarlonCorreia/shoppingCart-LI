package server

import (
	"os"
	"shoppingCart-LI/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getMode() string {
	godotenv.Load()
	mode := os.Getenv("GIN_DEBUG_MODE")

	if mode == "true" || mode == "" {
		return gin.DebugMode
	}

	return gin.ReleaseMode
}

func RunServer() {

	r := gin.Default()

	mode := getMode()
	gin.SetMode(mode)

	routers.CreateRouters(r)

	r.Run()

}
