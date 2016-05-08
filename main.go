package main

import (
    "html/template"
    "net/http"

    "github.com/ekyoung/gin-nice-recovery"
    "github.com/gin-gonic/contrib/renders/multitemplate"
    "github.com/gin-gonic/gin"

    "github.com/ekyoung/personal-site-go/lib/trips"
)

func main() {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(nice.Recovery("error", gin.H{
        "title": "Error",
    }))

    r.Static("/browser", "./browser")

    r.HTMLRender = createMyRender()

    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "home", gin.H{
            "title":        "Home",
            "isHomeActive": true,
        })
    })

    r.GET("/trips", func(c *gin.Context) {
        tripRepo := trips.NewTripRepository()
        trips := tripRepo.All()

        c.HTML(http.StatusOK, "trips", gin.H{
            "title":         "Trips",
            "isTripsActive": true,
            "trips":         trips,
        })
    })

    r.GET("/trips/:tripId", func(c *gin.Context) {
        tripId := c.Param("tripId")

        tripRepo := trips.NewTripRepository()
        trip := tripRepo.Lookup(tripId)

        if trip == nil {
            renderPageNotFound(c)
            return
        }

        c.HTML(http.StatusOK, "trips/gallery", gin.H{
            "title":         "Trips",
            "isTripsActive": true,
            "trip":          trip,
        })
    })

    r.GET("/trips/:tripId/slide-show", func(c *gin.Context) {
        tripId := c.Param("tripId")

        tripRepo := trips.NewTripRepository()
        trip := tripRepo.Lookup(tripId)

        if trip == nil {
            renderPageNotFound(c)
            return
        }

        c.HTML(http.StatusOK, "trips/slide-show", gin.H{
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

    r.GET("/error", func(c *gin.Context) {
        c.HTML(http.StatusOK, "error", gin.H{
            "title": "Error",
        })
    })

    r.NoRoute(func(c *gin.Context) {
        renderPageNotFound(c)
    })

    r.Run() // listen and server on 0.0.0.0:8080
}

func createMyRender() multitemplate.Render {
    r := multitemplate.New()
    r.Add("home", template.Must(template.New("index.view.tmpl").Delims("[[", "]]").ParseFiles("server/root/index.view.tmpl", "server/_shared/header.partial.tmpl", "server/_shared/main.layout.tmpl")))
    r.Add("trips", template.Must(template.New("index.view.tmpl").Delims("[[", "]]").ParseFiles("server/trips/index.view.tmpl", "server/_shared/header.partial.tmpl", "server/_shared/main.layout.tmpl", "server/trips/left-nav.partial.tmpl")))
    r.Add("trips/gallery", template.Must(template.New("gallery.view.tmpl").Delims("[[", "]]").ParseFiles("server/trips/gallery.view.tmpl", "server/_shared/header.partial.tmpl", "server/_shared/main.layout.tmpl", "server/trips/left-nav.partial.tmpl")))
    r.Add("trips/slide-show", template.Must(template.New("slide-show.view.tmpl").Delims("[[", "]]").ParseFiles("server/trips/slide-show.view.tmpl", "server/_shared/header.partial.tmpl", "server/_shared/main.layout.tmpl", "server/trips/left-nav.partial.tmpl")))
    r.Add("about-this-site", template.Must(template.New("about-this-site.view.tmpl").Delims("[[", "]]").ParseFiles("server/root/about-this-site.view.tmpl", "server/_shared/header.partial.tmpl", "server/_shared/main.layout.tmpl")))
    r.Add("resume", template.Must(template.New("resume.view.tmpl").Delims("[[", "]]").ParseFiles("server/root/resume.view.tmpl", "server/_shared/header.partial.tmpl", "server/_shared/main.layout.tmpl")))
    r.Add("error", template.Must(template.New("error.view.tmpl").Delims("[[", "]]").ParseFiles("server/root/error.view.tmpl", "server/_shared/header.partial.tmpl", "server/_shared/main.layout.tmpl")))
    r.Add("page-not-found", template.Must(template.New("page-not-found.view.tmpl").Delims("[[", "]]").ParseFiles("server/root/page-not-found.view.tmpl", "server/_shared/header.partial.tmpl", "server/_shared/main.layout.tmpl")))

    return r
}

func renderPageNotFound(c *gin.Context) {
    c.HTML(http.StatusNotFound, "page-not-found", gin.H{
        "title": "Page Not Found",
    })
}
