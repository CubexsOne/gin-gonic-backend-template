package controller

import "github.com/gin-gonic/gin"

func Api(router *gin.Engine) *gin.RouterGroup {
	apiController := router.Group("/api")
	exampleController(apiController)

	return apiController
}
