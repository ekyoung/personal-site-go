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
			"title":        "Home",
			"isHomeActive": true,
		})
	})

	r.GET("/trips", func(c *gin.Context) {
		c.HTML(http.StatusOK, "trips", gin.H{
			"title":         "Trips",
			"isTripsActive": true,
		})
	})

	r.GET("/about-this-site", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about-this-site", gin.H{
			"title":                 "About This Site",
			"isAboutThisSiteActive": true,
		})
	})

	r.GET("/resume", func(c *gin.Context) {
		c.HTML(http.StatusOK, "resume", gin.H{
			"title":          "Resume",
			"isResumeActive": true,
		})
	})

	r.Run() // listen and server on 0.0.0.0:8080
}

func createMyRender() multitemplate.Render {
	r := multitemplate.New()
	r.AddFromFiles("home", "server/_shared/layout.tmpl", "server/root/index.tmpl", "server/_shared/header.tmpl")
	r.AddFromFiles("trips", "server/_shared/layout.tmpl", "server/trips/index.tmpl", "server/_shared/header.tmpl")
	r.AddFromFiles("about-this-site", "server/_shared/layout.tmpl", "server/root/about-this-site.tmpl", "server/_shared/header.tmpl")
	r.AddFromFiles("resume", "server/_shared/layout.tmpl", "server/root/resume.tmpl", "server/_shared/header.tmpl")

	return r
}
