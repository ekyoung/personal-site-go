package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/browser", "./browser")

	r.HTMLRender = createMyRender()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home", gin.H{
			"title": "Home",
		})
	})

	r.GET("/trips", func(c *gin.Context) {
		c.HTML(http.StatusOK, "trips", gin.H{
			"title": "Trips",
		})
	})

	r.Run() // listen and server on 0.0.0.0:8080
}

func createMyRender() multitemplate.Render {
	r := multitemplate.New()
	r.AddFromFiles("home", "server/_shared/layout.tmpl", "server/root/index.tmpl")
	r.AddFromFiles("trips", "server/_shared/layout.tmpl", "server/trips/index.tmpl")

	return r
}
