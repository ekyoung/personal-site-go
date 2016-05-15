package wrapper

import (
	"github.com/gin-gonic/gin"
)

type Context interface {
	Param(key string) string
	HTML(code int, name string, obj interface{})
}

func Wrap(f func(Context)) func(*gin.Context) {
	return func(c *gin.Context) {
		f(c)
	}
}
