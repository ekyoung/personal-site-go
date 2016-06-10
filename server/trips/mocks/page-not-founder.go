package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/ekyoung/gin-mockable"
)

type PageNotFounder struct {
	mock.Mock
}

func (m *PageNotFounder) PageNotFound(c mockable.Context) {
	m.Called(c)
}
