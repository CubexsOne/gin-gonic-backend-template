package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		c.JSON(mapToHttpStatus(err))
		return
	}
}

func mapToHttpStatus(ginError *gin.Error) (int, *gin.Error) {
	httpStatus := ginError.Meta
	ginError.Meta = nil
	switch httpStatus {
	case 400:
		return http.StatusBadRequest, ginError
	case 401:
		return http.StatusUnauthorized, ginError
	case 403:
		return http.StatusForbidden, ginError
	case 404:
		return http.StatusNotFound, ginError
	default:
		return http.StatusInternalServerError, ginError
	}
}
