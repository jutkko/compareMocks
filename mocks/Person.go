// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

// Person is an autogenerated mock type for the Person type
type Person struct {
	mock.Mock
}

// DoNothing provides a mock function with given fields:
func (_m *Person) DoNothing() {
	_m.Called()
}

// DoThings provides a mock function with given fields: _a0, _a1
func (_m *Person) DoThings(_a0 string, _a1 int) (int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, int) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
