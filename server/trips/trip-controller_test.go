package trips_test

import (
	. "github.com/ekyoung/personal-site-go/server/trips"

	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/mock"

	ginMocks "github.com/ekyoung/gin-mockable/mocks"
	libTrips "github.com/ekyoung/personal-site-go/lib/trips"
	libTripsMocks "github.com/ekyoung/personal-site-go/lib/trips/mocks"
	serverTripsMocks "github.com/ekyoung/personal-site-go/server/trips/mocks"
)

var _ = Describe("Trip Controller", func() {
	var (
		mockContext        *ginMocks.Context
		mockTripRepo       *libTripsMocks.TripRepository
		mockPageNotFounder *serverTripsMocks.PageNotFounder

		controller *TripController

		trip *libTrips.Trip
	)

	tripId := "fun-trip"

	BeforeEach(func() {
		mockContext = new(ginMocks.Context)
		mockContext.On("Param", "tripId").Return(tripId)
		mockContext.On("HTML", mock.Anything, mock.Anything, mock.Anything).Return()

		mockTripRepo = new(libTripsMocks.TripRepository)

		mockPageNotFounder = new(serverTripsMocks.PageNotFounder)
		mockPageNotFounder.On("PageNotFound", mock.Anything).Return()

		controller = NewTripController(mockTripRepo, mockPageNotFounder)
	})

	Describe("Gallery Action", func() {
		It("should render an HTML template when the trip is found", func() {
			//ARRANGE
			trip = &libTrips.Trip{
				Id:   tripId,
				Name: "Fun Trip",
			}
			mockTripRepo.On("Lookup", tripId).Return(trip)

			//ACT
			controller.Gallery(mockContext)

			//ASSERT
			mockContext.AssertCalled(GinkgoT(), "HTML", 200, "trips/gallery", mock.Anything)
		})

		It("should call PageNotFound when the trip is not found", func() {
			//ARRANGE
			trip = nil
			mockTripRepo.On("Lookup", tripId).Return(trip)

			//ACT
			controller.Gallery(mockContext)

			//ASSERT
			mockPageNotFounder.AssertCalled(GinkgoT(), "PageNotFound", mockContext)
		})
	})

	Describe("Slide Show Action", func() {
		It("Should render an HTML template when the trip is found", func() {
			//ARRANGE
			trip = &libTrips.Trip{
				Id:   tripId,
				Name: "Fun Trip",
			}
			mockTripRepo.On("Lookup", tripId).Return(trip)

			//ACT
			controller.SlideShow(mockContext)

			//ASSERT
			mockContext.AssertCalled(GinkgoT(), "HTML", 200, "trips/slide-show", mock.Anything)
		})

		It("should call PageNotFound when the trip is not found", func() {
			//ARRANGE
			trip = nil
			mockTripRepo.On("Lookup", tripId).Return(trip)

			//ACT
			controller.SlideShow(mockContext)

			//ASSERT
			mockPageNotFounder.AssertCalled(GinkgoT(), "PageNotFound", mockContext)
		})
	})
})
