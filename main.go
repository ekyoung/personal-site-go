package main

import (
    "html/template"

    "github.com/ekyoung/gin-nice-recovery"
    "github.com/gin-gonic/contrib/renders/multitemplate"
    "github.com/gin-gonic/gin"

    libTrips "github.com/ekyoung/personal-site-go/lib/trips"

    "github.com/ekyoung/personal-site-go/server/root"
    "github.com/ekyoung/personal-site-go/server/trips"
)

func main() {
    rootController := &root.RootController{}

    tripRepo := libTrips.NewHardcodedTripRepository()
    tripController := trips.NewTripController(tripRepo, rootController)

    r := gin.New()
    r.Use(gin.Logger())
    r.Use(nice.Recovery("error", gin.H{
        "title": "Error",
    }))

    r.Static("/browser", "./browser")

    r.HTMLRender = createMyRender()

    r.GET("/", rootController.Index)

    r.GET("/trips", tripController.Index)
    r.GET("/trips/:tripId", tripController.Gallery)
    r.GET("/trips/:tripId/slide-show", tripController.SlideShow)

    r.GET("/about-this-site", rootController.AboutThisSite)
    r.GET("/resume", rootController.Resume)

    r.NoRoute(rootController.PageNotFound)

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
