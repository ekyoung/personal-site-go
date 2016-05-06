package main

import (
    "errors"
    "fmt"

    "net/http"

    "html/template"

    "github.com/gin-gonic/contrib/renders/multitemplate"
    "github.com/gin-gonic/gin"

    "github.com/ekyoung/personal-site-go/lib/trips"
)

func main() {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(HandleErrors())

    r.Static("/browser", "./browser")

    r.HTMLRender = createMyRender()

    r.GET("/", func(c *gin.Context) {
        if false {
            panic("Aw, snap!")

        }

        if false {
            c.Error(errors.New("This is an error, not a panic.")) //Add an error to the list for the global error handler to deal with
            c.Abort()                                             //Don't run any subsequent handlers
            return                                                //Stop executing this handler
        }

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

    return r
}

func HandleErrors() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                fmt.Println("Handling a panic.")

                fmt.Println(err)

                //In a website, this would really be c.HTML with a pretty error page
                c.JSON(500, gin.H{
                    "status":  500,
                    "message": "Whomp, whomp! Panic!",
                })
            }
        }()
        c.Next() // execute all the handlers

        //Handle any errors collected
        //The gin.Logger middleware already prints errors so this is redundant
        for _, value := range c.Errors {
            fmt.Println("Collected error: " + value.Error())
        }

        if len(c.Errors) > 0 {
            //In a website, this would really be c.HTML with a pretty error page
            c.JSON(500, gin.H{
                "status":  500,
                "message": "Whomp, whomp! Server error!",
            })
        }
    }
}
