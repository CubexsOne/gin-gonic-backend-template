package main

import (
	"net/http"
	"os"

	"github.com/cubexsone/gin-gonic-backend-template/src/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	var env = utils.ToDefault(os.Getenv("ENV"), "DEV")
	if env == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000")
}
