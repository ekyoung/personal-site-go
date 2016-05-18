package server

import (
	"html/template"

	"github.com/gin-gonic/gin"

	"github.com/ekyoung/gin-nice-recovery"
	"github.com/ekyoung/personal-site-go/gin-mockable"
	"github.com/gin-gonic/contrib/renders/multitemplate"

	libTrips "github.com/ekyoung/personal-site-go/lib/trips"

	"github.com/ekyoung/personal-site-go/server/root"
	"github.com/ekyoung/personal-site-go/server/trips"
)

type Server struct {
	TripRepository libTrips.TripRepository
}

func (s *Server) Run() {
	rootController := &root.RootController{}
	tripController := trips.NewTripController(s.TripRepository, rootController)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(nice.Recovery(rootController.RecoveryHandler))

	r.Static("/browser", "./browser")

	r.HTMLRender = createMyRender()

	r.GET("/", mockable.Handler(rootController.Index))

	r.GET("/trips", mockable.Handler(tripController.Index))
	r.GET("/trips/:tripId", mockable.Handler(tripController.Gallery))
	r.GET("/trips/:tripId/slide-show", mockable.Handler(tripController.SlideShow))

	r.GET("/about-this-site", mockable.Handler(rootController.AboutThisSite))
	r.GET("/resume", mockable.Handler(rootController.Resume))

	r.NoRoute(mockable.Handler(rootController.PageNotFound))

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
