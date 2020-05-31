package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTuple_IsValidDimensionWithInvalidX(t *testing.T) {
	tup := &Tuple{
		X: -1,
		Y: 2,
	}
	assert.False(t, tup.IsValidDimension())
}

func TestTuple_IsValidDimensionWithInvalidY(t *testing.T) {
	tup := &Tuple{
		X: 0,
		Y: -1,
	}
	assert.False(t, tup.IsValidDimension())
}

func TestTuple_IsValidDimensionWithInvalidXY(t *testing.T) {
	tup := &Tuple{
		X: 4,
		Y: 1,
	}
	assert.False(t, tup.IsValidDimension())
}

func TestTuple_IsValidDimension(t *testing.T) {
	tup := &Tuple{
		X: 1,
		Y: 2,
	}
	assert.True(t, tup.IsValidDimension())
}
