package trips

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/ekyoung/personal-site-go/server"

    "github.com/ekyoung/personal-site-go/lib/trips"
)

func SlideShowController(c *gin.Context) {
    tripId := c.Param("tripId")

    tripRepo := trips.NewTripRepository()
    trip := tripRepo.Lookup(tripId)

    if trip == nil {
        server.RenderPageNotFound(c)
        return
    }

    c.HTML(http.StatusOK, "trips/slide-show", gin.H{
        "title":         "Trips",
        "isTripsActive": true,
        "trip":          trip,
    })
}
