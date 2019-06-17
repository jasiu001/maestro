// Code generated by mockery v1.0.0
package mocks

import bucket "github.com/jasiu001/maestro/bucket"

import mock "github.com/stretchr/testify/mock"

// BuckerSaver is an autogenerated mock type for the BuckerSaver type
type BuckerSaver struct {
	mock.Mock
}

// SaveBucket provides a mock function with given fields: b
func (_m *BuckerSaver) SaveBucket(b bucket.Bundle) error {
	ret := _m.Called(b)

	var r0 error
	if rf, ok := ret.Get(0).(func(bucket.Bundle) error); ok {
		r0 = rf(b)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
