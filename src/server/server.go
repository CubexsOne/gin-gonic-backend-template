package server

import (
	"github.com/cubexsone/gin-gonic-backend-template/src/controller"
	"github.com/cubexsone/gin-gonic-backend-template/src/environment"
	"github.com/cubexsone/gin-gonic-backend-template/src/server/middleware"
	"github.com/cubexsone/gin-gonic-backend-template/src/utils/log"
	"github.com/gin-gonic/gin"
)

func Server() {
	env := environment.ENV
	log.Info.Println("Server is running on:", env)
	if env == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(middleware.ErrorHandler)
	controller.Api(router)

	router.Run(":3000")
}
