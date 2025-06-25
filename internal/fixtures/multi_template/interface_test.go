package multitemplate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	testifyMock := NewMockTestifyFoo(t)
	testifyMock.EXPECT().Bar().Return("bar")
	assert.Equal(t, "foo", testifyMock.Bar())

	matryerMock := MockMatryerFoo{
		BarFunc: func() string {
			return "foo"
		},
	}
	assert.Equal(t, "foo", matryerMock.Bar())
}
