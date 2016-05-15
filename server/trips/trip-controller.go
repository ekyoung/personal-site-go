package trips

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ekyoung/personal-site-go/gin-wrapper"

	"github.com/ekyoung/personal-site-go/lib/trips"
)

type PageNotFounder interface {
	PageNotFound(c wrapper.Context)
}

type TripController struct {
	tripRepository trips.TripRepository
	pageNotFounder PageNotFounder
}

func NewTripController(tripRepository trips.TripRepository, pageNotFounder PageNotFounder) *TripController {
	return &TripController{
		tripRepository: tripRepository,
		pageNotFounder: pageNotFounder,
	}
}

func (ctrl *TripController) Index(c wrapper.Context) {
	trips := ctrl.tripRepository.All()

	c.HTML(http.StatusOK, "trips", gin.H{
		"title":         "Trips",
		"isTripsActive": true,
		"trips":         trips,
	})
}

func (ctrl *TripController) Gallery(c wrapper.Context) {
	tripId := c.Param("tripId")

	trip := ctrl.tripRepository.Lookup(tripId)

	if trip == nil {
		ctrl.pageNotFounder.PageNotFound(c)
		return
	}

	c.HTML(http.StatusOK, "trips/gallery", gin.H{
		"title":         "Trips",
		"isTripsActive": true,
		"trip":          trip,
	})
}

func (ctrl *TripController) SlideShow(c wrapper.Context) {
	tripId := c.Param("tripId")

	trip := ctrl.tripRepository.Lookup(tripId)

	if trip == nil {
		ctrl.pageNotFounder.PageNotFound(c)
		return
	}

	c.HTML(http.StatusOK, "trips/slide-show", gin.H{
		"title":         "Trips",
		"isTripsActive": true,
		"trip":          trip,
	})
}
