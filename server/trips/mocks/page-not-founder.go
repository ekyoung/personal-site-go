package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/ekyoung/personal-site-go/gin-wrapper"
)

type PageNotFounder struct {
	mock.Mock
}

func (m *PageNotFounder) PageNotFound(c wrapper.Context) {
	m.Called(c)
}
