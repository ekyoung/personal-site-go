package trips

import (
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/gin-gonic/gin"

	mockLibTrips "github.com/ekyoung/personal-site-go/lib/trips/mocks"
	mockServerTrips "github.com/ekyoung/personal-site-go/server/trips/mocks"
)

func TestGallery_TripNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tripID := "fun-trip"
	mockedContext, _ := mockContext(map[string]string{
		"tripId": tripID,
	})

	mockTripRepo := mockLibTrips.NewMockTripRepository(ctrl)
	mockTripRepo.EXPECT().Lookup(tripID).Return(nil)

	mockPageNotFounder := mockServerTrips.NewMockPageNotFounder(ctrl)
	pageNotFoundCalled := false
	mockPageNotFounder.EXPECT().PageNotFound(mockedContext).Do(func(c *gin.Context) {
		pageNotFoundCalled = true
	})

	controller := NewTripController(mockTripRepo, mockPageNotFounder)

	controller.Gallery(mockedContext)

	if !pageNotFoundCalled {
		t.Error("Expected PageNotFound to be called, but it was not.")
	}
}

func mockContext(params map[string]string) (*gin.Context, *httptest.ResponseRecorder) {

	gin.SetMode(gin.TestMode)

	context, responseRecorder, _ := gin.CreateTestContext()

	var paramsSlice []gin.Param
	for key, value := range params {
		paramsSlice = append(paramsSlice, gin.Param{
			Key:   key,
			Value: value,
		})
	}

	context.Params = paramsSlice

	return context, responseRecorder
}
