package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/ekyoung/personal-site-go/gin-mockable"
)

type PageNotFounder struct {
	mock.Mock
}

func (m *PageNotFounder) PageNotFound(c mockable.Context) {
	m.Called(c)
}
