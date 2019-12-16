package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := New()
	assert.NotNil(t, s)
	assert.Equal(t, 0, s.Len())

	_, ok := s.Pop()
	assert.False(t, ok)

	s.Push(10)
	assert.Equal(t, 1, s.Len())

	v, ok := s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 10, v)
}
