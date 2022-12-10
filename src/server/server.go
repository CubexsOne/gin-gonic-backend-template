package server

import (
	"os"

	"github.com/cubexsone/gin-gonic-backend-template/src/controller"
	"github.com/cubexsone/gin-gonic-backend-template/src/utils"
	"github.com/gin-gonic/gin"
)

func Server() {
	var env = utils.ToDefault(os.Getenv("ENV"), "DEV")
	if env == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	controller.Api(router)

	router.Run(":3000")
}
