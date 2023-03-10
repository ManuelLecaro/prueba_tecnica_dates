// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	core "github.com/ManuelLecaro/prueba_tecnica_dates/internal/domain/core"
	mock "github.com/stretchr/testify/mock"
)

// HolidayService is an autogenerated mock type for the HolidayService type
type HolidayService struct {
	mock.Mock
}

// GetHolidays provides a mock function with given fields: ctx, holiday
func (_m *HolidayService) GetHolidays(ctx context.Context, holiday core.Holiday) []core.Holiday {
	ret := _m.Called(ctx, holiday)

	var r0 []core.Holiday
	if rf, ok := ret.Get(0).(func(context.Context, core.Holiday) []core.Holiday); ok {
		r0 = rf(ctx, holiday)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]core.Holiday)
		}
	}

	return r0
}

type mockConstructorTestingTNewHolidayService interface {
	mock.TestingT
	Cleanup(func())
}

// NewHolidayService creates a new instance of HolidayService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHolidayService(t mockConstructorTestingTNewHolidayService) *HolidayService {
	mock := &HolidayService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
