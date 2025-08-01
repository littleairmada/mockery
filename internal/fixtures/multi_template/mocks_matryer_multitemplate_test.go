// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: matryer
// TEST MOCKERY BOILERPLATE

package multitemplate

import (
	"sync"
)

// Ensure that MockMatryerFoo does implement Foo.
// If this is not the case, regenerate this file with mockery.
var _ Foo = &MockMatryerFoo{}

// MockMatryerFoo is a mock implementation of Foo.
//
//	func TestSomethingThatUsesFoo(t *testing.T) {
//
//		// make and configure a mocked Foo
//		mockedFoo := &MockMatryerFoo{
//			BarFunc: func() string {
//				panic("mock out the Bar method")
//			},
//		}
//
//		// use mockedFoo in code that requires Foo
//		// and then make assertions.
//
//	}
type MockMatryerFoo struct {
	// BarFunc mocks the Bar method.
	BarFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// Bar holds details about calls to the Bar method.
		Bar []struct {
		}
	}
	lockBar sync.RWMutex
}

// Bar calls BarFunc.
func (mock *MockMatryerFoo) Bar() string {
	if mock.BarFunc == nil {
		panic("MockMatryerFoo.BarFunc: method is nil but Foo.Bar was just called")
	}
	callInfo := struct {
	}{}
	mock.lockBar.Lock()
	mock.calls.Bar = append(mock.calls.Bar, callInfo)
	mock.lockBar.Unlock()
	return mock.BarFunc()
}

// BarCalls gets all the calls that were made to Bar.
// Check the length with:
//
//	len(mockedFoo.BarCalls())
func (mock *MockMatryerFoo) BarCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockBar.RLock()
	calls = mock.calls.Bar
	mock.lockBar.RUnlock()
	return calls
}
