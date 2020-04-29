package numpygo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomFunc(t *testing.T) {
	arr := Rand("FLOAT64", 10)
	assert.Equal(t, 1, len(arr.Shape.Values))
	assert.Equal(t, 10, arr.Shape.Values[0])
	assert.Less(t, float64(0), arr.Elements.Values[0])
	assert.Greater(t, float64(1), arr.Elements.Values[0])
}

func TestMultidimensionalRandomFunc(t *testing.T) {
	arr := Rand("FLOAT64", 10, 10, 10)
	assert.Equal(t, 3, len(arr.Shape.Values))
	assert.Equal(t, 10, arr.Shape.Values[0])
	assert.Less(t, float64(0), arr.Elements.Values[0])
	assert.Greater(t, float64(1), arr.Elements.Values[0])
}
