// Code generated by mockery v2.36.0. DO NOT EDIT.

package ia

import mock "github.com/stretchr/testify/mock"

// mockVmAlert is an autogenerated mock type for the vmAlert type
type mockVmAlert struct {
	mock.Mock
}

// RequestConfigurationUpdate provides a mock function with given fields:
func (_m *mockVmAlert) RequestConfigurationUpdate() {
	_m.Called()
}

// newMockVmAlert creates a new instance of mockVmAlert. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockVmAlert(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockVmAlert {
	mock := &mockVmAlert{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
