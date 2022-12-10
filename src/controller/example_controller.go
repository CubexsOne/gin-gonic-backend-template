package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func exampleController(router *gin.RouterGroup) *gin.RouterGroup {
	controller := router.Group("/example")
	controller.GET("/special", specialRoute)

	return controller
}

func specialRoute(ctx *gin.Context) {
	ctx.String(http.StatusCreated, "Special")
}
