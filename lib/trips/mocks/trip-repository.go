package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/ekyoung/personal-site-go/lib/trips"
)

type TripRepository struct {
	mock.Mock
}

func (m *TripRepository) All() []*trips.Trip {
	ret := m.Called()
	return ret.Get(0).([]*trips.Trip)
}

func (m *TripRepository) Lookup(tripId string) *trips.Trip {
	ret := m.Called(tripId)
	return ret.Get(0).(*trips.Trip)
}
