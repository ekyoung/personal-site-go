package trips_test

import (
	. "github.com/ekyoung/personal-site-go/server/trips"

	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/mock"

	mockWrapper "github.com/ekyoung/personal-site-go/gin-wrapper/mocks"
	mockLibTrips "github.com/ekyoung/personal-site-go/lib/trips/mocks"
	mockServerTrips "github.com/ekyoung/personal-site-go/server/trips/mocks"

	libTrips "github.com/ekyoung/personal-site-go/lib/trips"
)

var _ = Describe("Trip Controller", func() {
	var (
		mockContext        *mockWrapper.Context
		mockTripRepo       *mockLibTrips.TripRepository
		mockPageNotFounder *mockServerTrips.PageNotFounder

		controller *TripController
	)

	BeforeEach(func() {
		mockContext = new(mockWrapper.Context)
		mockContext.On("HTML", mock.Anything, mock.Anything, mock.Anything).Return()

		mockTripRepo = new(mockLibTrips.TripRepository)

		mockPageNotFounder = new(mockServerTrips.PageNotFounder)
		mockPageNotFounder.On("PageNotFound", mock.Anything).Return()

		controller = NewTripController(mockTripRepo, mockPageNotFounder)
	})

	Describe("Gallery Action", func() {
		It("should render an HTML template when the trip is found", func() {
			//ARRANGE
			tripID := "fun-trip"
			mockContext.On("Param", "tripId").Return(tripID)

			trip := &libTrips.Trip{
				Id:   tripID,
				Name: "Fun Trip",
			}
			mockTripRepo.On("Lookup", tripID).Return(trip)

			//ACT
			controller.Gallery(mockContext)

			//ASSERT
			mockContext.AssertCalled(GinkgoT(), "HTML", 200, "trips/gallery", mock.Anything)
		})

		It("should call PageNotFound when the trip is not found", func() {
			//ARRANGE
			tripID := "fun-trip"
			mockContext.On("Param", "tripId").Return(tripID)

			mockTripRepo.On("Lookup", tripID).Return((*libTrips.Trip)(nil))

			//ACT
			controller.Gallery(mockContext)

			//ASSERT
			mockPageNotFounder.AssertCalled(GinkgoT(), "PageNotFound", mockContext)
		})
	})
})
