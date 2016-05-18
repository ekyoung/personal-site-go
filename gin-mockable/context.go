package mockable

import (
	"github.com/gin-gonic/gin"
)

type Context interface {
	Param(key string) string
	HTML(code int, name string, obj interface{})
}

type HandlerFunc func(Context)

func Handler(fn HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		fn(c)
	}
}
