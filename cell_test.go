package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCell(t *testing.T) {
	var c cell

	assert.False(t, c.IsAlive())
	c.Revival()
	assert.True(t, c.IsAlive())
	c.Die()
	assert.False(t, c.IsAlive())

	c.Revival()
	assert.Equal(t, byte(0), c.GetNeighbors())
	assert.True(t, c.IsAlive())
	c.SetNeighbors(4)
	assert.Equal(t, byte(4), c.GetNeighbors())
	assert.True(t, c.IsAlive())
}
