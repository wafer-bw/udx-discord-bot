// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// Handle provides a mock function with given fields: w, r
func (_m *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}