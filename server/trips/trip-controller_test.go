package trips_test

import (
	. "github.com/ekyoung/personal-site-go/server/trips"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"

	libTrips "github.com/ekyoung/personal-site-go/lib/trips"

	mockWrapper "github.com/ekyoung/personal-site-go/gin-wrapper/mocks"
	mockLibTrips "github.com/ekyoung/personal-site-go/lib/trips/mocks"
	mockServerTrips "github.com/ekyoung/personal-site-go/server/trips/mocks"
)

var _ = Describe("Trip Controller", func() {
	var (
		ctrl *gomock.Controller

		mockContext        *mockWrapper.MockContext
		mockTripRepo       *mockLibTrips.MockTripRepository
		mockPageNotFounder *mockServerTrips.MockPageNotFounder

		controller *TripController
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		mockContext = mockWrapper.NewMockContext(ctrl)
		mockTripRepo = mockLibTrips.NewMockTripRepository(ctrl)
		mockPageNotFounder = mockServerTrips.NewMockPageNotFounder(ctrl)

		controller = NewTripController(mockTripRepo, mockPageNotFounder)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Gallery Action", func() {
		It("should render an HTML template when the trip is found", func() {
			//RECORD aka Training my mocks aka ARRANGE
			tripID := "fun-trip"
			mockContext.EXPECT().Param("tripId").Return(tripID)

			trip := &libTrips.Trip{
				Id:   tripID,
				Name: "Fun Trip",
			}
			mockTripRepo.EXPECT().Lookup(tripID).Return(trip)

			//I really just want to assert that HTML is called with the right values
			mockContext.EXPECT().HTML(200, "trips/gallery", gomock.Any())

			//REPLAY aka ACT
			controller.Gallery(mockContext)
		})

		It("should call PageNotFound when the trip is not found", func() {
			//RECORD aka Training my mocks aka ARRANGE
			tripID := "fun-trip"
			mockContext.EXPECT().Param("tripId").Return(tripID)

			mockTripRepo.EXPECT().Lookup(tripID).Return(nil)

			//I really just want to assert that PageNotFound is called
			mockPageNotFounder.EXPECT().PageNotFound(mockContext)

			//REPLAY aka ACT
			controller.Gallery(mockContext)

		})
	})
})
