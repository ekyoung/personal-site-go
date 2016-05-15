package trips_test

import (
	. "github.com/ekyoung/personal-site-go/server/trips"

	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/mock"

	wrapperMocks "github.com/ekyoung/personal-site-go/gin-wrapper/mocks"
	libTrips "github.com/ekyoung/personal-site-go/lib/trips"
	libTripsMocks "github.com/ekyoung/personal-site-go/lib/trips/mocks"
	serverTripsMocks "github.com/ekyoung/personal-site-go/server/trips/mocks"
)

var _ = Describe("Trip Controller", func() {
	var (
		mockContext        *wrapperMocks.Context
		mockTripRepo       *libTripsMocks.TripRepository
		mockPageNotFounder *serverTripsMocks.PageNotFounder

		controller *TripController
	)

	tripId := "fun-trip"

	BeforeEach(func() {
		mockContext = new(wrapperMocks.Context)
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
			trip := &libTrips.Trip{
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
			mockTripRepo.On("Lookup", tripId).Return((*libTrips.Trip)(nil))

			//ACT
			controller.Gallery(mockContext)

			//ASSERT
			mockPageNotFounder.AssertCalled(GinkgoT(), "PageNotFound", mockContext)
		})
	})
})
