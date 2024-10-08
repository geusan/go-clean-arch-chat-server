// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	domain "api-server/domain"

	mock "github.com/stretchr/testify/mock"
)

// ChatroomRepository is an autogenerated mock type for the ChatroomRepository type
type ChatroomRepository struct {
	mock.Mock
}

type ChatroomRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *ChatroomRepository) EXPECT() *ChatroomRepository_Expecter {
	return &ChatroomRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: name, owner
func (_m *ChatroomRepository) Create(name string, owner *domain.User) *domain.Chatroom {
	ret := _m.Called(name, owner)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *domain.Chatroom
	if rf, ok := ret.Get(0).(func(string, *domain.User) *domain.Chatroom); ok {
		r0 = rf(name, owner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Chatroom)
		}
	}

	return r0
}

// ChatroomRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ChatroomRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - name string
//   - owner *domain.User
func (_e *ChatroomRepository_Expecter) Create(name interface{}, owner interface{}) *ChatroomRepository_Create_Call {
	return &ChatroomRepository_Create_Call{Call: _e.mock.On("Create", name, owner)}
}

func (_c *ChatroomRepository_Create_Call) Run(run func(name string, owner *domain.User)) *ChatroomRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*domain.User))
	})
	return _c
}

func (_c *ChatroomRepository_Create_Call) Return(_a0 *domain.Chatroom) *ChatroomRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChatroomRepository_Create_Call) RunAndReturn(run func(string, *domain.User) *domain.Chatroom) *ChatroomRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: id
func (_m *ChatroomRepository) Delete(id uint) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChatroomRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ChatroomRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id uint
func (_e *ChatroomRepository_Expecter) Delete(id interface{}) *ChatroomRepository_Delete_Call {
	return &ChatroomRepository_Delete_Call{Call: _e.mock.On("Delete", id)}
}

func (_c *ChatroomRepository_Delete_Call) Run(run func(id uint)) *ChatroomRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *ChatroomRepository_Delete_Call) Return(_a0 error) *ChatroomRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChatroomRepository_Delete_Call) RunAndReturn(run func(uint) error) *ChatroomRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Fetch provides a mock function with given fields:
func (_m *ChatroomRepository) Fetch() []domain.Chatroom {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 []domain.Chatroom
	if rf, ok := ret.Get(0).(func() []domain.Chatroom); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Chatroom)
		}
	}

	return r0
}

// ChatroomRepository_Fetch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fetch'
type ChatroomRepository_Fetch_Call struct {
	*mock.Call
}

// Fetch is a helper method to define mock.On call
func (_e *ChatroomRepository_Expecter) Fetch() *ChatroomRepository_Fetch_Call {
	return &ChatroomRepository_Fetch_Call{Call: _e.mock.On("Fetch")}
}

func (_c *ChatroomRepository_Fetch_Call) Run(run func()) *ChatroomRepository_Fetch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ChatroomRepository_Fetch_Call) Return(_a0 []domain.Chatroom) *ChatroomRepository_Fetch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChatroomRepository_Fetch_Call) RunAndReturn(run func() []domain.Chatroom) *ChatroomRepository_Fetch_Call {
	_c.Call.Return(run)
	return _c
}

// FindById provides a mock function with given fields: id
func (_m *ChatroomRepository) FindById(id uint) *domain.Chatroom {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *domain.Chatroom
	if rf, ok := ret.Get(0).(func(uint) *domain.Chatroom); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Chatroom)
		}
	}

	return r0
}

// ChatroomRepository_FindById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindById'
type ChatroomRepository_FindById_Call struct {
	*mock.Call
}

// FindById is a helper method to define mock.On call
//   - id uint
func (_e *ChatroomRepository_Expecter) FindById(id interface{}) *ChatroomRepository_FindById_Call {
	return &ChatroomRepository_FindById_Call{Call: _e.mock.On("FindById", id)}
}

func (_c *ChatroomRepository_FindById_Call) Run(run func(id uint)) *ChatroomRepository_FindById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *ChatroomRepository_FindById_Call) Return(_a0 *domain.Chatroom) *ChatroomRepository_FindById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChatroomRepository_FindById_Call) RunAndReturn(run func(uint) *domain.Chatroom) *ChatroomRepository_FindById_Call {
	_c.Call.Return(run)
	return _c
}

// NewChatroomRepository creates a new instance of ChatroomRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChatroomRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChatroomRepository {
	mock := &ChatroomRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
