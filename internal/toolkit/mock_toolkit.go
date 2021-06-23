// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package toolkit

import mock "github.com/stretchr/testify/mock"

// MockToolkit is an autogenerated mock type for the Toolkit type
type MockToolkit struct {
	mock.Mock
}

// AddMask provides a mock function with given fields: mask
func (_m *MockToolkit) AddMask(mask string) {
	_m.Called(mask)
}

// AddPath provides a mock function with given fields: paths
func (_m *MockToolkit) AddPath(paths ...string) error {
	_va := make([]interface{}, len(paths))
	for _i := range paths {
		_va[_i] = paths[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...string) error); ok {
		r0 = rf(paths...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Debug provides a mock function with given fields: a
func (_m *MockToolkit) Debug(a ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, a...)
	_m.Called(_ca...)
}

// Debugf provides a mock function with given fields: format, a
func (_m *MockToolkit) Debugf(format string, a ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, a...)
	_m.Called(_ca...)
}

// EndGroup provides a mock function with given fields:
func (_m *MockToolkit) EndGroup() {
	_m.Called()
}

// Error provides a mock function with given fields: a
func (_m *MockToolkit) Error(a ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, a...)
	_m.Called(_ca...)
}

// Errorc provides a mock function with given fields: context
func (_m *MockToolkit) Errorc(context MessageContext) {
	_m.Called(context)
}

// Errorf provides a mock function with given fields: format, a
func (_m *MockToolkit) Errorf(format string, a ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, a...)
	_m.Called(_ca...)
}

// ExportVariable provides a mock function with given fields: name, value
func (_m *MockToolkit) ExportVariable(name string, value string) error {
	ret := _m.Called(name, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(name, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetInput provides a mock function with given fields: name
func (_m *MockToolkit) GetInput(name string) (string, bool) {
	ret := _m.Called(name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetInputList provides a mock function with given fields: name
func (_m *MockToolkit) GetInputList(name string) ([]string, bool) {
	ret := _m.Called(name)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetState provides a mock function with given fields: name
func (_m *MockToolkit) GetState(name string) (string, bool) {
	ret := _m.Called(name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// IsDebug provides a mock function with given fields:
func (_m *MockToolkit) IsDebug() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SetOutput provides a mock function with given fields: name, value
func (_m *MockToolkit) SetOutput(name string, value string) {
	_m.Called(name, value)
}

// SetState provides a mock function with given fields: name, value
func (_m *MockToolkit) SetState(name string, value string) {
	_m.Called(name, value)
}

// StartGroup provides a mock function with given fields: title
func (_m *MockToolkit) StartGroup(title string) {
	_m.Called(title)
}

// Warning provides a mock function with given fields: a
func (_m *MockToolkit) Warning(a ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, a...)
	_m.Called(_ca...)
}

// Warningc provides a mock function with given fields: context
func (_m *MockToolkit) Warningc(context MessageContext) {
	_m.Called(context)
}

// Warningf provides a mock function with given fields: format, a
func (_m *MockToolkit) Warningf(format string, a ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, a...)
	_m.Called(_ca...)
}
