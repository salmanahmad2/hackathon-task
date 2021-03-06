// Code generated by mockery v1.0.0. DO NOT EDIT.

package utils

import mock "github.com/stretchr/testify/mock"
import time "time"

// MockClock is an autogenerated mock type for the Clock type
type MockClock struct {
	mock.Mock
}

// NowLocal provides a mock function with given fields:
func (_m *MockClock) NowLocal() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// NowUTC provides a mock function with given fields:
func (_m *MockClock) NowUTC() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// NowUnix provides a mock function with given fields:
func (_m *MockClock) NowUnix() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}
