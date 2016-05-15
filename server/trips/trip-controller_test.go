package trips

import (
	"github.com/golang/mock/gomock"
	"testing"

	libTrips "github.com/ekyoung/personal-site-go/lib/trips"

	mockWrapper "github.com/ekyoung/personal-site-go/gin-wrapper/mocks"
	mockLibTrips "github.com/ekyoung/personal-site-go/lib/trips/mocks"
	mockServerTrips "github.com/ekyoung/personal-site-go/server/trips/mocks"
)

func TestGallery_TripFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() //VERIFY (aka ASSERT for this test)

	//Setup that will be common to all tests
	mockContext := mockWrapper.NewMockContext(ctrl)
	mockTripRepo := mockLibTrips.NewMockTripRepository(ctrl)
	mockPageNotFounder := mockServerTrips.NewMockPageNotFounder(ctrl)
	controller := NewTripController(mockTripRepo, mockPageNotFounder)

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
}

func TestGallery_TripNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() //VERIFY (aka ASSERT for this test)

	//Setup that will be common to all tests
	mockContext := mockWrapper.NewMockContext(ctrl)
	mockTripRepo := mockLibTrips.NewMockTripRepository(ctrl)
	mockPageNotFounder := mockServerTrips.NewMockPageNotFounder(ctrl)
	controller := NewTripController(mockTripRepo, mockPageNotFounder)

	//RECORD aka Training my mocks aka ARRANGE
	tripID := "fun-trip"
	mockContext.EXPECT().Param("tripId").Return(tripID)

	mockTripRepo.EXPECT().Lookup(tripID).Return(nil)

	//I really just want to assert that PageNotFound is called
	mockPageNotFounder.EXPECT().PageNotFound(mockContext)

	//REPLAY aka ACT
	controller.Gallery(mockContext)
}
