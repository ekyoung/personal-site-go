package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.Static("/browser", "./browser")

	r.LoadHTMLGlob("server/**/*.tmpl")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "root/index.tmpl", gin.H{
			"title": "Home",
		})
	})

	r.GET("/trips", func(c *gin.Context) {
		c.HTML(http.StatusOK, "trips/index.tmpl", gin.H{
			"title": "Trips",
		})
	})

	r.Run() // listen and server on 0.0.0.0:8080
}
