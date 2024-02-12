package app

import (
	"fmt"
	"os"
	"time"

	"anggi.blog/utils/logs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	LoadEnv()
}

func StartApplication() {
	logs.Info.Println("Application started")

	router = gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8100"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "Content-Length", "X-CSRF-Token", "Token"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.AllowOriginFunc = func(origin string) bool {
		return origin == "http://localhost:8100"
	}
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))
	MapUrls()
	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))

}
