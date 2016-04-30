package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"

	"github.com/ekyoung/personal-site-go/lib/trips"
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

	r.GET("/trips/:tripId", func(c *gin.Context) {
		tripId := c.Param("tripId")

		tripRepo := trips.NewTripRepository()
		trip := tripRepo.Lookup(tripId)

		c.HTML(http.StatusOK, "trips/gallery", gin.H{
			"title":         "Trips",
			"isTripsActive": true,
			"trip":          trip,
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
	r.AddFromFiles("home", "server/root/index.view.tmpl", "server/_shared/header.partial.tmpl", "server/_shared/main.layout.tmpl")
	r.AddFromFiles("trips", "server/trips/index.view.tmpl", "server/_shared/header.partial.tmpl", "server/_shared/main.layout.tmpl", "server/trips/left-nav.partial.tmpl")
	r.AddFromFiles("trips/gallery", "server/trips/gallery.view.tmpl", "server/_shared/header.partial.tmpl", "server/_shared/main.layout.tmpl", "server/trips/left-nav.partial.tmpl")
	r.AddFromFiles("about-this-site", "server/root/about-this-site.view.tmpl", "server/_shared/header.partial.tmpl", "server/_shared/main.layout.tmpl")
	r.AddFromFiles("resume", "server/root/resume.view.tmpl", "server/_shared/header.partial.tmpl", "server/_shared/main.layout.tmpl")

	return r
}
