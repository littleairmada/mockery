package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	m := newMockfoo(t)
	m.EXPECT().Bar().Return(baz("foo"))
	assert.Equal(t, "foo", m.Bar())
}
