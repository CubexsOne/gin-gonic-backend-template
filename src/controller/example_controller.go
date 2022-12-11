package controller

import (
	"errors"
	"net/http"

	"github.com/cubexsone/gin-gonic-backend-template/src/utils"
	"github.com/gin-gonic/gin"
)

func exampleController(router *gin.RouterGroup) *gin.RouterGroup {
	controller := router.Group("/example")
	controller.GET("/special", specialRoute)
	controller.GET("/specialNotFound", specialRouteNotFound)

	return controller
}

func specialRoute(ctx *gin.Context) {
	ctx.String(http.StatusCreated, "Special")
}

func specialRouteNotFound(ctx *gin.Context) {
	err := errors.New("special error")
	utils.ToGinError(ctx, err, 404)
}
