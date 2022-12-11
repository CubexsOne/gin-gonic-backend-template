package utils

import "github.com/gin-gonic/gin"

func ToGinError(ctx *gin.Context, err error, statusCode int) {
	ctx.Error(err).SetMeta(statusCode)
}
