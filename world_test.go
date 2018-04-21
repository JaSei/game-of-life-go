package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	w := New(10, 5)
	assert.Len(t, w.latest, 10)
	assert.Len(t, w.latest[0], 5)

}

func TestIsStable(t *testing.T) {
	w := New(5, 5)

	assert.True(t, w.IsStable())

	w.latest[0][0] = 5 //Neighbors 2, allive
	assert.False(t, w.IsStable())

	w.NextGeneration()
	assert.True(t, w.IsStable())
}
