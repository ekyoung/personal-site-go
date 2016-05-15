package mocks

import "github.com/stretchr/testify/mock"

type Context struct {
	mock.Mock
}

func (m *Context) Param(key string) string {
	ret := m.Called(key)
	return ret.String(0)
}

func (m *Context) HTML(code int, name string, obj interface{}) {
	m.Called(code, name, obj)
}
