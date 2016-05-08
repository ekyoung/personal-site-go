package trips

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/ekyoung/personal-site-go/lib/trips"
)

func IndexController(c *gin.Context) {
    tripRepo := trips.NewTripRepository()
    trips := tripRepo.All()

    c.HTML(http.StatusOK, "trips", gin.H{
        "title":         "Trips",
        "isTripsActive": true,
        "trips":         trips,
    })
}
