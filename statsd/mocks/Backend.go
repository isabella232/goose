// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"

import time "time"

// Backend is an autogenerated mock type for the Backend type
type Backend struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, name, value, tags, rate
func (_m *Backend) Count(ctx context.Context, name string, value int64, tags []string, rate float64) error {
	ret := _m.Called(ctx, name, value, tags, rate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, []string, float64) error); ok {
		r0 = rf(ctx, name, value, tags, rate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Decr provides a mock function with given fields: ctx, name, tags, rate
func (_m *Backend) Decr(ctx context.Context, name string, tags []string, rate float64) error {
	ret := _m.Called(ctx, name, tags, rate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, float64) error); ok {
		r0 = rf(ctx, name, tags, rate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Distribution provides a mock function with given fields: ctx, name, value, tags, rate
func (_m *Backend) Distribution(ctx context.Context, name string, value float64, tags []string, rate float64) error {
	ret := _m.Called(ctx, name, value, tags, rate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, float64, []string, float64) error); ok {
		r0 = rf(ctx, name, value, tags, rate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Gauge provides a mock function with given fields: ctx, name, value, tags, rate
func (_m *Backend) Gauge(ctx context.Context, name string, value float64, tags []string, rate float64) error {
	ret := _m.Called(ctx, name, value, tags, rate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, float64, []string, float64) error); ok {
		r0 = rf(ctx, name, value, tags, rate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Histogram provides a mock function with given fields: ctx, name, value, tags, rate
func (_m *Backend) Histogram(ctx context.Context, name string, value float64, tags []string, rate float64) error {
	ret := _m.Called(ctx, name, value, tags, rate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, float64, []string, float64) error); ok {
		r0 = rf(ctx, name, value, tags, rate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Incr provides a mock function with given fields: ctx, name, tags, rate
func (_m *Backend) Incr(ctx context.Context, name string, tags []string, rate float64) error {
	ret := _m.Called(ctx, name, tags, rate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, float64) error); ok {
		r0 = rf(ctx, name, tags, rate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Set provides a mock function with given fields: ctx, name, value, tags, rate
func (_m *Backend) Set(ctx context.Context, name string, value string, tags []string, rate float64) error {
	ret := _m.Called(ctx, name, value, tags, rate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []string, float64) error); ok {
		r0 = rf(ctx, name, value, tags, rate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Timing provides a mock function with given fields: ctx, name, value, tags, rate
func (_m *Backend) Timing(ctx context.Context, name string, value time.Duration, tags []string, rate float64) error {
	ret := _m.Called(ctx, name, value, tags, rate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration, []string, float64) error); ok {
		r0 = rf(ctx, name, value, tags, rate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}